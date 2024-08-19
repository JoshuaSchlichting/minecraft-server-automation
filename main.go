package main

import (
	"time"

	log "github.com/JoshuaSchlichting/minecraft-server-automation/logger"
)

func main() {

	log.SetLogLevel(log.DEBUG)
	client, err := newRCONClient("192.168.1.xxx", 25575, "rcon_password")
	if err != nil {
		log.Fatal("Error creating RCON adapter:", err)
	}
	defer client.Close()

	rcon, err := NewRCONAdapter(client)
	if err != nil {
		log.Fatal("Error creating RCON adapter:", err)
	}
	services := NewService(rcon)
	services.StartPrintConnectedPlayers()
	services.StartDiamondRoulette()
	services.StartRandomSnappleFacts()
	services.StartZombieHordeRaid(Coordinates{X: 375, Y: 63, Z: 537})
	services.StartLightningStorms()
	services.StartMineRailGiveaway()
	// if err != nil {
	// 	log.Fatal("Error getting player location:", err)
	// }

	// client.GiveItem("playername", "minecraft:villager_spawn_egg", 2)
	// rcon.SetBlock("minecraft:gold_block", Coordinates{X: 0, Y: 0, Z: 0})

	// run forever until someone hit's ctrl+c
	select {}
}
