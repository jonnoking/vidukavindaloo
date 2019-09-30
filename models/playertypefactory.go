package models

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
)

func NewPlayerTypesFromByteArray(bootstrap []byte) (*PlayerTypes, error) {

	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
	}

	ts := []PlayerType{}

	var bs map[string]interface{}
	json.Unmarshal([]byte(bootstrap), &bs)

	//log.Printf("%+v", bs)

	pt := bs["element_types"].([]interface{})

	for _, v := range pt {

		var playerType PlayerType

		config.Result = &playerType

		decoder, _ := mapstructure.NewDecoder(config)
		i, _ := v.(map[string]interface{})

		decoder.Decode(i)

		//mapstructure.WeakDecode(v, &player)

		ts = append(ts, playerType)
	}

	r, e := NewPlayerTypes(ts)

	return r, e

}

func NewPlayerTypes(p []PlayerType) (*PlayerTypes, error) {

	pt := PlayerTypes{
		PlayerTypes: p,
	}

	pt.Positions = map[int]PlayerType{}

	for _, player := range p {
		pt.Positions[player.ID] = player
	}

	return &pt, nil
}

func NewPlayerTypesFromBootStrapMap(bootstrap map[string]interface{}) (*PlayerTypes, error) {

	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
	}

	ts := []PlayerType{}
	playerTypes := bootstrap["element_types"].([]interface{})

	for _, v := range playerTypes {

		var playerType PlayerType

		config.Result = &playerType

		decoder, _ := mapstructure.NewDecoder(config)
		i, _ := v.(map[string]interface{})

		decoder.Decode(i)

		//mapstructure.WeakDecode(v, &player)

		ts = append(ts, playerType)
	}

	r, e := NewPlayerTypes(ts)

	return r, e
}
