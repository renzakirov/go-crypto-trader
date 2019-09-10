package api

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"local/go-crypto-trader/model"
)

// MarketAPI ...
type MarketAPI struct {
	Name    string
	BaseURL string
	Client  *http.Client
	key     string
	secret  string
}

// New - create new client for Market Client
func New() *MarketAPI {
	exmo := MarketAPI{
		Name:    "EXMO",
		BaseURL: "https://api.exmo.com/v1/",
		Client:  &http.Client{},
		key:     "K-5e1e4ea78b30514e440c173dd06901177b855ee7",
		secret:  "S-c0276f38c5ddede119c4acef38a1ae509a7ca0d3",
	}

	return &exmo
}

// GetTicker ...
func (ma *MarketAPI) GetTicker() (*model.TickerResponse, error) {
	method := "ticker"
	resp, err := ma.doRequest(method, nil)
	if err != nil {
		log.Println("Ошибка в doRequest ", err)
		return nil, err
	}
	_ = ioutil.WriteFile("ticker.json", *resp, 0644)

	var dat model.TickerResponse
	err2 := json.Unmarshal([]byte(*resp), &dat)
	if err2 != nil {
		log.Println("Ошибка в Unmarshal ", err2)
		return nil, err2
	}

	result := dat["result"]
	// if ok && result.(bool) != true {
	// 	return nil, errors.New(dat["error"].(string))
	// }

	// ETC_USD :=

	log.Println(result)

	return &dat, nil
}

func (ma *MarketAPI) doRequest(method string, params *model.ApiParams) (*[]byte, error) {
	postParams := url.Values{}
	postParams.Add("nonce", ma.nonce())
	if params != nil {
		for key, value := range *params {
			postParams.Add(key, value)
		}
	}
	postContent := postParams.Encode()

	sign := ma.makeSign(postContent, ma.secret)

	req, _ := http.NewRequest("POST", ma.BaseURL+method, bytes.NewBuffer([]byte(postContent)))
	req.Header.Set("Key", ma.key)
	req.Header.Set("Sign", sign)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(postContent)))

	resp, err := ma.Client.Do(req)
	if err != nil {
		log.Println("Ошибка в DO ", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		log.Println("Код возврата не 200 ", resp.Status)
		return nil, errors.New("http status: " + resp.Status)
	}

	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		log.Println("Ошибка в ReadAll ", err1)
		return nil, err1
	}

	return &body, nil
	/*
		var dat model.ApiResponse
		err2 := json.Unmarshal([]byte(body), &dat)
		if err2 != nil {
			return nil, err2
		}

		if result, ok := dat["result"]; ok && result.(bool) != true {
			return nil, errors.New(dat["error"].(string))
		}

		return &dat, nil
	*/
}

func (ma *MarketAPI) nonce() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func (ma *MarketAPI) makeSign(message string, secret string) string {
	mac := hmac.New(sha512.New, []byte(secret))
	mac.Write([]byte(message))
	return fmt.Sprintf("%x", mac.Sum(nil))
}
