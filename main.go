package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/JoshuaSchlichting/minecraft-server-automation/logger"
)

func main() {
	create := flag.Bool("create", false, "A boolean flag to create the structure")
	destroy := flag.Bool("destroy", false, "A boolean flag to destroy the structure")
	give := flag.String("give", "", "give this item to the player (player:item), trailed by :<count> if more than one is desired")
	creativePlayer := flag.String("creative-mode", "", "enable creative mode for the player")
	survivalPlayer := flag.String("survival-mode", "", "enable survivalmode for the player")
	flag.Parse()

	if *create && *destroy {
		log.Fatal("Flags 'create' and 'destroy' are mutually exclusive")
	}
	log.SetLogLevel(log.DEBUG)

	rcon, err := NewRCONAdapter("192.168.50.29", 25575, "ALEXB")
	if err != nil {
		log.Fatal("error creating RCON adapter:", err)
	}
	defer rcon.Close()
	if *create {
		blockStartingPosition := BlockCoordinates{
			X: 434,
			Y: 100,
			Z: 395,
		}
		buildHouse(rcon, blockStartingPosition, "testblocks.txt")
		os.Exit(0)
	}
	if *destroy {
		deleteBlocks(rcon, "testblocks.txt")
		os.Exit(0)
	}
	if *give != "" {
		parts := strings.Split(*give, ":")
		if len(parts) < 2 {
			log.Error("Invalid format for 'give' flag. Expected playername:itemname[:count]")
			return
		}

		playerName := parts[0]
		itemName := parts[1]
		count := 1
		if len(parts) > 2 {
			count, _ = strconv.Atoi(parts[2])
		}
		log.Info(fmt.Sprintf("Giving %d %s to %s", count, itemName, playerName))
		_, err := rcon.GiveItem(playerName, itemName, count)
		if err != nil {
			log.Error("Error giving item:", err)
		}
		os.Exit(0)
	}

	if *creativePlayer != "" {
		log.Info(fmt.Sprintf("Enabling creative mode for player %s", *creativePlayer))
		_, err := rcon.CreativeMode(*creativePlayer)
		if err != nil {
			log.Error("Error enabling creative mode:", err)
		}
		os.Exit(0)
	}
	if *survivalPlayer != "" {
		log.Info(fmt.Sprintf("Enabling survival mode for player %s", *survivalPlayer))
		_, err := rcon.SurvivalMode(*survivalPlayer)
		if err != nil {
			log.Error("Error enabling survival mode:", err)
		}
		os.Exit(0)
	}
	services := NewService(rcon)
	services.StartPrintConnectedPlayers()
	services.StartDiamondRoulette()
	services.StartRandomSnappleFacts()
	// 	// services.StartZombieHordeRaid(Coordinates{X: 375, Y: 63, Z: 537})
	// 	// services.StartLightningStorms()
	services.StartMineRailGiveaway()
	// 	// _, err = client.GiveItem("SchlitzMaltLiqy", `minecraft:glass`, 64)
	// 	// _, err = rcon.GiveItem("KillerKora", `minecraft:sheep_spawn_egg`, 64)
	// if err != nil {
	// 	log.Error("Error giving item:", err)
	// }
	// GiveEnchantedDiamondArmorSet("KillerKora", client)
	players, err := rcon.ListPlayerNames()
	if err != nil {
		log.Fatal("Error getting players:", err)
	}
	log.Info(players)

	killerKorasPos, err := rcon.GetPlayerLocation(players["KillerKora"])
	if err != nil {
		log.Error("error getting player location:", err)
		return
	}
	log.Info(killerKorasPos)
	// blockStartingPosition := BlockCoordinates{
	// 	X: 434,
	// 	Y: 100,
	// 	Z: 390,
	// }
	// buildHouse(rcon, blockStartingPosition, "testblocks.txt")
	// deleteBlocks(rcon, "testblocks.txt")
	// run forever until someone hit's ctrl+c
	// zombieAttackOnPlayer(rcon, players["KillerKora"])
	log.Info("Server is running...")

	// BuildRail(*rcon, East, BlockCoordinates{585, 68, -2493}, BlockCoordinates{6100, 68, -2493})

	select {}
}

func zombieAttackOnPlayer(rcon *RCONAdapter, player Player) {
	// Find the coordinates of the player
	log.Info(player.Name, "is being attacked by zombies!")
	playerLocation, err := rcon.GetPlayerLocation(player)
	if err != nil {
		log.Error("Error getting player location:", err)
		return
	}

	rcon.SendMessage(player.Name, "It's close to midnight...")
	time.Sleep(time.Duration(3) * time.Second)
	rcon.SendMessage(player.Name, "Something evil's lurking in the dark...")
	time.Sleep(time.Duration(3) * time.Second)
	rcon.SpawnZombie(playerLocation, 1)
	rcon.SendMessage(player.Name, "Under the moonlight...")
	time.Sleep(time.Duration(2) * time.Second)
	rcon.SendMessage(player.Name, "You see a sight that almost stops your heart...")
	time.Sleep(time.Duration(2) * time.Second)
	rcon.SendMessage(player.Name, "You try to scream...")
	time.Sleep(time.Duration(2) * time.Second)
	rcon.SpawnZombie(playerLocation, 1)
	rcon.SendMessage(player.Name, "But terror takes the sound before you make it...")
	time.Sleep(time.Duration(2) * time.Second)
	rcon.SendMessage(player.Name, "You start to freeze...")
	time.Sleep(time.Duration(2) * time.Second)
	rcon.SpawnZombie(playerLocation, 1)

	rcon.SendMessage(player.Name, "As horror looks you right between the eyes...")
	time.Sleep(time.Duration(2) * time.Second)
	rcon.SendMessage(player.Name, "You're paralyzed...")
	rcon.SpawnZombie(playerLocation, 1)

	time.Sleep(time.Duration(2) * time.Second)
	rcon.SendMessage(player.Name, "Cause this is THRILLER! THRILLER NIGHT!!!...")
	rcon.SpawnZombie(playerLocation, 1)

}

