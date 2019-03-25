package main

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/oauth1"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func CreateClient() *http.Client {
	//Create oauth client with consumer keys and access token
	config := oauth1.NewConfig("k6n3lQczhaBC3iJS7rrOQQ9qb", "kvAUFhTPGCSYIFOR77DEVozReKRxxmlFR4tkqr3I2hPSpuaZmI")
	token := oauth1.NewToken("717048519631183872-qL2loJ2abKgKfRxRLmYGQtC71SShYQx", "H18Vhs9UaspmysUrSzZXD3LtFUiCDGZoy1JnMVxiBEqml")

	return config.Client(oauth1.NoContext, token)
}
func registerWebhook(){
	fmt.Println("Registering webhook...")
	httpClient := CreateClient()

	path := "https://api.twitter.com/1.1/account_activity/all/dev/webhooks.json"
	values := url.Values{}
	values.Set("url", "https://"+os.Getenv("APP_URL")+"/twitter/webhook")
	resp, _ := httpClient.PostForm(path, values)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		panic(err)
	}
	fmt.Println("Webhook id of " + data["id"].(string) + " has been registered")
}
