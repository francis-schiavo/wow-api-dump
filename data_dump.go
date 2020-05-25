package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"sync"
)

type AHDump struct {
	OutputDir string
	realms    chan int
	waitGroup *sync.WaitGroup
}

func (dump *AHDump) worker() {
	for realmId := range dump.realms {
		response := APIClient.Auction(realmId, nil)
		if response.Error != nil || response.Status == 404 {
			fmt.Printf("Failed to download auction data for connected realm %d\n", realmId)
		}

		filename := fmt.Sprintf("%s/ConnectedRealm%d.json", dump.OutputDir, realmId)

		err := ioutil.WriteFile(filename, response.Body, 0644)
		if err != nil {
			fmt.Printf("Failed to save dump file [%s] for connected realm %d: %s\n", filename, realmId, err.Error())
		} else {
			fmt.Printf("Auction data for connected realm %d saved in: %s.\n", realmId, filename)
		}
		dump.waitGroup.Done()
	}
}

func (dump *AHDump) Run(concurrency int) {
	connectedRealmsResponse := APIClient.ConnectedRealmIndex(nil)
	if connectedRealmsResponse.Error != nil {
		log.Fatal("Failed to fetch connected realm index")
	}

	var connectedRealms ConnectedRealms
	connectedRealmsResponse.Parse(&connectedRealms)

	dump.waitGroup = &sync.WaitGroup{}
	dump.realms = make(chan int)

	for w := 1; w <= concurrency; w++ {
		go dump.worker()
	}

	regexID, _ := regexp.Compile("/([0-9]+)\\?")
	for _, connectedRealm := range connectedRealms.ConnectedRealms {
		id, _ := strconv.Atoi(regexID.FindStringSubmatch(connectedRealm.Href)[1])
		dump.realms <- id
		dump.waitGroup.Add(1)
	}

	dump.waitGroup.Wait()
}
