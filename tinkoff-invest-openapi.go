package tinkoff_invest_openapi

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

type Money struct {
	Currency string  `json:"currency"`
	Value    float64 `json:"value"`
}
