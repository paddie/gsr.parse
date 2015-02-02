package main

// func validateFeed(merchs MerchantReviewMap, scheme XmlSchema) {

// 	reviews := 0
// 	merchants := 0
// 	merchant_ids := make(map[string]bool)
// 	review_ids := make(map[string]bool)

// 	for merch, rs := range merchs {
// 		reviews += len(rs)
// 		merchants += 1

// 		if !merchant_ids[merch.Id] {

// 			if outputMerchantIds {
// 				if scheme == V4 {
// 					fmt.Printf("- merchant: ObjectId(\"%s\")\n", merch.Id)
// 				} else {
// 					fmt.Println("- merchant: " + merch.Id)
// 				}
// 			}

// 			if outputMerchantUrls {
// 				fmt.Printf("- url: %s\n", merch.Url)
// 			}

// 			merchant_ids[merch.Id] = true
// 		} else {
// 			log.Println("ISSUE: Duplicate merchant_id: " + merch.Id)
// 		}

// 		var indent = ""
// 		if outputMerchantIds || outputMerchantUrls {
// 			indent = "\t"
// 		}

// 		for _, r := range rs {
// 			if !review_ids[r.Id] {
// 				review_ids[r.Id] = true
// 				if outputReviewIds {
// 					if scheme == V4 {
// 						fmt.Printf("%sObjectId(\"%s\")\n", indent, r.Id)
// 					} else {
// 						fmt.Printf(indent + r.Id)
// 					}
// 				}
// 			} else {
// 				log.Println("ISSUE: Duplicate review_id: " + r.Id)
// 			}
// 		}
// 	}

// 	if stats {
// 		log.Printf("Merchants: %d\n", merchants)
// 		log.Printf("Reviews: %d\n", reviews)
// 	}
// }
