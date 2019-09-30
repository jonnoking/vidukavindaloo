package models

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"log"
)

// error - circular package reference
// func GetPlayerByCode(code int) (*Player, error) {
// 	var plr Player

// 	if fpl.Players == nil {
// 		return fmt.Errorf("Players collection not populated"), plr
// 	}

// 	plr := fpl.Players[code]
// 	plr.PlayerType := fpl.PlayerTypes[plr.PlayerTypeID]

// 	if (plr == nil) {
// 		return fmt.Errorf("Player not found"), plr
// 	}

// 	return plr, nil
// }

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
