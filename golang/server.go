package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	resvbt "./resvbt"
	utils "./utils"
)

const timeOut time.Duration = 10 // In Sec

func main() {
	port := flag.String("p", "8000", "port")
	dir := flag.String("d", "./static", "dir")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*dir)))
	log.Printf("Serving %s on Http port: %s\n", *dir, *port)
	go http.ListenAndServe(":"+*port, nil)

	for true {
		oldPrice := utils.GetCryptPrice()
		log.Println("oldPrice: ", oldPrice)

		timer1 := time.NewTimer(timeOut * time.Second)
		<-timer1.C

		newPrice := utils.GetCryptPrice()
		log.Println("newPrice: ", newPrice)

		resvbt.ResolveBet(utils.ConvertToPredict(oldPrice, newPrice))
	}
}
