package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Newsource struct {
	Status      string `json:"status"`
Sources []Sources `json:"sources"`
}

type Sources struct {
	Id string  `json:"id"`
	Name string  `json:"name"`
	Description string  `json:"description"`
	Category string  `json:"category"`
	Language string  `json:"Language"`
	Country string  `json:"country"`
	Sortby []Sortby `json:"sortsbyavailable"`
}

type Sortby struct {

}

type Numverify struct {
	Status      string `json:"status"`
	Source       string `json:"source"`
	SortBy       string `json:"sortBy"`
Newscontent []Newscontent `json:"articles"`
}

type Newscontent struct {
	Author string  `json:"author"`
	Title string  `json:"title"`
	Description string  `json:"description"`
	Url string  `json:"url"`
	UrlToImage string  `json:"urlToImage"`
	PublishedAt string  `json:"publishedAt"`
}


func main() {



	url :="https://newsapi.org/v1/sources?language=en"
		// Build the request
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal("NewRequest: ", err)
			return
		}

		// For control over HTTP client headers,
		// redirect policy, and other settings,
		// create a Client
		// A Client is an HTTP client
		client := &http.Client{}

		// Send the request via a client
		// Do sends an HTTP request and
		// returns an HTTP response
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal("Do: ", err)
			return
		}

		// Callers should close resp.Body
		// when done reading from it
		// Defer the closing of the body
		defer resp.Body.Close()

		// Fill the record with the data from the JSON
		var result Newsource

		// Use json.Decode for reading streams of JSON data
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			log.Println(err)
		}


		fmt.Println("Status   = ", result.Status)
		fmt.Println("News Count  = ",len(result.Sources))

		for i := 0; i < len(result.Sources); i++ {
			fmt.Println(result.Sources[i].Id)
			fmt.Println(result.Sources[i].Name)
			fmt.Println(result.Sources[i].Category)
			fmt.Println(result.Sources[i].Description)
			fmt.Println(result.Sources[i].Country)
			fmt.Println(result.Sources[i].Language)
		}







	return
	//phone := "14158586273"
	// QueryEscape escapes the phone string so
	// it can be safely placed inside a URL query
	//safePhone := url.QueryEscape(phone)

	///url := fmt.Sprintf("http://apilayer.net/api/validate?access_key=YOUR_ACCESS_KEY&number=%s", safePhone)





////////// source ////////////////////////////////////////
url ="https://newsapi.org/v1/articles?source=cnn&sortBy=top&apiKey=a0602ff6939d455c8460e8b19fb2c4fa"
	// Build the request
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client = &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record Numverify

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}


	fmt.Println("Status   = ", record.Status)
	fmt.Println("Source  = ", record.Source)
	fmt.Println("SortBy  = ", record.SortBy)
	fmt.Println("News Count  = ",len(record.Newscontent))

	for i := 0; i < len(record.Newscontent); i++ {
		fmt.Println(record.Newscontent[i].Author)
		fmt.Println(record.Newscontent[i].Description)
		fmt.Println(record.Newscontent[i].Title)
		fmt.Println(record.Newscontent[i].Url)
		fmt.Println(record.Newscontent[i].UrlToImage)
		fmt.Println(record.Newscontent[i].PublishedAt)
	}


}
