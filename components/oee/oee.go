package oee

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
)

type OEEComponent interface {
	CalculateOEE(ctx context.Context, data ProductionData) (float64, error)
}

type oeeComponent struct {
	weaver.Implements[OEEComponent]
}

type ProductionData struct {
	ProductionTime float64 `json:"productionTime"`
	AvailableTime  float64 `json:"availableTime"`
	IdealProdCount int     `json:"idealProdCount"`
	RejCount       int     `json:"rejCount"`
	GoodCount      int     `json:"goodCount"`
	weaver.AutoMarshal
}

func (o *oeeComponent) CalculateOEE(ctx context.Context, data ProductionData) (float64, error) {
	if data.AvailableTime == 0 || data.IdealProdCount == 0 || data.GoodCount+data.RejCount == 0 {
		return 0, fmt.Errorf("Dados insuficientes para calcular o OEE")
	}

	availability := data.ProductionTime / data.AvailableTime

	totalProdCount := data.GoodCount + data.RejCount
	performance := float64(totalProdCount) / float64(data.IdealProdCount)

	quality := float64(data.GoodCount) / float64(totalProdCount)

	oee := availability * performance * quality

	return oee, nil
}