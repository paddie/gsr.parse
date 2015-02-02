package main

import "log"

type feed_V3 struct {
	Merchants []merchantNode_v3 `xml:"merchants>merchant"`
}

type merchantNode_v3 struct {
	Id          string   `xml:"id,attr"`
	Name        string   `xml:"name"`
	MerchantUrl string   `xml:"merchant_url"`
	Reviews     []Review `xml:"reviews>review"`
}

func (f *feed_V3) Map() Feed {
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
