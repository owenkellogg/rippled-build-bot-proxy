package main

import(
  "github.com/julienschmidt/httprouter"
  "github.com/parnurzeal/gorequest"
  "net/http"
  "crypto/tls"
  "fmt"
  "log"
  "io/ioutil"
  "os"
)

var (
  certFile = "./server.crt"
) 

func GithubWebhook(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
  
  cert, _ := tls.LoadX509KeyPair(certFile, "")
  tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}

  request := gorequest.New().TLSClientConfig(tlsConfig)

  url := "https://54.254.176.227/github"

  requestBody, _ := ioutil.ReadAll(req.Body)

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

