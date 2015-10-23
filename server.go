package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/sqs"
	"github.com/julienschmidt/httprouter"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

var (
	certFile = "./server.crt"
	auth     = aws.Auth{
		AccessKey: os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}
	conn       = sqs.New(auth, aws.USWest2)
	fromSteven = regexp.MustCompile(`stevenzeiler\/rippled`)
	fromRipple = regexp.MustCompile(`ripple\/rippled`)
)

type Body struct {
	Compare string `json:"compare"`
}

func postHttp(requestBody []byte) {
	cert, _ := tls.LoadX509KeyPair(certFile, "")
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}

	request := gorequest.New().TLSClientConfig(tlsConfig)

	url := "https://" + os.Getenv("HOST") + "/github"
	request.Post(url).
		Send(string(requestBody)).
		End()
}

func GithubWebhook(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	requestBody, _ := ioutil.ReadAll(req.Body)

	fmt.Println(string(requestBody[:]))

	body := new(Body)

	if err := json.Unmarshal(requestBody, body); err != nil {
		errMsg := "error parsing json body"
		fmt.Println(errMsg)
		panic(errMsg)
	}

	var queueName string

	if fromSteven.MatchString(body.Compare) {
		queueName = os.Getenv("SQS_QUEUE_DEMO")
	} else if fromRipple.MatchString(body.Compare) {
		queueName = os.Getenv("SQS_QUEUE_PRODUCTION")
	} else {
		errMsg := "commit not from stevenzeiler or ripple!"
		fmt.Println(errMsg)
		panic(errMsg)
	}

	queue, err := conn.GetQueue(queueName)
	resp, err := queue.SendMessage(string(requestBody))

	if err != nil {
		fmt.Println("Error sending message to queue")
	} else {
		fmt.Sprintf("Send message to queue %", resp)
	}

	fmt.Fprintf(w, "Success!\n")
	go postHttp(requestBody)
}

func main() {
	router := httprouter.New()
	router.POST("/github", GithubWebhook)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
