package keb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetRuntimes(endpoint string) (*Runtimes, error) {
	resp, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)

	var runtimes Runtimes

	err = json.NewDecoder(resp.Body).Decode(&runtimes)

	if err != nil {
		fmt.Printf("invalid JSON: %v", err)
		return nil, err
	}
	return &runtimes, err
}
