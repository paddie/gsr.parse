package main

import (
	"bytes"
	"fmt"
	"time"
)

type Feed struct {
	Merchants []Merchant `xml:"merchant"`
}

type Merchant struct {
	BusinessUnitId string   `xml:"id,attr"`
	Url            string   `xml:"merchant_info>merchant_url"`
	Name           string   `xml:"merchant_info>name"`
	Reviews        []Review `xml:"review"`
}

type Review struct {
	Id         string    `xml:"id,attr"`
	ReviewerId string    `xml:"reviewer_id"`
	ReviewUrl  string    `xml:"review_url"`
	Date       time.Time `xml:"review_date"`
}

func (m *Merchant) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Merchant:\n")
	buffer.WriteString("\tName: " + m.Name + "\n")
	buffer.WriteString("\tBusinessUnitId: " + m.BusinessUnitId + "\n")
	buffer.WriteString("\tUrl: " + m.Url + "\n")

	var reviewCount = len(m.Reviews)

	cutoff := time.Now().AddDate(-1, 0, 0)
	considered := 0
	for _, r := range m.Reviews {
		if r.Date.After(cutoff) {
			considered++
		}
	}

	if considered >= 30 {
		buffer.WriteString(fmt.Sprintf("\tReviews Last 12 months: %d (total=%d)\n", considered, reviewCount))
	} else {
		buffer.WriteString(fmt.Sprintf("\tReviews Last 12 months: %d/30 (total=%d)\n", considered, reviewCount))
	}

	return buffer.String()
}

func (r *Review) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Review:\n")
	buffer.WriteString("\tId: " + r.Id + "\n")
	buffer.WriteString("\tConsumerId: " + r.ReviewerId + "\n")
	buffer.WriteString("\tUrl: " + r.ReviewUrl + "\n")

	return buffer.String()
}
