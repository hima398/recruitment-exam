package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

type Result struct {
	Winner string `json: winner`
	Loser  string `json: loser`
}

var baseUrl = "https://ob6la3c120.execute-api.ap-northeast-1.amazonaws.com/Prod"

func Battle(m1, m2 string) *Result {

	c := &http.Client{}
	resp, err := c.Get(fmt.Sprintf("%s/battle/%s+%s", baseUrl, m1, m2))
	if err != nil {
		// TODO return err
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var result Result
	_ = json.Unmarshal(body, &result)

	return &result

}
