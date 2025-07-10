package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func fromWei(wei *big.Int, decimals int) string {
	fWei := new(big.Float).SetInt(wei)
	unit := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
	result := new(big.Float).Quo(fWei, unit)
	return result.Text('f', decimals)

}

func main() {
	client, err := ethclient.Dial("xxxxx")
	if err != nil {
		log.Fatal(err)
	}

	// 先解析出私钥
	privateKey, err := crypto.HexToECDSA("xxxxx")
	if err != nil {
		log.Fatal(err)
	}

	// 从私钥提取公钥, 签名数据需要
	publicKey := privateKey.Public()
	// 断言为ECDSA类型
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	// 根据公钥获取from address
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 查询账户余额
	fromBalance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 设置需要转账的金额
	toAmount := big.NewInt(1_000_000_000_000_000)

	// 对比转账的金额, 如果账户余额不足，就提示
	if fromBalance.Cmp(toAmount) == -1 {
		log.Fatalf("Current Balance: %s ETH, The amount to be transferred amount %s ETH, and the are still %s ETH missing \n", fromWei(fromBalance, 18), fromWei(toAmount, 18), fromWei(new(big.Int).Sub(toAmount, fromBalance), 18))
	}

	log.Printf("Transfer Amount: %s \n", fromWei(toAmount, 18))
	// 获取下一次交易的nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)

	var gasLimit uint64 = 21_000
	gasPrice, err := client.SuggestGasPrice(context.Background())
	//gasPrice = new(big.Int).Add(gasPrice, big.NewInt(1_000_000_000))
	log.Printf("gas price: %s \n", gasPrice)
	log.Printf("gas limit: %s \n", gasLimit)
	log.Printf("nonce %d \n", nonce)

	toAddress := common.HexToAddress("0xe72b86f1b85565D781632E659e1e9C42f0A9EB4f")
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    toAmount,
		Gas:      gasLimit,
		GasPrice: gasPrice,
	})

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())

	// final output
	//2025/07/11 00:06:57 Transfer Amount: 0.001000000000000000
	//2025/07/11 00:06:57 gas price: 1676014
	//2025/07/11 00:06:57 gas limit: %!s(uint64=21000)
	//2025/07/11 00:06:57 nonce 0
	//2025/07/11 00:06:57 Transaction sent: 0x4ccfac3f304211aee685ae2d1a017957957c2a235978225eaed488a877da83cf
}
