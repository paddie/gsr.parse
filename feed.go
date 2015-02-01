package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

type Review struct {
	Id         string `xml:"id,attr"`
	ReviewerId string `xml:"reviewer_id"`
}

type Feeder interface {
	MerchantReviews() map[string][]Review
}

func ParseFeed(r io.Reader) (Feeder, error) {

	rd := bufio.NewReader(r)

	b, err := rd.Peek(600)
	if err != nil {
		return nil, err
	}

	header := string(b)

	if strings.Contains(header, "3.0/merchant_reviews.xsd") {
		log.Println("Schema version 3.0 detected")
		return parseFeed_V3(rd)
	} else if strings.Contains(header, "4.0/merchant_reviews.xsd") {
		log.Println("Schema version 4.0 detected")
		return parseFeed_V4(rd)
	}

	return nil, fmt.Errorf("Unable to determine Schema version in header '%s'", header)
}
