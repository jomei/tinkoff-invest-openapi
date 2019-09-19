package tinkoff_invest_openapi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	url     = "https://api-invest.tinkoff.ru"
	timeout = 10
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

type ErrorResponse struct {
	Response
	Payload struct {
		Message string `json:"message"`
		Code    string `json:"code"`
	} `json:"payload"`
}

type Money struct {
	Currency string  `json:"currency"`
	Value    float64 `json:"value"`
}

func doRequest(conn *Connection, url string, method string, requestBody interface{}) (*http.Response, error) {
	client := http.Client{
		Timeout: timeout,
	}

	body, err := json.Marshal(requestBody)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer"+conn.token)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, nil // todo: fix
	}

	return resp, nil
}
