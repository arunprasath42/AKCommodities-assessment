// service/exchange_rates.go
package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"web-api/src/repository"
)

type ExchangeRatesService struct{}
type ExchangeRates struct {
	ID       int       `json:"id" gorm:"primary_key"`
	Date     time.Time `json:"date"`
	Base     string    `json:"base"`
	Currency string    `json:"currency"`
	Rate     float64   `json:"rate"`
}

type ExchangeRateResponse struct {
	Disclaimer string             `json:"disclaimer"`
	License    string             `json:"license"`
	Timestamp  int64              `json:"timestamp"`
	Base       string             `json:"base"`
	Rates      map[string]float64 `json:"rates"`
}

func (s *ExchangeRatesService) FetchHistoricalExchangeRates() (string, error) {
	appID := "8c97ab25682a4b22873780655a250683"
	baseURL := "https://openexchangerates.org/api/"

	// Calculate the start date as 10 days ago
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -10)

	for date := startDate; date.Before(endDate); date = date.AddDate(0, 0, 1) {
		formattedDate := date.Format("2006-01-02")
		url := baseURL + "historical/" + formattedDate + ".json?app_id=" + appID

		response, err := http.Get(url)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		defer response.Body.Close()

		var exchangeRateResponse ExchangeRateResponse
		err = json.NewDecoder(response.Body).Decode(&exchangeRateResponse)
		if err != nil {
			fmt.Println("Error decoding response:", err)
			continue
		}

		// Insert the extracted exchange rates into the database
		for currency, rate := range exchangeRateResponse.Rates {
			exchangeRate := ExchangeRates{
				Date:     date,
				Base:     exchangeRateResponse.Base,
				Currency: currency,
				Rate:     rate,
			}

			fmt.Println("Inserting exchange rate into DB:", exchangeRate)
			err = repository.Repo.Insert(&exchangeRate)
			if err != nil {
				fmt.Printf("Error inserting exchange rate into DB: %v\n", err)
			} else {
				fmt.Printf("Inserted exchange rate into DB: %+v\n", exchangeRate)
			}
		}
	}

	fmt.Println("Done!")
	return "Done!", nil
}

// GetExchangeRates retrieves exchange rates from the database with optional filtering by base currency and specific currency
func (s *ExchangeRatesService) GetExchangeRates(baseCurrency, currency string) ([]ExchangeRates, error) {
	var exchangeRates []ExchangeRates
	err := repository.Repo.FindAll(&exchangeRates, map[string]interface{}{"base": baseCurrency, "currency": currency})
	if err != nil {
		return nil, err
	}
	return exchangeRates, nil
}
