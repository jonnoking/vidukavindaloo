package models

import (
	"encoding/json"
	//"fmt"
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

	// config := &mapstructure.DecoderConfig{
	// 	TagName: "json",
	// }

	ts := []FPLPlayer{}
	players := bootstrap["elements"].([]interface{})

	for _, v := range players {

		var player FPLPlayer

		// decoder, _ := mapstructure.NewDecoder(config)
		// config.Result = &player
		// i, ok := v.(map[string]interface{})
		// if !ok {
		// 	log.Fatal(fmt.Sprintf("%#+v", v))
		// }
		// decoder.Decode(i)

		mapstructure.WeakDecode(v, &player)

		ts = append(ts, player)
		//log.Println(player)
	}

	r, e := NewFPLPlayers(ts)

	return r, e
}
