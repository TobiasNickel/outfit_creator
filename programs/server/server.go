package main

import (
	"flag"

	"github.com/TobiasNickel/outfit_creator/frontend"
	"github.com/TobiasNickel/outfit_creator/_internal/server"
)

func main() {
	// option parsing
	port := flag.String("port", "8080", "port number to listen")
	developMode := flag.Bool(
		"develop",
		false,
		"in development use the frontend from FS(not the embedded) the path is ./frontend/public",
	)
	help := flag.Bool("help", false, "show help")
	flag.Parse()
	if *help == true {
		flag.PrintDefaults()
		return
	}

	staticfs := frontend.BuildHTTPFS()
	if *developMode == true {
		staticfs = frontend.BuildHTTPDevelopmentFS()
	}

	r := server.New(staticfs)

	// server start
	r.Run("0.0.0.0:" + *port)
}
