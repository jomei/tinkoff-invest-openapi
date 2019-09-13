package tinkoff_invest_openapi

import "testing"

func TestNew(t *testing.T) {
	token := "someToken"
	connection := NewConnection(token)
	if token != connection.token {
		t.Errorf("token mismatch")
	}
}
