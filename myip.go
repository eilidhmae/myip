package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type MyIP struct{
	Address	string `json:"ip"`
}

func (r *MyIP) Label() string {
	switch r.Address {
	default:
		return "UNKNOWN"
	}
}

func New() (*MyIP, error) {
	apiUrl := `https://api.ipify.org?format=json`
	myip := new(MyIP)
	resp, err := http.Get(apiUrl)
	if err != nil {
		return myip, err
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return myip, err
	}
	if err := json.Unmarshal(buf, &myip); err != nil {
		return myip, err
	}
	return myip, nil
}
