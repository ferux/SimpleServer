package server

import (
	"net"
	"net/http"
	"log"
	"os"
	"time"
)

//Run function runs the server for listening on specified ip:port.
func Run(addr, assets string, debug bool) error {
	serv := http.Server {
		ReadTimeout: 60 * time.Second,
		WriteTimeout: 60 * time.Second,
		MaxHeaderBytes: 1 << 16,
	}

	if debug {
		logger := log.New(os.Stdout, "httpSimpleServer", log.Lshortfile)
		serv.ErrorLog = logger
		serv.ConnState = checkConnState
	}


	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	if _, err := os.Stat(assets); os.IsNotExist(err) {
		return err
	}

	http.Handle("/", http.FileServer(http.Dir(assets)))
	go serv.Serve(listener)
	defer serv.Close()
	log.Printf("\nWelcome to simple HTTP server!\nServer is available on %s\nPut your static files here: %s\nDebug mode: %v", addr, assets, debug)
	log.Println("Hit enter to exit")
	buf := make([]byte, 1)
	os.Stdin.Read(buf)
	return nil
}

func checkConnState (conn net.Conn, cState http.ConnState) {
	if cState == http.StateNew {
		log.Printf("New connection from %v", conn.RemoteAddr())	
	}
}
