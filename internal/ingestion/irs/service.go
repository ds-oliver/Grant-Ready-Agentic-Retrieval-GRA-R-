package irs

import (
	"encoding/xml"
	"fmt"
	"io"
)

// FinancialMetrics captures normalized financial values derived from IRS Form 990 data.
type FinancialMetrics struct {
	TotalRevenue           float64
	TotalExpenses          float64
	ProgramExpenses        float64
	AdministrativeExpenses float64
	FundraisingExpenses    float64
	ProgramExpenseRatio    float64
}

// Service parses IRS Form 990 XML data and produces financial metrics.
type Service struct{}

// ParseForm990 parses the provided XML stream and maps values into FinancialMetrics.
func (s Service) ParseForm990(r io.Reader) (FinancialMetrics, error) {
	var form form990XML
	decoder := xml.NewDecoder(r)
	if err := decoder.Decode(&form); err != nil {
		return FinancialMetrics{}, fmt.Errorf("decode form 990 xml: %w", err)
	}

	metrics := FinancialMetrics{
		TotalRevenue:           form.TotRevenue,
		TotalExpenses:          form.TotExpenses,
		ProgramExpenses:        form.ProgramServiceExpenses,
		AdministrativeExpenses: form.ManagementGeneralExpenses,
		FundraisingExpenses:    form.FundraisingExpenses,
	}

	metrics.ProgramExpenseRatio = calculateProgramExpenseRatio(metrics.ProgramExpenses, metrics.TotalExpenses)

	return metrics, nil
}

func calculateProgramExpenseRatio(programExpenses, totalExpenses float64) float64 {
	if totalExpenses <= 0 {
		return 0
	}
	return programExpenses / totalExpenses
}

type form990XML struct {
	XMLName                   xml.Name `xml:"Return"`
	TotRevenue                float64  `xml:"ReturnData>IRS990>TotalRevenueCurrentYear"`
	TotExpenses               float64  `xml:"ReturnData>IRS990>TotalExpensesCurrentYear"`
	ProgramServiceExpenses    float64  `xml:"ReturnData>IRS990>ProgramServiceRevenue"`
	ManagementGeneralExpenses float64  `xml:"ReturnData>IRS990>ManagementAndGeneral"`
	FundraisingExpenses       float64  `xml:"ReturnData>IRS990>Fundraising"`
}
