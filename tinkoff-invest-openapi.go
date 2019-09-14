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
	url         = "https://api-invest.tinkoff.ru/openapi/sandbox/sandbox"
	registerUrl = url + "/register"
	balanceUrl  = url + "/balance"
	timeout     = 10
)

type Connection struct {
	token string
}

func NewConnection(token string) *Connection {
	return &Connection{
		token: token,
	}
}

type Register struct {
	TrackingID string `json:"trackingId"`
	Status     string `json:"status"`
}

func (conn *Connection) Register() (*Register, error) {
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
		return nil, nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Can't read register response: %s", err)
	}

	var register Register
	err = json.Unmarshal(respBody, &register)

	if err != nil {
		log.Fatalf("Can't unmarshal register response: '%s' \nwith error: %s", string(respBody), err)
	}

	if strings.ToUpper(register.Status) != "OK" {
		log.Fatalf("Register failed, trackingId: '%s'", register.TrackingID)
	}

	return &register, nil
}

type Balance struct {
	TrackingID string `json:"trackingId"`
	Status     string `json:"status"`
}

func (conn *Connection) Balance(currency string, amount float64) (*Balance, error) {
	client := http.Client{
		Timeout: timeout,
	}

	type bodyStruct struct {
		Currency string  `json:"currency"`
		Amount   float64 `json:"amount"`
	}

	body, err := json.Marshal(bodyStruct{Currency: currency, Amount: amount})

	req, err := http.NewRequest("POST", balanceUrl, bytes.NewBuffer(body))

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

	var balance Balance
	err = json.Unmarshal(respBody, &balance)

	if err != nil {
		log.Fatalf("Can't unmarshal balance response: '%s' \nwith error: %s", string(respBody), err)
	}

	return &balance, nil
}
