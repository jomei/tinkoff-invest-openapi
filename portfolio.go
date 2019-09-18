package tinkoff_invest_openapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	portfolioUrl          = url + "/portfolio"
	portfolioCurreniesUrl = portfolioUrl + "/currencies"
)

type PortfolioResponse struct {
	Response
	Payload struct {
		Positions []Position `json:"positions"`
	} `json:"payload"`
}

type Position struct {
	Figi           string  `json:"figi"`
	Ticket         string  `json:"ticket"`
	Isin           string  `json:"isin"`
	InstrumentType string  `json:"instrumentType"`
	Balance        float64 `json:"balance"`
	Blocked        float64 `json:"blocked"`
	ExpectedYield  Money   `json:"expectedYield"`
	Lots           int32   `json:"lots"`
}

type PortfolioCurrenciesResponse struct {
	Response
	Payload struct {
		Currencies []Currency `json:"currencies"`
	} `json:"payload"`
}

type Currency struct {
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance"`
	Blocked  float64 `json:"blocked"`
}

func (conn *Connection) GetPortfolio() (*PortfolioResponse, error) {
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("GET", portfolioUrl, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer"+conn.token)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Get portfolio, bad response code '%s' from '%s'", resp.Status, url)
		return nil, nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Can't read get portfolio response: %s", err)
	}

	var pr PortfolioResponse
	err = json.Unmarshal(respBody, &pr)

	if err != nil {
		log.Fatalf("Can't unmarshal get portfolio response: '%s' \nwith error: %s", string(respBody), err)
	}

	return &pr, nil
}

func (conn *Connection) GetPortfolioCurrencies() (*PortfolioCurrenciesResponse, error) {
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("GET", portfolioCurreniesUrl, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer"+conn.token)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Get portfolio currencies, bad response code '%s' from '%s'", resp.Status, url)
		return nil, nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Can't read get portfolio curencies response: %s", err)
	}

	var pr PortfolioCurrenciesResponse
	err = json.Unmarshal(respBody, &pr)

	if err != nil {
		log.Fatalf("Can't unmarshal get portfolio response: '%s' \nwith error: %s", string(respBody), err)
	}

	return &pr, nil
}
