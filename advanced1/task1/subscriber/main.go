package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"time"
)

func main() {
	client, err := ethclient.Dial("xxxx")
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	subHeader, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-subHeader.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Printf("\n=== 新区块通知 ===\n")
			fmt.Printf("区块号: %v\n", header.Number)
			fmt.Printf("区块哈希: %v\n", header.Hash().Hex())
			fmt.Printf("时间戳: %v\n", time.Unix(int64(header.Time), 0))

			// 使用区块号查询，比哈希查询更可靠
			var block *types.Block
			var err error

			// 添加重试机制
			for i := 0; i < 3; i++ {
				block, err = client.BlockByNumber(context.Background(), header.Number)
				if err == nil {
					// 验证区块哈希是否匹配
					if block.Hash() == header.Hash() {
						break
					} else {
						fmt.Printf("区块哈希不匹配，重试中... (%d/3)\n", i+1)
						err = fmt.Errorf("block hash mismatch, header Hash: %v, block Hash: %v", header.Hash(), block.Hash())
						fmt.Printf("Header Number: %v\n", header.Number)
						fmt.Printf("Block Number: %v\n", block.Number())
					}
				} else {
					fmt.Printf("获取区块失败，重试中... (%d/3): %v\n", i+1, err)
				}

				if i < 2 { // 不是最后一次重试
					time.Sleep(time.Second * 10) // 等待2秒后重试
				}
			}

			if err != nil {
				log.Printf("获取区块详情失败，跳过此区块: %v\n", err)
				continue // 跳过这个区块，继续处理下一个
			}

			// 成功获取区块，输出详细信息
			fmt.Printf("\n--- 区块详细信息 ---\n")
			fmt.Printf("区块哈希: %v\n", block.Hash().Hex())
			fmt.Printf("父区块哈希: %v\n", block.ParentHash().Hex())
			fmt.Printf("交易根哈希: %v\n", block.TxHash().Hex())
			fmt.Printf("收据根哈希: %v\n", block.ReceiptHash().Hex())
			fmt.Printf("区块时间: %v\n", time.Unix(int64(block.Time()), 0))
			fmt.Printf("区块号: %v\n", block.Number())
			fmt.Printf("区块随机数: %v\n", block.Nonce())
			fmt.Printf("Gas限制: %v\n", block.GasLimit())
			fmt.Printf("Gas使用: %v\n", block.GasUsed())
			fmt.Printf("交易数量: %v\n", len(block.Transactions()))

			// 如果有交易，显示前几个交易的哈希
			txs := block.Transactions()
			if len(txs) > 0 {
				fmt.Printf("\n--- 交易列表 (前5个) ---\n")
				for i, tx := range txs {
					if i >= 5 { // 只显示前5个交易
						fmt.Printf("... 还有 %d 个交易\n", len(txs)-5)
						break
					}
					fmt.Printf("交易 %d: %v\n", i+1, tx.Hash().Hex())
				}
			}
			fmt.Printf("==================\n")
		}
	}
}
