package ClientRest

import (
	"log"

	"github.com/go-resty/resty/v2"
)

func RestGet(url string) (*resty.Response, error) {
	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get(url)

	if err != nil {
		log.Fatal(err)
	}
	return resp, err

}

func RestPost(url string, payload string) (*resty.Response, error) {
	// Create a Resty Client
	client := resty.New()

	// POST JSON string
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		//  SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
		Post(url)

	return resp, err

}
