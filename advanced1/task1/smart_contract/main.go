package main

import (
	"context"
	"github.com/caryxiao/learning-go/advanced1/task1/smart_contract/contracts/simple_storage"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

const contractAddr = "0xA716167B3147a6717fa8d3B70954372689dF8592"

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/API_KEY")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("private_key")
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return
	}

	simpleStorageContract, err := simple_storage.NewSimpleStorage(common.HexToAddress(contractAddr), client)

	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := simpleStorageContract.Store(opt, big.NewInt(168))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("stored tx hash: %v \n", tx.Hash().Hex())
	callOpt := &bind.CallOpts{Context: context.Background()}

	storageNumber, err := simpleStorageContract.Retrieve(callOpt)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("stored storage number: %v \n", storageNumber.String())

}
