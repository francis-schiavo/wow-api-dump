package main

type ConnectedRealms struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ConnectedRealms []struct {
		Href string `json:"href"`
	} `json:"connected_realms"`
}