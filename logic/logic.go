package logic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/matisidler/smiles-bot/models"
)

func GetRequest(year, month, day, originAirport, destinationAirport string, onlyDayBefore bool) ([]models.Day, error) {
	requestURL := fmt.Sprintf("https://api-prd-airlines-carousel.smiles.com.br/v1/airlines/carousel/pricing?adults=1&cabinType=all&children=0&currencyCode=ARS&departureDate=%s-%s-%s&destinationAirportCode=%s&infants=0&isFlexibleDateChecked=false&originAirportCode=%s&tripType=2&forceCongener=true&r=ar", year, month, day, destinationAirport, originAirport)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return nil, err
	}

	var bp models.BPModel
	err = json.NewDecoder(res.Body).Decode(&bp)
	if err != nil {
		fmt.Printf("client: error decoding json: %s\n", err)
		return nil, err
	}
	if day == "28" {
		if bp.BestPricingSegmentList[0].ListOfDatesAndMiles[len(bp.BestPricingSegmentList[0].ListOfDatesAndMiles)-1].Date[5:7] == month {
			return bp.BestPricingSegmentList[0].ListOfDatesAndMiles[3:], nil
		}
		return bp.BestPricingSegmentList[0].ListOfDatesAndMiles[3 : len(bp.BestPricingSegmentList[0].ListOfDatesAndMiles)-1], nil

	}

	if onlyDayBefore {
		newArray := []models.Day{bp.BestPricingSegmentList[0].ListOfDatesAndMiles[2]}
		return newArray, nil
	}

	return bp.BestPricingSegmentList[0].ListOfDatesAndMiles, nil
}

func LookForBestPrices(bp []models.Day) []models.Day {
	sort.SliceStable(bp, func(i, j int) bool {
		return bp[i].Miles < bp[j].Miles
	})
	return bp
}

func Remove0Values(bp []models.Day) []models.Day {
	var newArray []models.Day
	for _, day := range bp {
		if day.Miles != 0 {
			newArray = append(newArray, day)
		}
	}
	return newArray
}
