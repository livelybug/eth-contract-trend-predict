package resvbt

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	predict "./contracts/predict"
	storage "./contracts/storage"
)

const rinkebyPK string = "A2220868D5BF40CD443368CFA6C618DC3C82C14CFDF21A40E1E6B9ED049F34DB"
const contractAdd string = "0xd67caF6E55ADF71005fa0E9Cc0D2A5F7E8522EF9"
const gasPrice int = 2900000

func ResolveBet() {

	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")

	// Start verifying the deployed contract
	address := common.HexToAddress(contractAdd)
	fmt.Println(address.Hex())

	predictInstance, err := predict.NewPredict(address, client)
	if err != nil {
		log.Fatal(err)
	}

	storageNum := new(big.Int).SetInt64(0)
	storageAdd, err := predictInstance.StoreAdds(nil, storageNum)
	fmt.Println("Storage 0: ", storageAdd.Hex())

	storageInstance, err := storage.NewStorage(address, client)
	if err != nil {
		log.Fatal(err)
	}
	_ = storageInstance

	// Finish verifying the deployed contract

	poolSum, err := predictInstance.PoolSum(nil)
	if err != nil {
		log.Fatal(err)
	}
	if poolSum.Uint64() == 0 {
		log.Println("No bet yet!")
		return
	}

	privateKey, err := crypto.HexToECDSA(rinkebyPK)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	//     log.Fatal(err)
	// }

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(gasPrice) // in units
	auth.GasPrice = new(big.Int).SetUint64(auth.GasLimit)

	tx, err := predictInstance.Resolve(auth, new(big.Int).SetInt64(0))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
}
