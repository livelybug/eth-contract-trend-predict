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

	utils "github.com/livelybug/eth-contract-trend-predict/golang/utils"
	predict "github.com/livelybug/eth-contract-trend-predict/golang/resvbt/contracts/predict"
	storage "github.com/livelybug/eth-contract-trend-predict/golang/resvbt/contracts/storage"
)

const rinkebyPK string = "A2220868D5BF40CD443368CFA6C618DC3C82C14CFDF21A40E1E6B9ED049F34DB"
const contractAdd string = "0xd67caF6E55ADF71005fa0E9Cc0D2A5F7E8522EF9"
const gasPrice int = 3000000
const gasLimit int = 3900000

// ResolveBet ... Finalize current bet
func ResolveBet(trend uint) {

	client, err := ethclient.Dial("https://rinkeby.infura.io")
	utils.LogErr(err)

	fmt.Println("we have a connection")

	// Start verifying the deployed contract
	address := common.HexToAddress(contractAdd)
	fmt.Println(address.Hex())

	predictInstance, err := predict.NewPredict(address, client)
	utils.LogErr(err)

	storageNum := new(big.Int).SetInt64(0)
	storageAdd, err := predictInstance.StoreAdds(nil, storageNum)
	fmt.Println("Storage 0: ", storageAdd.Hex())

	storageInstance, err := storage.NewStorage(address, client)
	utils.LogErr(err)
	_ = storageInstance

	// Finish verifying the deployed contract

	poolSum, err := predictInstance.PoolSum(nil)
	utils.LogErr(err)
	if poolSum.Uint64() == 0 {
		log.Println("No bet yet!")
		return
	}

	privateKey, err := crypto.HexToECDSA(rinkebyPK)
	utils.LogErr(err)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	utils.LogErr(err)

	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// utils.LogErr(err)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(gasLimit) // in units
	auth.GasPrice = new(big.Int).SetUint64(uint64(gasPrice))

	tx, err := predictInstance.Resolve(auth, big.NewInt(int64(trend)))
	utils.LogErr(err)

	fmt.Println("tx sent: ", tx.Hash().Hex())
}
