package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func (api *API) ExecuteFPLGet(url string) ([]byte, error) {

	apiURL := url

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	r, _ := api.BuildFPLRequest(apiURL, "GET")

	resp, respErr := client.Do(r)
	if respErr != nil {
		return nil, respErr
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("ExecuteFPLGet : status code: %d - %s \n %s", resp.StatusCode, resp.Status, url)
	}

	byteValue, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		//log.Fatal(readErr)
		return nil, readErr
	}

	return byteValue, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
