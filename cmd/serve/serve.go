package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Stats struct {
	DomainsBeingBlocked int `json:"domains_being_blocked"`
}

func main() {
	// Can we hit the API?
	remote := "http://howard-dns/admin/api.php"
	resp, err := http.Get(remote)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	stats := Stats{}
	err = json.Unmarshal(body, &stats)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(stats.DomainsBeingBlocked)
}
