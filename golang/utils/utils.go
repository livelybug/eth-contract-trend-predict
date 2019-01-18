package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

// CryptocomparePrice ...
type CryptocomparePrice struct {
	USD  float64   `json:"USD"`
	DATE time.Time `json:"DATE"`
}

var prm sync.Map

// Usd ...
// const Usd string = "USD"
// Date ...
// const Date string = "date"
// New ...
// const New string = "new"

// Old ...
const Old string = "old"

// LogErr ...
func LogErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

// GetCryptPrice ...
func GetCryptPrice(status string) float64 {
	var cp CryptocomparePrice

	resp, err := http.Get("https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms=USD&e=Coinbase")
	LogErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	LogErr(err)
	err = json.Unmarshal(body, &cp)
	LogErr(err)

	if status == Old {
		prm.Store(Old, CryptocomparePrice{cp.USD, time.Now()})
	}

	return cp.USD
}

// ConvertToPredict ...
func ConvertToPredict(oldPr float64, newPr float64) uint {
	if oldPr < newPr {
		return 0
	} else if oldPr == newPr {
		return 1
	} else {
		return 2
	}
}

// HTTPHandler ...
func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	jsonSync := make(map[string]interface{})
	prm.Range(func(k interface{}, v interface{}) bool {
		jsonSync[k.(string)] = v
		return true
	})
	j, err := json.Marshal(&jsonSync)
	LogErr(err)
	// fmt.Printf("%s\n", j)

	fmt.Fprintf(w, string(j))
}
