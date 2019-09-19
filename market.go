package tinkoff_invest_openapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	marketUrl        = url + "/market"
	getStocksUrl     = marketUrl + "/stocks"
	getBondsUrl      = marketUrl + "/bonds"
	getCurrenciesUrl = marketUrl + "/bonds"
)

type MarketResponse struct {
	Response
	Payload struct {
		Total       float64      `json:"total"`
		Instruments []Instrument `json:"instruments"`
	} `json:"payload"`
}

type Instrument struct {
	Figi              string  `json:"figi"`
	Ticket            string  `json:"ticket"`
	Isin              string  `json:"isin"`
	MinPriceIncrement float64 `json:"minPriceIncrement"`
	Lot               int32   `json:"lot"`
	Currency          string  `json:"currency"`
}

func doMarkerRequest(conn *Connection, url string, requestType string) (*MarketResponse, error) {
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("GET", getStocksUrl, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer"+conn.token)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("%s, bad response code '%s' from '%s'", requestType, resp.Status, url)
		return nil, nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Can't read %s response: %s", requestType, err)
	}

	var r MarketResponse
	err = json.Unmarshal(respBody, &r)

	if err != nil {
		log.Fatalf("Can't unmarshal %s response: '%s' \nwith error: %s", requestType, string(respBody), err)
	}

	return &r, nil
}

func (conn *Connection) GetStocks() (*MarketResponse, error) {
	return doMarkerRequest(conn, getStocksUrl, "get stocks")
}

func (conn *Connection) GetBonds() (*MarketResponse, error) {
	return doMarkerRequest(conn, getBondsUrl, "get bonds")
}

func (conn *Connection) GetCurrencies() (*MarketResponse, error) {
	return doMarkerRequest(conn, getCurrenciesUrl, "get currencies")
}
