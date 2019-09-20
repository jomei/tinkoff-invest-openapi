package tinkoff_invest_openapi

import (
	"encoding/json"
	"io/ioutil"
)

const (
	operationsUrl = url + "/operations"
)

type GetOperationsResponse struct {
	Response
	Payload struct {
		Operations []Operation `json:"operations"`
	} `json:"payload"`
}

type Operation struct {
	Id             string  `json:"id"`
	Status         string  `json:"status"`
	Trades         []Trade `json:"trades"`
	Commission     Money   `json:"commission"`
	Currency       string  `json:"currency"`
	Payment        float64 `json:"payment"`
	Price          float64 `json:"price"`
	Quantity       int32   `json:"quantity"`
	Figi           string  `json:"figi"`
	InstrumentType string  `json:"instrumentType"`
	IsMarginCall   bool    `json:"isMarginCall"`
	Date           string  `json:"date"` // todo: to date
	OperationType  string  `json:"operationType"`
}

type Trade struct {
	Id       string  `json:"tradeId"`
	Date     string  `json:"date"` // todo: to datetime
	Price    float64 `json:"price"`
	Quantity int32   `json:"quantity"`
}

func (conn *Connection) GetOperations() (*GetOperationsResponse, error) {
	resp, err := doRequest(conn, operationsUrl, "GET", nil)

	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	var r GetOperationsResponse
	err = json.Unmarshal(respBody, &r)

	if err != nil {
		return nil, err
	}

	return &r, nil
}
