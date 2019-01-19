package mysev

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"time"

	resvbt "./resvbt"
	utils "./utils"
)

// Mysev ...
func Mysev() {
	dirStatic, err := filepath.Abs("golang/static")
	utils.LogErr(err)

	port := flag.String("p", "8000", "port")
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
