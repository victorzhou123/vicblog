package dto

type DashboardDataDto struct {
	ArticleCounts       int `json:"articleCounts"`
	CategoryCounts      int `json:"categoryCounts"`
	TagCounts           int `json:"tagCounts"`
	ArticleVisitsCounts int `json:"articleVisitsCounts"`
}

type HeatMapDto struct {
	HeatMap map[string]int `json:"heatMap"`
}
