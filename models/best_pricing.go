package models

type BPModel struct {
	BestPricingSegmentList []struct {
		ListOfDatesAndMiles []Day `json:"calendarDayList"`
	} `json:"bestPricingSegmentList"`
}

type Day struct {
	Date  string `json:"date"`
	Miles int    `json:"miles"`
}
