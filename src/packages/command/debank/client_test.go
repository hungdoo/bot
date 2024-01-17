package debank

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestDebank(t *testing.T) {
	// URL := "https://api.debank.com/portfolio/project_list?user_addr=0x1c3925201a001281001749Bb0655d82e8934e9C8"
	URL := "https://api.debank.com/user/used_chains?id=0x1c3925201a001281001749Bb0655d82e8934e9C8"
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		t.Fatal("Error creating request:", err)
	}

	// Set the User-Agent and X-App-Id headers
	// 	x-api-nonce: n_dnNAOSWvzwmQiOQu52o7i0ohyaT7vnL89JvKD1tI
	// x-api-sign: f0f31f0078adbbe23a0784bf63f0c07dce637d977b51a9d4c6218e29cbc7afd5
	// x-api-ts: 1664716589
	// x-api-ver: v2
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("x-api-nonce", "n_dnNAOSWvzwmQiOQu52o7i0ohyaT7vnL89JvKD1tI")
	req.Header.Set("x-api-sign", "f0f31f0078adbbe23a0784bf63f0c07dce637d977b51a9d4c6218e29cbc7afd5")
	req.Header.Set("x-api-ts", "1664716589")
	req.Header.Set("x-api-ver", "v2")

	// Send the request using http.DefaultClient
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal("Error sending request:", err)
	}

	// Don't forget to close the response body
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf(resp.Status)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
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
