package dto

import "github.com/victorzhou123/vicblog/statistics/domain/entity"

type VisitsOfAWeekDto struct {
	Counts []int    `json:"counts"`
	Dates  []string `json:"dates"`
}

// inputs must be ordered as ascend
func ToVisitsOfAWeekDto(visits []entity.ArticleDailyVisits) VisitsOfAWeekDto {

	counts, dates := make([]int, len(visits)), make([]string, len(visits))

	for i := range visits {
		counts[i] = visits[i].Total.Amount()
		dates[i] = visits[i].Date.TimeMonthDayOnly()
	}

	return VisitsOfAWeekDto{
		Counts: counts,
		Dates:  dates,
	}
}
