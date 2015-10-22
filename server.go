package main

import (
	"crypto/tls"
	"fmt"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/sqs"
	"github.com/julienschmidt/httprouter"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	certFile = "./server.crt"
	auth     = aws.Auth{
		AccessKey: os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}
	conn      = sqs.New(auth, aws.APSoutheast)
	queueName = os.Getenv("SQS_QUEUE_PRODUCTION")
)

func GithubWebhook(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	cert, _ := tls.LoadX509KeyPair(certFile, "")
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}

	request := gorequest.New().TLSClientConfig(tlsConfig)

	url := "https://" + os.Getenv("HOST") + "/github"

	requestBody, _ := ioutil.ReadAll(req.Body)

	for key, value := range req.Header {
		fmt.Println("Key:", key, "Value:", value)
	}

	fmt.Println(string(requestBody[:]))

	queue, err := conn.GetQueue(queueName)

	resp, err := queue.SendMessage(string(requestBody))

	if err != nil {
		fmt.Println("Error sending message to queue")
	} else {
		fmt.Sprintf("Send message to queue %", resp)
	}

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
