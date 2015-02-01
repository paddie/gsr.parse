package main

import (
	"encoding/xml"
	"io"
	"log"
)

type Feed_V3 struct {
	Merchants []Merchant_v3 `xml:"merchants>merchant"`
}

type Merchant_v3 struct {
	Id           string   `xml:"id,attr"`
	Name         string   `xml:"name"`
	Merchant_url string   `xml:"merchant_url"`
	Reviews      []Review `xml:"reviews>review"`
}

func (f *Feed_V3) MerchantReviews() map[string][]Review {
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

func parseFeed_V3(r io.Reader) (*Feed_V3, error) {

	feed := &Feed_V3{}

	dec := xml.NewDecoder(r)
	err := dec.Decode(feed)

	return feed, err
}
