package main

import (
	"crypto/tls"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	certFile = "./server.crt"
)

func GithubWebhook(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	cert, _ := tls.LoadX509KeyPair(certFile, "")
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}

	request := gorequest.New().TLSClientConfig(tlsConfig)

	url := "https://"+os.Getenv("HOST")+"/github"

	requestBody, _ := ioutil.ReadAll(req.Body)

  fmt.Println(requestBody)

	request.Post(url).
		Send(string(requestBody)).
		End()

	fmt.Fprintf(w, "Success!\n")
}

func main() {
	router := httprouter.New()
	router.POST("/github", GithubWebhook)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
