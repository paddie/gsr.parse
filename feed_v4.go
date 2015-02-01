package main

import (
	"encoding/xml"
	"io"
	"log"
)

type Feed_V4 struct {
	Merchants []Merchant_v4 `xml:"merchant"`
}

func (f *Feed_V4) MerchantReviews() map[string][]Review {
	merchs := make(map[string][]Review)
	for _, m := range f.Merchants {
		if _, ok := merchs[m.Id]; ok {
			log.Printf("ISSUE: Duplicate Merchant Id = %s", m.Id)
			continue
		} else {
			merchs[m.Id] = m.Reviews
		}
	}

	return merchs
}

type Merchant_v4 struct {
	Id           string   `xml:"id,attr"`
	Name         string   `xml:"merchant_info>name"`
	Merchant_url string   `xml:"merchant_url"`
	Reviews      []Review `xml:"review"`
}

func parseFeed_V4(r io.Reader) (*Feed_V4, error) {

	feed := &Feed_V4{}

	dec := xml.NewDecoder(r)
	err := dec.Decode(feed)

	return feed, err
}
