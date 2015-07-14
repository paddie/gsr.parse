package main

import (
	"bytes"
	"fmt"
)

type Feed map[string]Merchant

type Merchant struct {
	BusinessUnitId string
	Url            string
	Name           string
	Reviews        []Review
}

func (m *Merchant) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Merchant:\n")
	buffer.WriteString("\tName: " + m.Name + "\n")
	buffer.WriteString("\tBusinessUnitId: " + m.BusinessUnitId + "\n")
	buffer.WriteString("\tUrl: " + m.Url + "\n")
	buffer.WriteString(fmt.Sprintf("\tReviews: %d\n", len(m.Reviews)))

	return buffer.String()
}

type Review struct {
	Id         string `xml:"id,attr"`
	ReviewerId string `xml:"reviewer_id"`
	ReviewUrl  string `xml:"review_url"`
}

func (r *Review) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Review:\n")
	buffer.WriteString("\tId: " + r.Id + "\n")
	buffer.WriteString("\tConsumerId: " + r.ReviewerId + "\n")
	buffer.WriteString("\tUrl: " + r.ReviewUrl + "\n")

	return buffer.String()
}
