package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func extractText() string {
	url := "https://randomincategory.toolforge.org/Featured_articles?site=en.wikipedia.org"
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	pageTitle := strings.TrimPrefix((*resp.Request.URL).String(), "https://en.wikipedia.org/wiki/")

	req, err := http.NewRequest("GET", "https://en.wikipedia.org/w/api.php", nil)

	if err != nil {
		log.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("action", "query")
	q.Add("format", "json")
	q.Add("titles", pageTitle)
	q.Add("prop", "extracts")
	q.Add("exsentences", "10")
	q.Add("explaintext", "true")

	req.URL.RawQuery = q.Encode()

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	// var result Response
	var result map[string]interface{}

	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	// fmt.Println(result["query"].(map[string]interface{})["pages"])
	pageJson := result["query"].(map[string]interface{})["pages"].(map[string]interface{})

	var text string
	for key := range pageJson {
		text = pageJson[key].(map[string]interface{})["extract"].(string)

	}
	return text
}
