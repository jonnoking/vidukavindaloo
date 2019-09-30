package models

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
)

func NewPhasesFromByteArray(bootstrap []byte) (*Phases, error) {

	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
	}

	ts := []Phase{}

	var bs map[string]interface{}
	json.Unmarshal([]byte(bootstrap), &bs)

	//log.Printf("%+v", bs)

	pt := bs["phases"].([]interface{})

	for _, v := range pt {

		var phase Phase

		config.Result = &phase

		decoder, _ := mapstructure.NewDecoder(config)
		i, _ := v.(map[string]interface{})

		decoder.Decode(i)

		//mapstructure.WeakDecode(v, &player)

		ts = append(ts, phase)
	}

	r, e := NewPhases(ts)

	return r, e

}

func NewPhases(p []Phase) (*Phases, error) {

	pt := Phases{}

	pt.Phases = map[int]Phase{}

	for _, phase := range p {
		pt.Phases[phase.ID] = phase
	}

	return &pt, nil
}

func NewPhasesFromBootStrapMap(bootstrap map[string]interface{}) (*Phases, error) {

	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
	}

	ts := []Phase{}
	phases := bootstrap["phases"].([]interface{})

	for _, v := range phases {

		var phase Phase

		config.Result = &phase

		decoder, _ := mapstructure.NewDecoder(config)
		i, _ := v.(map[string]interface{})

		decoder.Decode(i)

		//mapstructure.WeakDecode(v, &player)

		ts = append(ts, phase)
	}

	r, e := NewPhases(ts)

	return r, e
}
