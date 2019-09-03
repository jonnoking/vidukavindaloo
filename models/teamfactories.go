package models

import (
	"encoding/json"
	//	"fmt"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	//	"log"
)

func NewFPLTeams(teams []FPLTeam) (*FPLTeams, error) {
	ts := map[int]FPLTeam{}
	for _, team := range teams {
		ts[team.Code] = team
	}

	t := new(FPLTeams)
	t.Teams = ts

	return t, nil
}

func NewFPLTeamsFromBootStrap(filename string) (*FPLTeams, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	r, e := NewFPLTeamsFromBootStrapByteArray(f)

	return r, e
}

func NewFPLTeamsFromBootStrapByteArray(bootstrap []byte) (*FPLTeams, error) {

	var result map[string]interface{}
	json.Unmarshal([]byte(bootstrap), &result)

	r, e := NewFPLTeamsFromBootStrapMap(result)

	return r, e
}

func NewFPLTeamsFromBootStrapMap(bootstrap map[string]interface{}) (*FPLTeams, error) {

	// config := &mapstructure.DecoderConfig{
	// 	TagName: "json",
	// }

	ts := []FPLTeam{}
	teams := bootstrap["teams"].([]interface{})

	for _, v := range teams {
		var team FPLTeam

		// decoder, _ := mapstructure.NewDecoder(config)
		// config.Result = &team
		// i, ok := v.(map[string]interface{})
		// if !ok {
		// 	log.Fatal(fmt.Sprintf("%#+v", v))
		// }
		// decoder.Decode(i)

		mapstructure.Decode(v, &team)
		ts = append(ts, team)
	}

	r, e := NewFPLTeams(ts)

	return r, e
}
