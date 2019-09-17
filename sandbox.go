package tinkoff_invest_openapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

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
