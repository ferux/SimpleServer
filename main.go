package main

import (
	"flag"
	"SimpleServer/server"
	"log"
	"os"
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
	log.Println(server.Run(listenPath, assetsPath, enableDebug))
}