package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime"
	"time"
)

type Newsource struct {
	Status  string    `json:"status"`
	Sources []Sources `json:"sources"`
}

type Sources struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Language    string   `json:"Language"`
	Country     string   `json:"country"`
	Sortby      []Sortby `json:"sortsbyavailable"`
}

type Sortby struct {
}

type Numverify struct {
	Status      string        `json:"status"`
	Source      string        `json:"source"`
	SortBy      string        `json:"sortBy"`
	Newscontent []Newscontent `json:"articles"`
}

type Newscontent struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	UrlToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
}

var tpl string

func index(w http.ResponseWriter, r *http.Request) {

	/////////////////////////////////////////// news sources/////////////////////////////////////////////////////////////

	url := "https://newsapi.org/v1/sources?language=en"
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
	fmt.Println("News Count  = ", len(result.Sources))

	tpl = `<!doctype html>
<html lang="en" class="no-js">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link href='http://fonts.googleapis.com/css?family=Open+Sans:400,300,700' rel='stylesheet' type='text/css'>
<link rel="stylesheet" href="static/css/reset.css"> <!-- CSS reset -->
<link rel="stylesheet" href="static/css/style.css"> <!-- Resource style -->
<script src="static/js/modernizr.js"></script> <!-- Modernizr -->
<title>All News | Top News</title>
</head>
<body>
<header class="cd-header">
<h1>ALL IN ONE NEWS CHANNEL</h1>
</header>
<div  class="cd-pricing-container cd-has-margins">
<ul class="cd-pricing-list">
<li>
<header class="cd-pricing-header">
<h2>Choose News Source</h2>
</header> <!-- .cd-pricing-header -->
<div class="cd-pricing-body">
<ul class="cd-pricing-features">`

	fmt.Fprintf(w, tpl)

	for i := 0; i < len(result.Sources); i++ {
		fmt.Println(result.Sources[i].Id)
		fmt.Println(result.Sources[i].Name)
		fmt.Println(result.Sources[i].Category)
		fmt.Println(result.Sources[i].Description)
		fmt.Println(result.Sources[i].Country)
		fmt.Println(result.Sources[i].Language)

		tpl = `<li title="Description: ` + result.Sources[i].Description + `"> <a href="?source=` + result.Sources[i].Id + `"><em>Category:</em> ` + result.Sources[i].Category + ` || <em>Name:</em> ` + result.Sources[i].Name + ` </a> </li>`

		fmt.Fprintf(w, tpl)
	}

	//////////////////////////////////////////end news source///////////////////////////////////////////////////////////

	tpl = `</ul>
</div> <!-- .cd-pricing-body -->
<footer class="cd-pricing-footer">
<a class="cd-select" href="http://codyhouse.co/?p=429">Select</a>
</footer> <!-- .cd-pricing-footer -->
</li>
<li class="cd-popular">
<header class="cd-pricing-header">
<h2>Choose News</h2>
</header> <!-- .cd-pricing-header -->
<div class="cd-pricing-body">
<ul class="cd-pricing-features">
`

	fmt.Fprintf(w, tpl)
	//return

	//////////////////////////////////////// News Articles ///////////////////////////////////////////////////////////

	source := r.FormValue("source")

	if source != "" {

		url = "https://newsapi.org/v1/articles?source=" + source + "&sortBy=top&apiKey=a0602ff6939d455c8460e8b19fb2c4fa"
	} else {
		url = "https://newsapi.org/v1/articles?source=cnn&sortBy=top&apiKey=a0602ff6939d455c8460e8b19fb2c4fa"
	}
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
	fmt.Println("News Count  = ", len(record.Newscontent))

	for i := 0; i < len(record.Newscontent); i++ {
		fmt.Println(record.Newscontent[i].Author)
		fmt.Println(record.Newscontent[i].Description)
		fmt.Println(record.Newscontent[i].Title)
		fmt.Println(record.Newscontent[i].Url)
		fmt.Println(record.Newscontent[i].UrlToImage)
		fmt.Println(record.Newscontent[i].PublishedAt)

		tpl = `<li> <input name="news_image" type="image" value="img" src="` + record.Newscontent[i].UrlToImage + `" alt="` + record.Newscontent[i].Title + `" align="left" width="60" height="50" border="1"> <em>Author:</em> ` + record.Newscontent[i].Author + ` <br> <a href="` + record.Newscontent[i].Url + `" title="` + record.Newscontent[i].Description + `" target="_blank" > ` + record.Newscontent[i].Title + ` </a> <br> <em>PublishedAt:</em> ` + record.Newscontent[i].PublishedAt + ` </li>`

		//tpl = `<li> <input name="news_image" type="image" value="img" src=" ` + record.Newscontent[i].UrlToImage + ` " alt="`+ record.Newscontent[i].Title +`" align="left" width="60" height="50" border="1"> <em> Author:</em> `+ record.Newscontent[i].Author +` <br> <a href="`+ record.Newscontent[i].Url +`" title="`+ record.Newscontent[i].Description +`" target="_blank" > ` + record.Newscontent[i].Title + `</a> <br> <em>PublishedAt:</em> `+  record.Newscontent[i].PublishedAt + ` </li>`
		fmt.Fprintf(w, tpl)
	}

	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	ua := r.Header.Get("User-Agent")
	// get current timestamp
	currenttime := time.Now().Local().String()
	os := runtime.GOOS

	tpl = `
	</ul>
	</div> <!-- .cd-pricing-body -->
	<footer class="cd-pricing-footer">
	<a class="cd-select" href="http://codyhouse.co/?p=429">Select</a>
	</footer> <!-- .cd-pricing-footer -->
	</li>
	<li>
	<header class="cd-pricing-header">
	<h2>User Information</h2>
	</header> <!-- .cd-pricing-header -->

	<div class="cd-pricing-body">
	<ul class="cd-pricing-features">
	<li><em>User IP Address</em> ` + ip + `</li>
	<li><em>User Browser</em> ` + ua + `</li>
	<li><em>Current Time</em> ` + currenttime + `</li>
	<li><em>Operating System</em> ` + os + `</li>
	<li><em>Unlimited</em> Bandwidth</li>
	</ul>
	</div> <!-- .cd-pricing-body -->
	<footer class="cd-pricing-footer">
	<a class="cd-select" href="http://codyhouse.co/?p=429">Select</a>
	</footer> <!-- .cd-pricing-footer -->
	</li>
	</ul> <!-- .cd-pricing-list -->
	</div> <!-- .cd-pricing-container -->
	<script src="static/js/jquery-2.1.1.js"></script>
	<script src="static/js/main.js"></script> <!-- Resource jQuery -->
	</body>
	</html>`

	fmt.Fprintf(w, tpl)
	//////////////////////////////////////// News Articles ///////////////////////////////////////////////////////////

}

func main() {

	//timeout("nduson2k@gmail.com")

	//return

	http.HandleFunc("/", index)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//http.ListenAndServe(":6060", nil)
	http.ListenAndServe(":8080", nil)

}
