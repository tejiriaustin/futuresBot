package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type BinanceOpenSpotRequests struct {
	Symbols       string `json:"symbols"`
	OrderId       int    `json:"orderId"`
	OrderListId   int    `json:"orderListId"`
	ClientOrderId string `json:"clientOrderId"`
	TransactTime  int    `json:"transactTime"`
}

type BinanceOpenSpotResponse struct {
	Code                    int     `json:"code"`
	Symbols                 string  `json:"symbols"`
	OrderId                 int     `json:"orderId"`
	OrderListId             int     `json:"orderListId"`
	ClientOrderId           string  `json:"clientOrderId"`
	TransactTime            int     `json:"transactTime"`
	Price                   float64 `json:"price"`
	OriginQty               int     `json:"origQty"`
	ExecutedQty             float64 `json:"executedQty"`
	CummulativeQuoteQty     float64 `json:"cummulativeQuoteQty"`
	Status                  string  `json:"status"`
	TimeInForce             string  `json:"timeInForce"`
	Type                    string  `json:"type"`
	Side                    string  `json:"side"`
	StategyId               string  `json:"stategyId"`
	StategyType             string  `json:"stategyType"`
	WorkingTime             string  `json:"workingTime"`
	SelfTradePreventionMode string  `json:"selfTradePreventionMode"`
}

type Requests struct {
	Req *http.Request
}

func NewRequests(req *http.Request) *Requests {
	return &Requests{
		Req: req,
	}
}

func main() {
	fmt.Println("Hello World!")
	orderId := 1
	for {

		BuyBtc(orderId)

		time.Sleep(2 * time.Second)

		orderId++
	}
}

func NewBuyBtcRequest(req BinanceOpenSpotRequests) ([]byte, error) {
	return json.Marshal(req)
}

func NewBuyBtcResponse(res []byte) (BinanceOpenSpotResponse, error) {
	var response BinanceOpenSpotResponse
	err := json.Unmarshal(res, &response)
	if err != nil {
		return BinanceOpenSpotResponse{}, err
	}
	return response, nil
}

func PostRequest(url string, requestBody []byte) *http.Request {
	body := bytes.NewReader(requestBody)
	var b []byte
	_, err := body.Read(b)
	if err != nil {
		fmt.Println("fgdcjknlm;njhgvcjbknlmjbcghxjjk")
		return nil
	}
	fmt.Println(string(b))
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil
	}
	return request
}

func BuyBtc(OrderId int) {

	fmt.Println("buy btc")

	url := "https://api3.binance.com/api/v3/order"
	requestBody, err := NewBuyBtcRequest(
		BinanceOpenSpotRequests{
			Symbols:       "BTCUSDT",
			OrderId:       OrderId,
			OrderListId:   -1,
			ClientOrderId: "my_client_id",
		},
	)
	if err != nil {
		return
	}

	post := PostRequest(url, requestBody)
	r := NewRequests(post)

	postRequests, err := r.PostRequest()
	if err != nil {
		log.Println(err)
		return
	}

	body, err := ioutil.ReadAll(postRequests.Body)
	if err != nil {
		log.Println(err)
		return
	}

	response, err := NewBuyBtcResponse(body)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("buy btc end")
	log.Println(response)
	return
}

func (r *Requests) PostRequest() (*http.Response, error) {

	c := http.Client{}

	head := NewHeaderObj(r.Req.Header)

	head.
		SetHeader("Accept", "application/Json")

	response, err := c.Do(r.Req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return response, nil
}

type HeaderObj struct {
	head http.Header
}

func NewHeaderObj(head http.Header) *HeaderObj {
	return &HeaderObj{
		head: head,
	}
}

func (h HeaderObj) SetHeader(key string, value string) HeaderObj {
	h.head.Add(key, value)
	return h
}

func (h HeaderObj) GetHeader() http.Header {
	return h.head
}
