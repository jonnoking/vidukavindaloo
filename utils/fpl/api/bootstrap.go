package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	// "net/http"
	// "time"
)

func (api *API) LoadBootsrapFromCache() ([]byte, error) {
	f, err := ioutil.ReadFile(api.Config.Files.GetBootstrapFilename())
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (api *API) RefreshBootstrap() map[string]interface{} {

	byteValue, readErr := api.ExecuteFPLGet(api.Config.API.GetBoostrapAPI())
	if readErr != nil {
		log.Fatal(readErr)
	}

	api.Config.Files.SaveByteArrayToFile(byteValue, api.Config.Files.GetBootstrapFilename())

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return result
}
