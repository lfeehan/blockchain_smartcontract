 package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TODO: Automate this
const key = `{"address":"3590aca93338b0721966a8d0c96ebf2c4c87c544","crypto":{"cipher":"aes-128-ctr","ciphertext":"fb2deb799d1bdfb91aeecdfad15041a08ae759e6f104ea4cdbedbc8b742a5642","cipherparams":{"iv":"0700cd40c0ee2b706999ad5a3796fb04"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"cb8d5b3dc61b2d60092e6a79f83df894af1897bda5f966668d2e3cb3ac97c5a9"},"mac":"3785012c4140280a49078f0003663ea4b4c9971a7270056174ffbe529f238fad"},"id":"e0704e02-8a46-4ca6-a6b0-5152db4dfb52","version":3}`

func main() {
	Deploy()
}

func Deploy(){
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	provider := os.Getenv("PROVIDER_URI")
	account_pw := os.Getenv("PASSWORD")
	fmt.Printf("Deploying Smart Contract to: %s", provider)

	conn, err := ethclient.Dial(provider)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	_ = conn
	_ = err
	auth, err := bind.NewTransactor(strings.NewReader(key), account_pw)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	_ = auth

	// Deploy a new contract for the binding demo
	// method signature Deploy[contractname](auth, conn, **Deploy args)
	address, tx, token, err := DeployInbox(auth, conn, "test message")
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	_ = token

	// Don't even wait, check its presence in the local pending state
	time.Sleep(250 * time.Millisecond) // Allow it to be processed by the local node :P
}



