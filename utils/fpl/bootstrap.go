package fpl

import (
	"encoding/json"
	"github.com/jonnoking/vidukavindaloo/utils/cache"
	"github.com/jonnoking/vidukavindaloo/utils/config"
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
	// fplBootsrapURL := "https://fantasy.premierleague.com/api/bootstrap-static/"

	// fplClient := http.Client{
	// 	Timeout: time.Second * 5,
	// }

	// req, err := http.NewRequest(http.MethodGet, fplBootsrapURL, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// req.Header.Set("User-Agent", "cloudjonno")

	// res, getErr := fplClient.Do(req)
	// if getErr != nil {
	// 	log.Fatal(getErr)
	// }

	// sbErr := cache.SaveBodyToFile(res.Body, "./fpl-bootstrap.json")
	// if sbErr != nil {
	// 	log.Fatal(sbErr)
	// }

	// defer res.Body.Close()

	// byteValue, readErr := ioutil.ReadAll(res.Body)
	// if readErr != nil {
	// 	log.Fatal(readErr)
	// }

	byteValue, readErr := ExecuteFPLGet(config.GetBootstrapAPI())
	if readErr != nil {
		log.Fatal(readErr)
	}

	cache.SaveByteArrayToFile(byteValue, config.GetBootrapFilename())

	// fErr := ioutil.WriteFile("./fpl-bootstrap.json", byteValue, 0644)
	// check(fErr)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	//

	return result
}
