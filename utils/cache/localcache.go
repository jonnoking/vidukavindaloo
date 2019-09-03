package cache

import (
	"io"
	"io/ioutil"
)

// SaveBodyToFile save response body to file
func SaveBodyToFile(body io.ReadCloser, filename string) error {

	byteValue, readErr := ioutil.ReadAll(body)
	if readErr != nil {
		return readErr
	}

	fErr := ioutil.WriteFile(filename, byteValue, 0644)
	if fErr != nil {
		return fErr
	}
	return nil
}

// doesn't work - look at closures?
// func ReadJsonFromCache(i map[string]interface{}, filename string) (map[string]interface{}, error) {
// 	c := i
// 	cookies, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}

// 	json.Unmarshal(cookies, &c)

// 	return c, nil
// }
