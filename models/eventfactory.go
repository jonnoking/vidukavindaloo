package models

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
)

func NewEvents(teams []Event) (*Events, error) {
	ts := map[int]Event{}
	for _, event := range teams {
		ts[event.ID] = event
	}

	t := new(Events)
	t.Events = ts

	return t, nil
}

func NewEventsFromBootStrap(filename string) (*Events, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	r, e := NewEventsFromBootStrapByteArray(f)

	return r, e
}

func NewEventsFromBootStrapByteArray(bootstrap []byte) (*Events, error) {

	var result map[string]interface{}
	json.Unmarshal([]byte(bootstrap), &result)

	r, e := NewEventsFromBootStrapMap(result)

	return r, e
}

func NewEventsFromBootStrapMap(bootstrap map[string]interface{}) (*Events, error) {

	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
	}

	ts := []Event{}
	teams := bootstrap["events"].([]interface{})

	for _, v := range teams {
		var event Event

		config.Result = &event

		decoder, _ := mapstructure.NewDecoder(config)
		i, _ := v.(map[string]interface{})

		decoder.Decode(i)

		//mapstructure.Decode(v, &team)
		ts = append(ts, event)
	}

	r, e := NewEvents(ts)

	return r, e
}
