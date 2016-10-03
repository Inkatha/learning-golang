package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type CarDetails []struct {
    Color          string `json:"color"`
    Make           string `json:"make"`
    State          string `json:"state"`
    TowDate        string `json:"tow_date"`
    TowPhoneNumber string `json:"tow_facility_phone"`
    TowedToAddress string `json:"towed_to_address"`
}

func main() {
    url := fmt.Sprintf("https://data.cityofchicago.org/resource/rp42-fxjt.json")

    // Build the request
    req, err := http.NewRequest("GET", url, nil)
    if (err != nil) {
        log.Fatal("NewRequest: ", err)
        return
    }

    /// For control over HTTP client headers,
    // redirect policy, and other settings,
    // create a Client
    // A client is an HTTP client
    client := &http.Client{}

    // For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
    resp, err := client.Do(req)
    if (err != nil) {
        log.Fatal("Do: ", err)
        return
    }

    // Callers should close resp.Body
    // When done reading from it
    // Defer the closing of the Body
    defer resp.Body.Close()

    // Fill the record with the data from the json
    var record CarDetails

    if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
        log.Println(err)
    }

    for _, element := range record {
        fmt.Println("Color = ", element.Color)        
        fmt.Println("Make   = ", element.Make)
        fmt.Println("State  = ", element.State)
        fmt.Println("TowDate   = ", element.TowDate)
        fmt.Println("TowPhoneNumber  = ", element.TowPhoneNumber)
        fmt.Println("TowedToAddress  = ", element.TowedToAddress)
    }
}