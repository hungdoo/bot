package debank

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/shopspring/decimal"
)

type Portfolio struct {
	Name   string `json:"name"`
	Detail struct {
		HealthRate decimal.Decimal `json:"health_rate"`
	} `json:"detail"`
	Stats struct {
		DebtUsdValue decimal.Decimal `json:"debt_usd_value"`
	} `json:"stats"`
}
type Project struct {
	Id            string      `json:"id"`
	PortfolioList []Portfolio `json:"portfolio_list"`
}
type DebankProject struct {
	Data []Project `json:"data"`
}

const baseUrl = "https://api.debank.com/portfolio/project_list?user_addr="

func GetDebt(userAddr string) (decimal.Decimal, error) {
	resp, err := http.Get(baseUrl + userAddr)
	if err != nil {
		return decimal.Decimal{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return decimal.Decimal{}, err
	}

	var cResp DebankProject
	if err := json.Unmarshal(body, &cResp); err != nil {
		return decimal.Decimal{}, err
	}

	for _, p := range cResp.Data {
		if p.Id == "ftm_scream" {
			for _, portfo := range p.PortfolioList {
				if portfo.Name == "Lending" {
					// fmt.Println(portfo.Detail.HealthRate)
					// fmt.Println(portfo.Stats.DebtUsdValue)
					return portfo.Stats.DebtUsdValue, nil
				}
			}
		}
	}

	return decimal.Decimal{}, fmt.Errorf("DebtUsdValue not found")
}
