package api_v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

type Client struct {
	SubDomain string
	ApiKey    string
}

func (c *Client) doRequest(action string, input interface{}, outputRef interface{}) {
	url := fmt.Sprintf("https://%s.kanbanize.com/index.php/api/kanbanize/%s/format/json", c.SubDomain, action)

	jsonPayload, err := json.Marshal(input)
	if err != nil {
		log.Fatalf("unable to marshal payload, %v", err)
	}

	client := &http.Client {
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalf("unable to create request, %v", err)
	}
	req.Header.Add("apikey", c.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("request error, %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("unable to read response body, %v", err)
	}

	// some responses contain dynamically created structure
	// you can pass a *string to get the raw response data
	t := reflect.TypeOf(outputRef).String()
	if t == "*string" {
		s := outputRef.(*string)
		*s = string(body)
		return
	}

	err = json.Unmarshal(body, outputRef)
	if err != nil {
		log.Fatalf("unable to parse response json, %v, %v", err, string(body))
	}
}
