package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	resvbt "./resvbt"
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
		timer1 := time.NewTimer(timeOut * time.Second)
		<-timer1.C
		log.Println("Timer 1 expired")
		resvbt.ResolveBet()
	}
}
