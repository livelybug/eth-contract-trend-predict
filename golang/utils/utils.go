package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// CryptocomparePrice ...
type CryptocomparePrice struct {
	USD float64 `json:"USD"`
}

// LogErr ...
func LogErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

// GetCryptPrice ...
func GetCryptPrice() float64 {
	var cp CryptocomparePrice

	resp, err := http.Get("https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms=USD&e=Coinbase")
	LogErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	LogErr(err)
	err = json.Unmarshal(body, &cp)
	LogErr(err)
	return cp.USD
}

// ConvertToPredict ...
func ConvertToPredict(oldPr float64, newPr float64) uint {
	if(oldPr < newPr) {
		return 0
	} else if(oldPr == newPr) {
		return 1
	} else {
		return 2
	}
}