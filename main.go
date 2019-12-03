package main

import (
	"flag"
	"log"
	"net/http"
)

func main()  {

	var help bool
	var port string
	var urlPath string
	var resources string

	flag.StringVar(&port,"p","8080","server listen on port")
	flag.StringVar(&urlPath,"u","/","url path")
	flag.StringVar(&resources,"r",".","resources path")
	flag.BoolVar(&help,"h",false,"help info show this")

	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}else{
		log.Print("server listen on port:",port)
		log.Print("url path on:",urlPath)
		log.Print("resources path in path:",resources)
	}

	http.Handle(urlPath,http.FileServer(http.Dir(resources)))

	err:=http.ListenAndServe(":"+port,nil)

	if err != nil {
		log.Fatal(err)
	}


}

