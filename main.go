package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-ini/ini"
)

type btResult struct {
	ID     int `json:"id"`
	Result []struct {
		BgID              int       `json:"bg_id"`
		BgShortDesc       string    `json:"bg_short_desc"`
		BgReportedDate    string    `json:"bg_reported_date"`
		BgLastUpdatedDate string    `json:"bg_last_updated_date"`
		BgLatestDate      time.Time `json:"bg_latest_date"`
		BgBlock           string    `json:"bg_block"`
		StName            string    `json:"st_name"`
		BgAssignedToUser  string    `json:"bg_assigned_to_user"`
		UsAlias           string    `json:"us_alias"`
		BgClass           string    `json:"bg_class"`
	} `json:"result"`
}

func main() {
	url, err := readCfg("")
	if err != nil {
		log.Fatal("ErrorReadConfig:", err)
		return
	}

	// building request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest:", err)
		return
	}

	// building client
	client := &http.Client{}

	// building response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do:", err)
		return
	}

	defer resp.Body.Close()

	var record btResult

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	// fmt.Println(record.Result)

	// print Header
	// FIXME: Need to limit result to fit the screen!
	fmt.Printf("|%-8s|%-50s|%-10s|%-10s|%-6s|%-6s|%-20s|\n\n", "BT ID", "DESCRIPTION", "CREATED", "LAST MOD", "ASSIGN", "BLOCKED", "STATUS")
	for _, r := range record.Result {
		// print Body
		fmt.Printf("|%-8d|%-50s|%-10s|%-10s|%-6s|%-6s|%-20s|\n", r.BgID, r.BgShortDesc, r.BgReportedDate, r.BgLastUpdatedDate, r.UsAlias, r.BgBlock, r.StName)
	}
}

// readCfg - function that is quite complete to feed in non-default config file using flag in example
func readCfg(p string) (string, error) {
	if p == "" {
		p = "config.ini"
	}
	iniCfg, err := ini.Load(p)
	if err != nil {
		fmt.Println("CONFIG")
		return "", err
	}
	iniSection := iniCfg.Section("server")
	iniHost := iniSection.Key("host").String()
	iniAPI1 := iniSection.Key("api1").String()
	urlAPI := iniHost + iniAPI1

	return urlAPI, nil
}
