package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Match struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func main() {
	resp, err := http.Get("https://api.draftkings.com/contests/v1/featured?format=json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var matches []Match
	if err := json.Unmarshal(body, &matches); err != nil {
		panic(err)
	}

	for _, match := range matches {
		fmt.Printf("%s: %s\n", match.ID, match.Name)
		fmt.Println(match.Description)
		fmt.Println()
	}
}
