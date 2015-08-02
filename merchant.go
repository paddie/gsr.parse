package main

import (
	"bytes"
	"fmt"
	"time"
)

type Merchant struct {
	BusinessUnitId string   `xml:"id,attr"`
	Url            string   `xml:"merchant_info>merchant_url"`
	Name           string   `xml:"merchant_info>name"`
	Reviews        []Review `xml:"review"`
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
