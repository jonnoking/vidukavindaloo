package models

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
)

func NewTeams(teams []Team) (*Teams, error) {
	ts := map[int]Team{}
	for _, team := range teams {
		ts[team.Code] = team
	}

	t := new(Teams)
	t.TeamsByCode = ts
	t.Teams = teams

	return t, nil
}

func NewTeamsFromBootStrap(filename string) (*Teams, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	r, e := NewTeamsFromBootStrapByteArray(f)

	return r, e
}

func NewTeamsFromBootStrapByteArray(bootstrap []byte) (*Teams, error) {

	var result map[string]interface{}
	json.Unmarshal([]byte(bootstrap), &result)

	r, e := NewTeamsFromBootStrapMap(result)

	return r, e
}

func NewTeamsFromBootStrapMap(bootstrap map[string]interface{}) (*Teams, error) {

	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
	}

	ts := []Team{}
	teams := bootstrap["teams"].([]interface{})

	for _, v := range teams {
		var team Team

		config.Result = &team

		decoder, _ := mapstructure.NewDecoder(config)
		i, _ := v.(map[string]interface{})

		decoder.Decode(i)

		//mapstructure.Decode(v, &team)
		ts = append(ts, team)
	}

	r, e := NewTeams(ts)

	return r, e
}
