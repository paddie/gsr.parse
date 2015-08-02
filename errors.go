package main

import (
	"errors"
)

var (
	MerchantNotFund = errors.New("MerchantNotFound")
	ReviewNotFound  = errors.New("ReviewNotFound")
)
