package tinkoff_invest_openapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	orderUrl       = url + "/orders"
	limitOrderUrl  = orderUrl + "/limit-order"
	cancelOrderUrl = orderUrl + "/cancel"
)

type OrdersResponse struct {
	Response
	Payload []Order `json:"payload"`
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

func (conn *Connection) GetOrders() (*OrdersResponse, error) {
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("GET", orderUrl, nil)

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
		log.Fatalf("Limit order, bad response code '%s' from '%s'", resp.Status, url)
		return nil, nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Can't read get order response: %s", err)
	}

	var lor LimitOrderResponse
	err = json.Unmarshal(respBody, &lor)

	if err != nil {
		log.Fatalf("Can't unmarshal limit order response: '%s' \nwith error: %s", string(respBody), err)
	}

	return &lor, nil
}

func (conn *Connection) orderCancel(orderId string) (*Response, error) {
	client := http.Client{
		Timeout: timeout,
	}

	type bodyStruct struct {
		OrderId string `json:"orderId"`
	}
	body, err := json.Marshal(bodyStruct{OrderId: orderId})

	req, err := http.NewRequest("POST", cancelOrderUrl, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer"+conn.token)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Cancel order, bad response code '%s' from '%s'", resp.Status, url)
		return nil, nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Can't read cancel order response: %s", err)
	}

	var r Response
	err = json.Unmarshal(respBody, &r)

	if err != nil {
		log.Fatalf("Can't unmarshal cancel order response: '%s' \nwith error: %s", string(respBody), err)
	}

	return &r, nil
}
