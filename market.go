package tinkoff_invest_openapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const (
	marketUrl        = url + "/market"
	getStocksUrl     = marketUrl + "/stocks"
	getBondsUrl      = marketUrl + "/bonds"
	getCurrenciesUrl = marketUrl + "/currencies"
	getEtfsUrl       = marketUrl + "/etfs"
	searchUrl        = marketUrl + "/search"
	getByFigiUrl     = searchUrl + "/by-figi"
)

type InstrumentsResponse struct {
	Response
	Payload struct {
		Total       float64      `json:"total"`
		Instruments []Instrument `json:"instruments"`
	} `json:"payload"`
}

type Instrument struct {
	Figi              string  `json:"figi"`
	Ticker            string  `json:"ticker"`
	Isin              string  `json:"isin"`
	MinPriceIncrement float64 `json:"minPriceIncrement"`
	Lot               int32   `json:"lot"`
	Currency          string  `json:"currency"`
}

type GetByFigiResponse struct {
	Response
	Payload Instrument `json:"payload"`
}

func doMarketRequest(conn *Connection, url string, requestType string) (*InstrumentsResponse, error) {
	resp, err := doRequest(conn, url, "GET", nil)

	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Can't read %s response: %s", requestType, err)
	}

	var r InstrumentsResponse
	err = json.Unmarshal(respBody, &r)

	if err != nil {
		log.Fatalf("Can't unmarshal %s response: '%s' \nwith error: %s", requestType, string(respBody), err)
	}

	return &r, nil
}

func (conn *Connection) GetStocks() (*InstrumentsResponse, error) {
	return doMarketRequest(conn, getStocksUrl, "get stocks")
}

func (conn *Connection) GetBonds() (*InstrumentsResponse, error) {
	return doMarketRequest(conn, getBondsUrl, "get bonds")
}

func (conn *Connection) GetCurrencies() (*InstrumentsResponse, error) {
	return doMarketRequest(conn, getCurrenciesUrl, "get currencies")
}

func (conn *Connection) GetEtfs() (*InstrumentsResponse, error) {
	return doMarketRequest(conn, getEtfsUrl, "get etfs")
}

func (conn *Connection) GetByFigi(figi string) (*GetByFigiResponse, error) {
	type body struct {
		Figi string `json:"figi"`
	}
	resp, err := doRequest(conn, getByFigiUrl, "GET", body{Figi: figi})

	if err != nil {
		return nil, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	var r GetByFigiResponse
	err = json.Unmarshal(respBody, &r)

	if err != nil {
		log.Fatalf("Can't unmarshal %s response: '%s' \nwith error: %s", "get by figi", string(respBody), err)
	}

	return &r, nil

}
