package finnhub

//CompanyExecutives is the response for /stock/executive
// https://finnhub.io/docs/api#company-executive
type CompanyExecutives struct {
	Executive []CompanyExecutive `json:"executive"`
}

// CompanyExecutive data structure for company executives
type CompanyExecutive struct {
	Name         string  `json:"name"`
	Title        string  `json:"title"`
	Since        string  `json:"since"`
	Sex          string  `json:"sex"`
	Compensation float64 `json:"compensation"`
	Currency     string  `json:"currency"`
}
