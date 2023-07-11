package httpclient

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type HttpClient struct {
	client *http.Client
}

func New() HttpClient {
	client := &http.Client{}
	return HttpClient{client}
}

func (this *HttpClient) Get(url string, response any) error {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return errors.Wrap(err, "cant create a new request")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Fallstudie, Tobias Nickel")

	res, err := this.client.Do(req)
	if err != nil {
		return errors.Wrap(err, "cant do the request")
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&response)
	if err != nil {
		return errors.Wrap(err, "can't partse json")
	}
	return nil
}
