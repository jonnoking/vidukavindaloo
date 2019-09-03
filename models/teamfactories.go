package models

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
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

	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
	}

	ts := []FPLTeam{}
	teams := bootstrap["teams"].([]interface{})

	for _, v := range teams {
		var team FPLTeam

		config.Result = &team

		decoder, _ := mapstructure.NewDecoder(config)
		i, _ := v.(map[string]interface{})

		decoder.Decode(i)

		//mapstructure.Decode(v, &team)
		ts = append(ts, team)
	}

	r, e := NewFPLTeams(ts)

	return r, e
}
