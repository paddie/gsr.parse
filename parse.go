package main

import (
	"encoding/xml"
	"io"
)

func findReviewId(reader io.Reader, reviewId string) (*Review, error) {

	decoder := xml.NewDecoder(reader)

	for {

		t, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				return nil, ReviewNotFound
			}
			return nil, err
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local != "review" {
				continue
			}

			for _, attr := range se.Attr {
				if attr.Name.Local == "id" && attr.Value == reviewId {
					review := &Review{}
					err = decoder.DecodeElement(review, &se)
					return review, err
				}
			}
		}
	}

	return nil, ReviewNotFound
}

func findMerchantId(reader io.Reader, merchantId string) (*Merchant, error) {

	decoder := xml.NewDecoder(reader)

	for {
		t, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				return nil, MerchantNotFund
			}
			return nil, err
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local != "merchant" {
				continue
			}

			for _, attr := range se.Attr {
				if attr.Name.Local == "id" && attr.Value == merchantId {

					merchant := &Merchant{}

					err = decoder.DecodeElement(merchant, &se)
					return merchant, err
				}
			}
		}
	}

	return nil, MerchantNotFund
}
