package models

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"log"
)

func NewFPLPlayers(players []FPLPlayer) (*FPLPlayers, error) {
	ts := map[int]FPLPlayer{}
	for _, player := range players {
		ts[player.Code] = player
	}

	t := new(FPLPlayers)
	t.Players = ts
	log.Println(len(t.Players)) //working
	return t, nil
}

func NewFPLPlayersFromBootStrap(filename string) (*FPLPlayers, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	r, e := NewFPLPlayersFromBootStrapByteArray(f)

	return r, e
}

func NewFPLPlayersFromBootStrapByteArray(bootstrap []byte) (*FPLPlayers, error) {

	var result map[string]interface{}
	json.Unmarshal([]byte(bootstrap), &result)

	r, e := NewFPLPlayersFromBootStrapMap(result)

	return r, e
}

func NewFPLPlayersFromBootStrapMap(bootstrap map[string]interface{}) (*FPLPlayers, error) {

	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
	}

	ts := []FPLPlayer{}
	players := bootstrap["elements"].([]interface{})

	for _, v := range players {

		var player FPLPlayer

		config.Result = &player

		decoder, _ := mapstructure.NewDecoder(config)
		i, _ := v.(map[string]interface{})

		decoder.Decode(i)

		//mapstructure.WeakDecode(v, &player)

		ts = append(ts, player)
	}

	r, e := NewFPLPlayers(ts)

	return r, e
}
