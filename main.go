package main

import (
	"log"
	"net/http"
)

// create a Handaler function
// Handaler is responsible for
// executing application logic and
// for writing HTTP Response header and  bodies
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello /from Snippetbox"))
}

func main() {
	//initialize a new servermux which is responsible for
	//storing a mapping between url and corrosponding handaler function
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	//use listenandservefunction to start a new web server
	//two parameters : the tcp network address to listen on a port
	///and the second marameter is servermux

	log.Println("Startiong a new seerver on Port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
