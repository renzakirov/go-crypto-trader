package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"local/go-crypto-trader/model"
	"log"
)

func main() {
	// exmo := api.New()

	fmt.Printf("-------------\n")

	// params := model.ApiParams{"pair": "BTC_USD"}
	// _, err1 := exmo.GetTicker()
	// if err1 != nil {
	// 	fmt.Printf("api error: %s\n", err1.Error())
	// } else {
	// 	fmt.Println("api result: OK")

	// }
	ticker, _ := getTickerFromFile()

	BTC_USD := (*ticker)["BTC_USD"]
	fmt.Printf("BTC_USD = %#v \n", BTC_USD)
	fmt.Println("BTC_USD.updated = ", BTC_USD.Updated)

}

func getTickerFromFile() (*model.TickerResponse, error) {
	var result model.TickerResponse
	byteValue, err := ioutil.ReadFile("ticker.json")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &result, nil
}
