package api

import (
	"encoding/json"
	"github.com/jonnoking/vidukavindaloo/utils/cache"
	"github.com/jonnoking/vidukavindaloo/utils/fpl/config"
	"io/ioutil"
	"log"
	// "net/http"
	// "time"
)

func LoadBootsrapFromCache() ([]byte, error) {
	f, err := ioutil.ReadFile("./fpl-bootstrap.json")
	if err != nil {
		return nil, err
	}
	return f, nil
}

func RefreshBootstrap() map[string]interface{} {

	byteValue, readErr := ExecuteFPLGet(config.GetBoostrapAPI())
	if readErr != nil {
		log.Fatal(readErr)
	}

	cache.SaveByteArrayToFile(byteValue, config.GetBootstrapFilename())

	// fErr := ioutil.WriteFile("./fpl-bootstrap.json", byteValue, 0644)
	// check(fErr)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	//

	return result
}
