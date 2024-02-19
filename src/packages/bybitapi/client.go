package bybitapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type IdoResponse struct {
	Success bool   `json:"success"`
	Mesage  string `json:"ret_msg"`
	Result  struct {
		ProjectList []Project `json:"projectList"`
	} `json:"result"`
}

type Project struct {
	Name                  string   `json:"name"`
	RegistrationEndTime   UnixTime `json:"registrationEndTime"`
	RegistrationStartTime UnixTime `json:"registrationStartTime"`
	StartTime             UnixTime `json:"startTime"`
	EndTime               UnixTime `json:"endTime"`
	SnapshotStartTime     UnixTime `json:"snapshotStartTime"`
	SnapshotEndTime       UnixTime `json:"snapshotEndTime"`
}

var IDO_URL = "https://api2.bybit.com/spot/api/web3/ido/v1/project/list"

func FetchProjects(url string) ([]Project, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var res IdoResponse
	err = json.Unmarshal(bodyBytes, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	if !res.Success {
		return nil, fmt.Errorf("fetch failed: %s", res.Mesage)
	}

	return res.Result.ProjectList, nil
}

func GetUpcomingProjects(url string, since time.Time) ([]Project, error) {
	projects, err := FetchProjects(url)
	if err != nil {
		return nil, fmt.Errorf("failed to FetchProjects: %w", err)
	}

	var upcomings []Project
	for _, j := range projects {
		if since.Before(j.RegistrationEndTime.Time) {
			upcomings = append(upcomings, j)
		}
	}
	return upcomings, nil
}

func GetLatestProject(url string) (*Project, error) {
	projects, err := FetchProjects(url)
	if err != nil {
		return nil, fmt.Errorf("failed to FetchProjects: %w", err)
	}

	var latestIdx = int(0)
	for i := 1; i < len(projects); i++ {
		if projects[i].RegistrationEndTime.Time.After(projects[latestIdx].RegistrationEndTime.Time) {
			latestIdx = i
		}
	}
	return &projects[latestIdx], nil
}
