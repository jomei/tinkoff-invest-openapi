package tinkoff_invest_openapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	sandboxUrl           = url + "/openapi/sandbox"
	registerUrl          = sandboxUrl + "/register"
	currenciesBalanceUrl = sandboxUrl + "/currencies/balance"
	positionsBalanceUrl  = sandboxUrl + "/positions/balance"
	clearUrl             = sandboxUrl + "/clear"
)

func (conn *Connection) SandboxRegister() (*Response, error) {
	resp, err := doRequest(conn, registerUrl, "POST", nil)

	if err != nil {
		return nil, err
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
	type bodyStruct struct {
		Currency string  `json:"currency"`
		Balance  float64 `json:"balance"`
	}

	resp, err := doRequest(conn, currenciesBalanceUrl, "POST", bodyStruct{Currency: currency, Balance: balance})

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
	type bodyStruct struct {
		Balance float64 `json:"balance"`
		Figi    string  `json:"figi"`
	}

	resp, err := doRequest(conn, positionsBalanceUrl, "POST", bodyStruct{Figi: figi, Balance: balance})

	if err != nil {
		return nil, err
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
	resp, err := doRequest(conn, clearUrl, "POST", nil)

	if err != nil {
		return nil, err
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
