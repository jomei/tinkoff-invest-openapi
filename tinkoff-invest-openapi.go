package tinkoff_invest_openapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	url                  = "https://api-invest.tinkoff.ru"
	sandboxUrl           = url + "/openapi/sandbox"
	registerUrl          = sandboxUrl + "/register"
	currenciesBalanceUrl = sandboxUrl + "/currencies/balance"
	positionsBalanceUrl  = sandboxUrl + "/positions/balance"
	clearUrl             = sandboxUrl + "/clear"
	orderUrl             = url + "/orders"
	limitOrderUrl        = orderUrl + "/limit-order"
	timeout              = 10
)

type Connection struct {
	token string
}

func NewConnection(token string) *Connection {
	return &Connection{
		token: token,
	}
}

type Response struct {
	TrackingID string `json:"trackingId"`
	Status     string `json:"status"`
}

type OrdersResponse struct {
	Response
	Payload []*Order `json:"payload"`
}

type Order struct {
	OrderId       string  `json:"orderId"`
	Figi          string  `json:"figi"`
	Operation     string  `json:"operation"`
	Status        string  `json:"status"`
	RequestedLots int32   `json:"requestedLots"`
	ExecutedLots  int32   `json:"executedLots"`
	Type          string  `json:"type"`
	Price         float64 `json:"prise"`
}

type Money struct {
	Currency string  `json:"currency"`
	Value    float64 `json:"value"`
}

type LimitOrderResponse struct {
	Response
	Payload []*LimitOrder `json:"payload"`
}

type LimitOrder struct {
	OrderId       string `json:"orderId"`
	Operation     string `json:"operation"`
	Status        string `json:"status"`
	RejectReason  string `json:"rejectReason"`
	RequestedLots int32  `json:"requestedLots"`
	ExecutedLots  int32  `json:"executedLots"`
	Commission    Money  `json:"commission"`
}

func (conn *Connection) SandboxRegister() (*Response, error) {
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("POST", registerUrl, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer"+conn.token)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Register, bad response code '%s' from '%s'", resp.Status, url)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Can't read register response: %s", err)
	}

	var register Response
	err = json.Unmarshal(respBody, &register)

	if err != nil {
		log.Fatalf("Can't unmarshal register response: '%s' \nwith error: %s", string(respBody), err)
	}

	if strings.ToUpper(register.Status) != "OK" {
		log.Fatalf("Register failed, trackingId: '%s'", register.TrackingID)
	}

	return &register, nil
}

func (conn *Connection) SandboxCurrencyBalance(currency string, balance float64) (*Response, error) {
	client := http.Client{
		Timeout: timeout,
	}

	type bodyStruct struct {
		Currency string  `json:"currency"`
		Balance  float64 `json:"balance"`
	}

	body, err := json.Marshal(bodyStruct{Currency: currency, Balance: balance})

	req, err := http.NewRequest("POST", currenciesBalanceUrl, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer"+conn.token)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Balance, bad response code '%s' from '%s'", resp.Status, url)
		return nil, nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Can't read balance response: %s", err)
	}

	var b Response
	err = json.Unmarshal(respBody, &balance)

	if err != nil {
		log.Fatalf("Can't unmarshal balance response: '%s' \nwith error: %s", string(respBody), err)
	}

	return &b, nil
}

func (conn *Connection) SandboxPositionBalance(figi string, balance float64) (*Response, error) {
	client := http.Client{
		Timeout: timeout,
	}

	type bodyStruct struct {
		Balance float64 `json:"balance"`
		Figi    string  `json:"figi"`
	}

	body, err := json.Marshal(bodyStruct{Figi: figi, Balance: balance})

	req, err := http.NewRequest("POST", positionsBalanceUrl, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer"+conn.token)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Balance, bad response code '%s' from '%s'", resp.Status, url)
		return nil, nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Can't read balance response: %s", err)
	}

	var b Response
	err = json.Unmarshal(respBody, &balance)

	if err != nil {
		log.Fatalf("Can't unmarshal balance response: '%s' \nwith error: %s", string(respBody), err)
	}

	return &b, nil
}

func (conn *Connection) SandboxClear() (*Response, error) {
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("POST", clearUrl, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer"+conn.token)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Clear, bad response code '%s' from '%s'", resp.Status, url)
		return nil, nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Can't read clear response: %s", err)
	}

	var clear Response
	err = json.Unmarshal(respBody, &resp)

	if err != nil {
		log.Fatalf("Can't unmarshal clear response: '%s' \nwith error: %s", string(respBody), err)
	}

	if strings.ToUpper(resp.Status) != "OK" {
		log.Fatalf("Clear failed, trackingId: '%s'", clear.TrackingID)
	}

	return &clear, nil
}

func (conn *Connection) GetOrders() (*OrdersResponse, error) {
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("POST", orderUrl, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer"+conn.token)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Get order, bad response code '%s' from '%s'", resp.Status, url)
		return nil, nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Can't read get order response: %s", err)
	}

	var or OrdersResponse
	err = json.Unmarshal(respBody, &or)

	if err != nil {
		log.Fatalf("Can't unmarshal get order response: '%s' \nwith error: %s", string(respBody), err)
	}

	return &or, nil
}

func (conn *Connection) limitOrder(figi string, lots int32, operation string, price float64) (*LimitOrderResponse, error) {
	client := http.Client{
		Timeout: timeout,
	}

	type bodyStruct struct {
		Figi      string  `json:"figi"`
		Lots      int32   `json:"lots"`
		Operation string  `json:"string"`
		Price     float64 `json:"price"`
	}
	body, err := json.Marshal(bodyStruct{Figi: figi, Lots: lots, Operation: operation, Price: price})

	req, err := http.NewRequest("POST", limitOrderUrl, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer"+conn.token)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Get order, bad response code '%s' from '%s'", resp.Status, url)
		return nil, nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Can't read get order response: %s", err)
	}

	var lor LimitOrderResponse
	err = json.Unmarshal(respBody, &lor)

	if err != nil {
		log.Fatalf("Can't unmarshal get order response: '%s' \nwith error: %s", string(respBody), err)
	}

	return &lor, nil
}
