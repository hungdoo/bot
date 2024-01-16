package debank

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestDebank(t *testing.T) {
	URL := "https://api.debank.com/portfolio/project_list?user_addr=0xddf169bf228e6d6e701180e2e6f290739663a784"
	resp, err := http.Get(URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Println(string(body))

	var cResp DebankProject
	if err := json.Unmarshal(body, &cResp); err != nil {
		t.Fatal(err)
	}

	for _, p := range cResp.Data {
		if p.Id == "ftm_scream" {
			fmt.Printf("%+v\n", p)
			for _, portfo := range p.PortfolioList {
				if portfo.Name == "Lending" {
					fmt.Println(portfo.Detail.HealthRate)
					fmt.Println(portfo.Stats.DebtUsdValue)
				}
			}
		}
	}
}
