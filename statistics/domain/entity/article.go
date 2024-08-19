package entity

import cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"

type ArticleDailyVisits struct {
	Total cmprimitive.Amount
	Date  cmprimitive.Timex
}