func (r *RCONAdapter) GiveEnchantedDiamondArmorSet(playerName string) {
	r.GiveItem(playerName, "minecraft:diamond_helmet", 1)
	r.GiveItem(playerName, "minecraft:diamond_chestplate", 1)
	r.GiveItem(playerName, "minecraft:diamond_leggings", 1)
	r.GiveItem(playerName, "minecraft:diamond_boots", 1)
	r.GiveItem(playerName, "minecraft:diamond_sword", 1)
	r.GiveItem(playerName, "minecraft:shield", 1)
	r.GiveItem(playerName, "minecraft:bow", 1)
	r.GiveItem(playerName, "minecraft:arrow", 64)
	r.GiveItem(playerName, "minecraft:diamon_pickaxe", 1)
	r.GiveItem(playerName, "minecraft:diamond_axe", 1)
	r.GiveItem(playerName, "minecraft:diamond_shovel", 1)
	r.GiveItem(playerName, "minecraft:diamond_hoe", 1)
	r.GiveItem(playerName, "minecraft:golden_apple", 1)
	r.GiveItem(playerName, "minecraft:cooked_beef", 64)
}

func buildHouse(rcon *RCONAdapter, c BlockCoordinates, persistentBlockLocationFilename string) (spawnedBlocks []BlockCoordinates) {

	rcon.SetBlock(c, "minecraft:stone")
	spawnedBlocks = append(spawnedBlocks, c)
	spawnedBlocks = append(spawnedBlocks, buildHouseFloor(rcon, c)...)
	spawnedBlocks = append(spawnedBlocks, buildHouseRoof(rcon, c)...)
	spawnedBlocks = append(spawnedBlocks, buildHouseWall(rcon, c, 5, 3)...)
	spawnedBlocks = append(spawnedBlocks, buildHouseDoor(rcon, c)...)
	spawnedBlocks = append(spawnedBlocks, buildHouseWindow(rcon, c)...)

	saveSpawnedBlocksToFile(spawnedBlocks, persistentBlockLocationFilename)
	return spawnedBlocks
}

func deleteBlocks(rcon *RCONAdapter, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		block := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(block[0])
		y, _ := strconv.Atoi(block[1])
		z, _ := strconv.Atoi(block[2])
		rcon.SetBlock(BlockCoordinates{x, y, z}, "minecraft:air")
	}
}

func saveSpawnedBlocksToFile(blocks []BlockCoordinates, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer file.Close()

	for _, block := range blocks {
		_, err := file.WriteString(fmt.Sprintf("%d,%d,%d\n", block.X, block.Y, block.Z))
		if err != nil {
			log.Fatal("Error writing to file:", err)
		}
	}
}

func buildHouseFloor(rcon *RCONAdapter, c BlockCoordinates) []BlockCoordinates {
	floorBlocks := []BlockCoordinates{}
	for x := c.X - 2; x <= c.X+2; x++ {
		for z := c.Z - 2; z <= c.Z+2; z++ {
			rcon.SetBlock(BlockCoordinates{x, c.Y - 1, z}, "minecraft:stone")
			floorBlocks = append(floorBlocks, BlockCoordinates{x, c.Y - 1, z})
		}
	}
	return floorBlocks
}

func buildHouseRoof(rcon *RCONAdapter, c BlockCoordinates) []BlockCoordinates {
	roofBlocks := []BlockCoordinates{}
	for x := c.X - 2; x <= c.X+2; x++ {
		for z := c.Z - 2; z <= c.Z+2; z++ {
			rcon.SetBlock(BlockCoordinates{x, c.Y, z}, "minecraft:stone")
			roofBlocks = append(roofBlocks, BlockCoordinates{x, c.Y, z})
		}
	}
	return roofBlocks
}

func buildHouseWall(rcon *RCONAdapter, bottomLeft BlockCoordinates, length, height int) []BlockCoordinates {
	wallBlocks := []BlockCoordinates{}
	for x := int(bottomLeft.X); x < int(bottomLeft.X)+length; x++ {
		for y := int(bottomLeft.Y); y < int(bottomLeft.Y)+height; y++ {
			rcon.SetBlock(BlockCoordinates{x, y, bottomLeft.Z}, "minecraft:stone")
			wallBlocks = append(wallBlocks, BlockCoordinates{x, y, bottomLeft.Z})
		}
	}
	return wallBlocks
}

func buildHouseDoor(rcon *RCONAdapter, c BlockCoordinates) []BlockCoordinates {
	doorBlocks := []BlockCoordinates{}
	for x := c.X - 1; x <= c.X+1; x++ {
		for z := c.Z - 1; z <= c.Z+1; z++ {
			rcon.SetBlock(BlockCoordinates{x, c.Y + 2, z}, "minecraft:wooden_door")
			doorBlocks = append(doorBlocks, BlockCoordinates{x, c.Y + 2, z})
		}
	}
	return doorBlocks
}

func buildHouseWindow(rcon *RCONAdapter, c BlockCoordinates) []BlockCoordinates {
	windowBlocks := []BlockCoordinates{}
	for x := c.X - 1; x <= c.X+1; x++ {
		for z := c.Z - 1; z <= c.Z+1; z++ {
			rcon.SetBlock(BlockCoordinates{x, c.Y + 3, z}, "minecraft:glass")
			windowBlocks = append(windowBlocks, BlockCoordinates{x, c.Y + 3, z})
		}
	}
	return windowBlocks
}
