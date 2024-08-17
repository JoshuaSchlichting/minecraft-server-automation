package main

import (
	"log"
)

func main() {
	client, err := newRCONClient("192.168.1.xxx", 25575, "rcon_password")
	if err != nil {
		log.Fatal("Error creating RCON adapter:", err)
	}
	defer client.Close()

	rcon, err := NewRCONAdapter(client)
	if err != nil {
		log.Fatal("Error creating RCON adapter:", err)
	}

	client.PrintConnectedPlayers()
	rcon.DiamondRoulette()
	SayRandomSnappleFacts(client)
	// playerLocation, err := rcon.GetPlayerLocation(players[0])
	// if err != nil {
	// 	log.Fatal("Error getting player location:", err)
	// }

	// client.GiveItem("playername", "minecraft:villager_spawn_egg", 2)
	// rcon.SetBlock("minecraft:gold_block", Coordinates{X: 0, Y: 0, Z: 0})

	// run forever until someone hit's ctrl+c
	select {}
}
