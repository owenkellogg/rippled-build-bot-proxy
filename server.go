package main

import (
	"encoding/json"
	"fmt"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/sqs"
	"github.com/julienschmidt/httprouter"
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
    fmt.Println(resp)
		fmt.Sprintf("Send message to queue %", queueName)
	}

	fmt.Fprintf(w, "Success!\n")
}

func main() {
	router := httprouter.New()
	router.POST("/github", GithubWebhook)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
