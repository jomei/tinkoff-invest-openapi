package tinkoff_invest_openapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
	resp, err := doRequest(conn, orderUrl, "GET", nil)

	if err != nil {
		return nil, err
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
	type bodyStruct struct {
		Figi      string  `json:"figi"`
		Lots      int32   `json:"lots"`
		Operation string  `json:"string"`
		Price     float64 `json:"price"`
	}

	resp, err := doRequest(conn, limitOrderUrl, "POST", bodyStruct{Figi: figi, Lots: lots, Operation: operation, Price: price})

	if err != nil {
		return nil, err
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
	type bodyStruct struct {
		OrderId string `json:"orderId"`
	}

	resp, err := doRequest(conn, cancelOrderUrl, "POST", bodyStruct{OrderId: orderId})

	if err != nil {
		return nil, err
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
