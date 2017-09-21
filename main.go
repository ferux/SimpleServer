package main

import (
	"net"
	"net/http"
	"log"
	"flag"
	"path/filepath"
	"os"
	"time"
)

func usage() {
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	var assetsPath,	listenPath string
	var enableDebug, help bool
	flag.Usage = usage
	
	flag.StringVar(&listenPath, "l", ":8080", "HTTP listen spec")
	flag.StringVar(&assetsPath, "a", "./", "Path to assets dir")
	flag.BoolVar(&enableDebug, "d", false, "Debug mode")
	flag.BoolVar(&help, "h", false, "Usage help")
	flag.Parse()	
	if help {
		usage()
	}
	serv := http.Server {
		ReadTimeout: 60 * time.Second,
		WriteTimeout: 60 * time.Second,
		MaxHeaderBytes: 1 << 16,
	}

	listener, err := net.Listen("tcp", listenPath)
	if err != nil {
		log.Fatalf("Error creating listener. Reason^ %v", err)
	}
	
	if enableDebug {
		logger := log.New(os.Stdout, "httpSimpleServer", log.Lshortfile)
		serv.ErrorLog = logger
		serv.ConnState = checkConnState
	}

	if _, err := os.Stat(assetsPath); os.IsNotExist(err) {
		log.Println("You should specify an existing directory.")
		os.Exit(2)
	}

	http.Handle("/", http.FileServer(http.Dir(assetsPath)))
	absPath, err := filepath.Abs(assetsPath)
	if err != nil {
		absPath = assetsPath
	}
	log.Printf("\nWelcome to simple HTTP server!\nServer is available on %s\nPut your static files here: %s\nDebug mode: %v", listenPath, absPath, enableDebug)
	go serv.Serve(listener)
	log.Println("Hit enter to exit")
	buf := make([]byte, 1)
	os.Stdin.Read(buf)
	log.Println("Got exit signal. Terminating...")
}

func checkConnState (conn net.Conn, cState http.ConnState) {
	if cState == http.StateNew {
		log.Printf("New connection from %v", conn.RemoteAddr())	
	}
}
