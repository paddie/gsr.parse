package main

import (
	"bufio"
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type XmlSchema int

const (
	v3 XmlSchema = iota
	v4
)

type Feeder interface {
	Map() Feed
}

func ParseFeed(path string) (Feed, error) {

	xml, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	if strings.HasSuffix(path, "xml.gz") {
		gzxml, err := gzip.NewReader(xml)
		if err != nil {
			return nil, err
		}

		return parse(gzxml)
	}

	return parse(xml)
}

func parse(r io.Reader) (Feed, error) {

	rd := bufio.NewReader(r)

	scheme, err := determineXmlSchema(rd)
	if err != nil {
		return nil, err
	}

	return parseSchema(rd, scheme)
}

func parseSchema(r io.Reader, scheme XmlSchema) (Feed, error) {

	var feed Feeder
	switch scheme {
	case v4:
		feed = &feed_V4{}
	case v3:
		feed = &feed_V3{}
	}

	dec := xml.NewDecoder(r)
	err := dec.Decode(feed)
	if err != nil {
		return nil, err
	}

	return feed.Map(), nil
}

func determineXmlSchema(rd *bufio.Reader) (XmlSchema, error) {

	b, err := rd.Peek(600)
	if err != nil {
		return 0, fmt.Errorf("Unable to determine 'merchant_reviews.xsd' from header '%s'", err.Error())
	}

	header := string(b)
	if strings.Contains(header, "3.0/merchant_reviews.xsd") {
		// log.Println("merchant_reviews.xsd v3.0 detected")
		return v3, nil
	} else if strings.Contains(header, "4.0/merchant_reviews.xsd") {
		// log.Println("merchant_reviews.xsd v4.0 detected")
		return v4, nil
	}

	return 0, fmt.Errorf("Unable to determine 'merchant_reviews.xsd' version from: '%s'", header)
}
