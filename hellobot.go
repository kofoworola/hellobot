package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main(){
	fmt.Println("Starting Server")

	if args := os.Args; len(args) > 1 && args[1] == "-register"{
		go registerWebhook()
	}
	//Create a new Mux Handler
	m := mux.NewRouter()
	//Listen to the base url and send a response
	m.HandleFunc("/", func(writer http.ResponseWriter, _ *http.Request) {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "Server is up and running")
	})
	//Listen to crc check and handle
	m.HandleFunc("/webhook/twitter", CrcCheck).Methods("GET")

	//Start Server
	server := &http.Server{
		Handler: m,
	}
	server.Addr = ":9090"
	server.ListenAndServe()
}

func CrcCheck(writer http.ResponseWriter, request *http.Request){
	//Set response header to json type
	writer.Header().Set("Content-Type", "application/json")
	//Get crc token in parameter
	token := request.URL.Query()["crc_token"]
	if len(token) < 1 {
		fmt.Fprintf(writer,"No crc_token given")
		return
	}

	//Encrypt and encode in base 64 then return
	h := hmac.New(sha256.New, []byte("kvAUFhTPGCSYIFOR77DEVozReKRxxmlFR4tkqr3I2hPSpuaZmIkvAUFhTPGCSYIFOR77DEVozReKRxxmlFR4tkqr3I2hPSpuaZmI"))
	h.Write([]byte(token[0]))
	encoded := base64.StdEncoding.EncodeToString(h.Sum(nil))
	//Generate response string map
	response := make(map[string]string)
	response["response_token"] =  "sha256=" + encoded
	//Turn response map to json and send it to the writer
	responseJson, _ := json.Marshal(response)
	fmt.Fprintf(writer, string(responseJson))
}

