package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup
var requests = 50

var opensearchUrl = "Insert Opensearch endpoint"
var opensearchBasicTolen = "Insert Basic Token"

func main() {
	jsonFile, err := os.Open("../resources/1mb.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened json file")
	defer jsonFile.Close()
	wg.Add(requests)
	for i := 0; i < requests; i++ {
		go sendRequest(jsonFile)
	}
	wg.Wait()
}

func sendRequest(jsonFile *os.File) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", opensearchUrl, jsonFile)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+opensearchBasicTolen)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
	wg.Done()
}
