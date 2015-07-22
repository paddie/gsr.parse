package main

import (
	"bufio"
	"compress/gzip"
	"encoding/xml"
	"io"
	"log"
	"os"
	"strings"
)

func ParseFeed(path string) (*Feed, error) {

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

func parse(r io.Reader) (*Feed, error) {

	rd := bufio.NewReader(r)
	dec := xml.NewDecoder(rd)

	feed := &Feed{}

	err := dec.Decode(feed)
	if err != nil {
		return nil, err
	}

	return feed, nil
}
