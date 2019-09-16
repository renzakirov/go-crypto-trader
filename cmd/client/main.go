package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"local/go-crypto-trader/internal/app/model"
	"local/go-crypto-trader/internal/app/store"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/config.toml", "Path to config file")
}

func main() {
	// flag.Parse()
	config := store.NewConfig()

	// _, err := toml.DecodeFile(configPath, config)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	s := store.New(config)

	if err := s.Open(); err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	t := model.TradeElement{
		ID:       0,
		TradeID:  168209,
		TypeID:   1,
		Type:     "sell",
		Price:    "0.00138776",
		Quantity: "10246.1102461",
		Amount:   "14.21914195",
		Date:     1568203090,
	}
	// _, err := s.Trade().FetchOperationTypeByID(&model.OperationType{ID: 2})
	err := s.Trade().SaveTradeElement(&t)
	if err != nil {
		log.Fatal(err)
	}

	// exmo := api.New()

	// fmt.Printf("-------------\n")

	// params := model.ApiParams{"pair": "BTC_USD"}
	// _, err1 := exmo.GetTicker()
	// if err1 != nil {
	// 	fmt.Printf("api error: %s\n", err1.Error())
	// } else {
	// 	fmt.Println("api result: OK")

	// }
	// ticker, _ := getTickerFromFile()

	// BTC_USD := (*ticker)["BTC_USD"]
	// fmt.Printf("BTC_USD = %#v \n", BTC_USD)
	// fmt.Println("BTC_USD.updated = ", BTC_USD.Updated)

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
