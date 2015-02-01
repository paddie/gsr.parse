package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	V3 XmlScheme = iota
	V4
)

type XmlScheme int

var (
	path string
)

func init() {
	flag.StringVar(&path, "p", "", "path to a GoogleSellerRatings Feed")
}

func main() {
	flag.Parse()

	if path == "" {
		log.Fatalf("Invalid feed path '%s'", path)
	}

	xml, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	feed, err := ParseFeed(xml)
	if err != nil {
		log.Fatal(err)
	}

	reviews := 0
	merchants := 0
	merchant_ids := make(map[string]bool)
	review_ids := make(map[string]bool)
	for merch, rs := range feed.MerchantReviews() {
		reviews += len(rs)
		merchants += 1

		if !merchant_ids[merch] {
			merchant_ids[merch] = true
		} else {
			log.Println("Duplicate Merchant Ids: " + merch)
		}

		for _, r := range rs {
			if !review_ids[r.Id] {
				review_ids[r.Id] = true
			} else {
				log.Println("Duplicate Review Ids: " + r.Id)
			}
		}
	}

	fmt.Println("-------------------------------------")
	fmt.Printf("Merchants: %d\n", merchants)
	fmt.Printf("Reviews: %d\n", reviews)

	fmt.Println("\nSuccess!")
}
