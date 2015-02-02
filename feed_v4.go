package main

import "log"

type feed_V4 struct {
	Merchants []merchantNode_v4 `xml:"merchant"`
}

type merchantNode_v4 struct {
	Id          string   `xml:"id,attr"`
	Name        string   `xml:"merchant_info>name"`
	MerchantUrl string   `xml:"merchant_info>merchant_url"`
	Reviews     []Review `xml:"review"`
}

func (f *feed_V4) Map() Feed {
	merchants := make(Feed, len(f.Merchants))

	for _, m := range f.Merchants {
		if _, ok := merchants[m.Id]; ok {
			log.Println("duplicate business unit id", m.Id)
			continue
		}

		merchants[m.Id] = Merchant{
			BusinessUnitId: m.Id,
			Url:            m.MerchantUrl,
			Name:           m.Name,
			Reviews:        m.Reviews,
		}
	}

	return merchants
}
