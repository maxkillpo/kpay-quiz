package service

import (
	"kpay-quiz/store"
	"time"
)

type SellReportsService interface {
	All(id string) (*ReportResponse, error)
}

type SellReportsServiceImprement struct {
	DB *store.DAO
}

type ReportResponse struct {
	Date     string `json:"date"`
	Products []struct {
		Name          string `json:"name"`
		SellingVolume int    `json:"selling_volume"`
	} `json:"products"`
	Accumulate float64 `json:"accumulate"`
}

func (s *SellReportsServiceImprement) All(id string) (*ReportResponse, error) {

	merchant, err := s.DB.FindMerchantById(id)

	var report ReportResponse

	if err != nil {
		return &report, err
	}

	historys := merchant.History

	for _, history := range historys {

		var product struct {
			Name          string `json:"name"`
			SellingVolume int    `json:"selling_volume"`
		}

		product.Name = history.ProductName
		product.SellingVolume = history.Amount
		report.Products = append(report.Products, product)
		report.Accumulate = report.Accumulate + history.TotalPrice
	}

	report.Date = time.Now().String()

	return &report, err
}
