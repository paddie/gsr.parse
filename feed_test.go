package main

import (
	"strings"
	"testing"
	"time"
)

const (
	MERCHANT_ID = "46d8a49d000064000500fa08"
	REVIEW_ID   = "546add240000640002b4e74a"
	REVIEWER_ID = "546adcf80000640001974608"
	FEED        = `
<?xml version="1.0" encoding="utf-8"?>
<feed xmlns="http://schemas.google.com/merchant_reviews/4.0"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:schemaLocation="http://schemas.google.com/merchant_reviews/4.0 http://www.google.com/shopping/reviews/schema/merchant/4.0/merchant_reviews.xsd">
<author>
<name>Trustpilot</name>
<email>support@trustpilot.com</email>
</author>
<merchant id="46d8a49d000064000500fa08">
	<merchant_info>
		<name>Westpac</name>
		<country>AU</country>
		<merchant_url>http://www.westpac.com.au/</merchant_url>
		<rating_url type="detail">https://au.trustpilot.com/review/www.westpac.com.au</rating_url>
	</merchant_info>
		<review id="546add240000640002b4e74a">
			<reviewer_type>user</reviewer_type>
			<may_show_full_content>true</may_show_full_content>
			<collection_method>after_fulfillment</collection_method>
			<review_url type="singleton">https://au.trustpilot.com/review/www.westpac.com.au/546add240000640002b4e74a</review_url>
			<reviewer_id>546adcf80000640001974608</reviewer_id>
			<reviewer>Sarah Klesser</reviewer>
			<review_date>2014-11-18T05:46:12Z</review_date>
			<language>en</language>
			<title>Convenient Personal Banking Services</title>
			<content>I’ve been a Westpac customer for almost five years and have to give them kudos for their convenient online and mobile banking services. I can now check my account balance, transfer funds, and even pay my telephone bills with just a click of my fingers…and all from the comfort of my own home. What I’m most happy with is the good customer service I receive whenever I run into any issues. To date, I've only had to contact customer support twice, and that’s because I’m not a very tech savvy person (and not because my money went missing or anything like that!) I highly recommend this bank.</content>
			<ratings>
				<overall min="1" max="5">5</overall>
			</ratings>
		</review>
	</merchant>
</feed>`
)

func TestFeed(t *testing.T) {

	c := Context{t}

	r := strings.NewReader(FEED)

	merchant, err := findMerchantId(r, MERCHANT_ID)
	c.Fatal(err)

	c.StringMatch(merchant.BusinessUnitId, MERCHANT_ID, "BusinessUnitId")

	c.IntMatch(len(merchant.Reviews), 1, "Review Count")

	review := merchant.Reviews[0]
	c.StringMatch(review.Id, REVIEW_ID, "Review.Id")

	date, err := time.Parse(time.RFC3339, "2014-11-18T05:46:12Z")
	c.Fatal(err)
	c.DateMatch(review.Date, date, "Review.Date")
}

type Context struct {
	t *testing.T
}

func (c Context) StringMatch(actual, expected, key string) {
	if actual != expected {
		c.t.Fatalf("Unexpected %s: Expected %s != %s Actual", key, expected, actual)
	}
}

func (c Context) IntMatch(actual, expected int, key string) {
	if actual != expected {
		c.t.Fatalf("Unexpected %s: Expected %s != %s Actual", key, expected, actual)
	}
}

func (c Context) DateMatch(actual, expected time.Time, key string) {
	if !actual.Equal(expected) {
		c.t.Fatalf("Unexpected %s: Expected %s != %s Actual", key, expected, actual)
	}
}

func (c Context) Fatal(err error) {
	if err != nil {
		c.t.Fatal(err)
	}
}
