package main

import (
	"bytes"
	"time"
)

type Review struct {
	Id         string    `xml:"id,attr"`
	ReviewerId string    `xml:"reviewer_id"`
	ReviewUrl  string    `xml:"review_url"`
	Date       time.Time `xml:"review_date"`
}

func (r *Review) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Review:\n")
	buffer.WriteString("\tId: " + r.Id + "\n")
	buffer.WriteString("\tConsumerId: " + r.ReviewerId + "\n")
	buffer.WriteString("\tUrl: " + r.ReviewUrl + "\n")

	return buffer.String()
}
