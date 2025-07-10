package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

func main() {
	dial, err := ethclient.Dial("xxxx")
	if err != nil {
		return
	}

	blockNumber := big.NewInt(8721075)
	header, err := dial.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(header.Number().String())
	fmt.Println(header.Hash().String())
	fmt.Println(header.Time())
	fmt.Println(header.ParentHash().String())
	fmt.Println(header.BaseFee().String())
	fmt.Println(header.Difficulty().String())
	fmt.Println(header.GasLimit())
	fmt.Println(header.GasUsed())
	fmt.Println(len(header.Transactions()))

	account := common.HexToAddress("0xe72b86f1b85565D781632E659e1e9C42f0A9EB4f")
	balance, err := dial.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balance: ", balance.String())
	formatBalance := new(big.Float).Quo(new(big.Float).SetInt(balance), new(big.Float).SetFloat64(1e18))
	formatBalance2 := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(math.Pow10(18)))
	fmt.Println("FormatBalance: ", formatBalance.String())
	fmt.Println("FormatBalance2: ", formatBalance2.String())

}
