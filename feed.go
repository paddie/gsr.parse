package main

type Feed map[string]Merchant

type Merchant struct {
	BusinessUnitId string
	Url            string
	Name           string
	Reviews        []Review
}

type Review struct {
	Id         string `xml:"id,attr"`
	ReviewerId string `xml:"reviewer_id"`
	ReviewUrl  string `xml:"review_url"`
}
