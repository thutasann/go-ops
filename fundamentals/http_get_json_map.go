package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Response interface {
	GetResponse() string
}

type Page struct {
	Name string `json:"page"`
}

type Words struct {
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func (w Words) GetResponse() string {
	return fmt.Sprintf("response: %s", strings.Join(w.Words, ", "))
}

type Occurrence struct {
	Words map[string]int `json:"words"`
}

func (o Occurrence) GetResponse() string {
	out := []string{}

	for k, v := range o.Words {
		out = append(out, fmt.Sprintf("%s (%d)", k, v))
	}

	return fmt.Sprintf("response: %s", strings.Join(out, ", "))
}

func Http_JSON_Get_Sample() {
	var (
		requestUrl string
		password   string
	)

	flag.StringVar(&requestUrl, "url", "", "url to access")
	flag.StringVar(&password, "password", "", "password to access our API")

	flag.Parse()

	res, err := do_request(requestUrl)
	if err != nil {
		if requestErr, ok := err.(RequestError); ok {
			fmt.Printf("Request Err: %s (HTTP code: %d, Body: %s\n)", requestErr.Err, requestErr.HTTPCode, requestErr.Body)
			os.Exit(1)
		}
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	if res == nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Response: %s", res.GetResponse())
}

func do_request(requestUrl string) (Response, error) {
	if _, err := url.ParseRequestURI(requestUrl); err != nil {
		return nil, fmt.Errorf("validation error: %s", err)
	}

	response, err := http.Get(requestUrl)

	if err != nil {
		return nil, fmt.Errorf("http get error: %s", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("ReadAll error: %s", err)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("invalid output (http code %d) %s", response.StatusCode, body)
	}

	var page Page

	err = json.Unmarshal(body, &page)
	if err != nil {
		return nil, RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(body),
			Err:      fmt.Sprintf("Page Unmarshal Error: %s", err),
		}
	}

	switch page.Name {
	case "words":
		var words Words
		err = json.Unmarshal(body, &words)
		if err != nil {
			return nil, RequestError{
				HTTPCode: response.StatusCode,
				Body:     string(body),
				Err:      fmt.Sprintf("Words Unmarshal Error: %s", err),
			}
		}

		return words, nil
	case "occurrence":
		var occurrence Occurrence
		err = json.Unmarshal(body, &occurrence)
		if err != nil {
			return nil, RequestError{
				HTTPCode: response.StatusCode,
				Body:     string(body),
				Err:      fmt.Sprintf("Occurrence Unmarshal Error: %s", err),
			}
		}

		return occurrence, nil
	}

	return nil, nil
}
