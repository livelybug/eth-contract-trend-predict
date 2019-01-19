package mysev

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	resvbt "github.com/livelybug/eth-contract-trend-predict/golang/resvbt"
	utils "github.com/livelybug/eth-contract-trend-predict/golang/utils"
)

// Mysev ...
func Mysev() {
	dirStatic, err := filepath.Abs("golang/static")
	utils.LogErr(err)

	port := flag.String("p", getPort(), "port")
	dir := flag.String("d", dirStatic, "dir")
	flag.Parse()

	http.HandleFunc("/price-date", utils.HTTPHandler)
	http.Handle("/", http.FileServer(http.Dir(*dir)))
	log.Printf("Serving %s on Http port: %s\n", *dir, *port)
	go http.ListenAndServe(":"+*port, nil)

	for true {
		mainTimer()

		timer1 := time.NewTimer(5 * time.Second)
		<-timer1.C
	}
}

// get the Port from the environment so we can run on Heroku
func getPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "8000"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return port
}

func mainTimer() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in mainTimer()", r)
		}
	}()

	for true {
		oldPrice := utils.GetCryptPrice("old")
		log.Println("oldPrice: ", oldPrice)

		timer1 := time.NewTimer(utils.TimeOut * time.Second)
		<-timer1.C

		newPrice := utils.GetCryptPrice("new")
		log.Println("newPrice: ", newPrice)

		resvbt.ResolveBet(utils.ConvertToPredict(oldPrice, newPrice))
	}
}
