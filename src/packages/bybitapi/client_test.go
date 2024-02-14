package bybitapi

import (
	"testing"
	"time"

	"github.com/hungdoo/bot/src/packages/utils"
)

func TestFetchProjects(t *testing.T) {
	projs, err := FetchProjects(IDO_URL)
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(projs)
	utils.PrettyPrint(projs)
}

func TestFetchUpcommingProjects(t *testing.T) {
	projs, err := GetUpcomingProjects(IDO_URL, time.Unix(1706781600, 0)) // 2024-02-01 10:00:00
	if err != nil {
		t.Fatal(err)
	}
	utils.PrettyPrint(projs)
}

func TestGetLatestProject(t *testing.T) {
	projs, err := GetLatestProject(IDO_URL)
	if err != nil {
		t.Fatal(err)
	}
	utils.PrettyPrint(projs)
}
