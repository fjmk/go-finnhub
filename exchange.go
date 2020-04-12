package finnhub

// Exchange is the abstract data structure for exchanges
type Exchange struct {
	Name     string `json:"name"`
	Code     string `json:"omitempty,code"`
	Currency string `json:"omitempty,currency"`
}
