package models

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"log"
)

func NewPlayers(players []Player) (*Players, error) {
	ts := map[int]Player{}
	tid := map[int]Player{}
	for _, player := range players {
		ts[player.Code] = player
		tid[player.ID] = player
	}

	t := new(Players)
	t.Players = ts
	t.PlayersByID = tid
	log.Println(len(t.Players)) //working
	return t, nil
}

func NewPlayersFromBootStrap(filename string) (*Players, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	r, e := NewPlayersFromBootStrapByteArray(f)

	return r, e
}

func NewPlayersFromBootStrapByteArray(bootstrap []byte) (*Players, error) {

	var result map[string]interface{}
	json.Unmarshal([]byte(bootstrap), &result)

	r, e := NewPlayersFromBootStrapMap(result)

	return r, e
}

func NewPlayersFromBootStrapMap(bootstrap map[string]interface{}) (*Players, error) {

	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
	}

	ts := []Player{}
	players := bootstrap["elements"].([]interface{})

	for _, v := range players {

		var player Player

		config.Result = &player

		decoder, _ := mapstructure.NewDecoder(config)
		i, _ := v.(map[string]interface{})

		decoder.Decode(i)

		//mapstructure.WeakDecode(v, &player)

		ts = append(ts, player)
	}

	r, e := NewPlayers(ts)

	return r, e
}

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
