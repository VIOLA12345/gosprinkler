package main

import (
	"fmt"
	//"html "
	"log"
	"net/http"
	"gobot.io/x/gobot/platforms/raspi"
)

var pi *raspi.Adaptor

func main() {

	pi = raspi.NewAdaptor()

	http.HandleFunc("/sprinkler/on", sprinklerOn)

	http.HandleFunc("/sprinkler/off", sprinklerOff)

	static := http.FileServer(http.Dir("../webcontent"))
	http.Handle("/", static)

	log.Printf("About to listen on 8081 to http://localhost:8081/")
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func sprinklerOn(w http.ResponseWriter, r *http.Request) {
	whichSprinkler := r.URL.Query()["which"]
	fmt.Fprintf(w, "Will attempt to turn on  %v", whichSprinkler)

	switch whichSprinkler[0] {
	case "A":
		pi.DigitalWrite("37", 0)
	case "B":
		pi.DigitalWrite("33", 1)
	case "C":
		pi.DigitalWrite("35", 1)
	case "D":
		pi.DigitalWrite("37", 1)
	default:
		fmt.Fprintf(w, "Invalid name for sprinkler  %v", whichSprinkler)
		return
	}

	fmt.Fprintf(w, "Sprinkler %v is turned on", whichSprinkler)
}

func sprinklerOff(w http.ResponseWriter, r *http.Request) {
	whichSprinkler := r.URL.Query()["which"]
	fmt.Fprintf(w, "Will attempt to turn off  %v", whichSprinkler)

	switch whichSprinkler[0] {
	case "A":
		pi.DigitalWrite("37", 1)
	case "B":
		pi.DigitalWrite("33", 0)
	case "C":
		pi.DigitalWrite("35", 0)
	case "D":
		pi.DigitalWrite("37", 0)
	default:
		fmt.Fprintf(w, "Invalid name for sprinkler  %v", whichSprinkler)
		return
	}

	fmt.Fprintf(w, "Sprinkler %v is turned off", whichSprinkler)
}
