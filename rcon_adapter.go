package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type RCONAdapter struct {
	client *rconClient
}

func NewRCONAdapter(c *rconClient) (*RCONAdapter, error) {

	return &RCONAdapter{client: c}, nil
}

type Player struct {
	Name string
	UUID uuid.UUID
}

func (a *RCONAdapter) GetPlayers() ([]Player, error) {
	playerListSentence, err := a.client.ListPlayers(true)
	if err != nil {
		return nil, err
	}

	// remove everything before the first :
	playerListSentence = strings.Split(playerListSentence, ": ")[1]
	players := make([]Player, 0)
	playerInfo := strings.Split(playerListSentence, ", ")

	for _, info := range playerInfo {
		playerData := strings.Split(info, " (")
		if len(playerData) != 2 {
			continue
		}

		name := playerData[0]
		uuidStr := strings.TrimSuffix(playerData[1], ")")

		uuid, err := uuid.Parse(uuidStr)
		if err != nil {
			continue
		}

		player := Player{Name: name, UUID: uuid}
		players = append(players, player)
	}

	return players, nil
}

type Coordinates struct {
	X float64
	Y float64
	Z float64
}

func (a *RCONAdapter) GetPlayerLocation(p Player) (Coordinates, error) {

	result, err := a.client.Execute(fmt.Sprintf("/data get entity @p[name=%s] Pos", p.Name))
	if err != nil {
		return Coordinates{}, err
	}

	// Parse the coordinates from the result
	coordsStr := strings.TrimPrefix(result, fmt.Sprintf("%s has the following entity data: [", p.Name))
	coordsStr = strings.TrimSuffix(coordsStr, "]")
	coordsData := strings.Split(coordsStr, ", ")

	if len(coordsData) != 3 {
		return Coordinates{}, fmt.Errorf("invalid coordinates data")
	}

	x, err := strconv.ParseFloat(strings.TrimSuffix(coordsData[0], "d"), 64)
	if err != nil {
		return Coordinates{}, err
	}

	y, err := strconv.ParseFloat(strings.TrimSuffix(coordsData[1], "d"), 64)
	if err != nil {
		return Coordinates{}, err
	}

	z, err := strconv.ParseFloat(strings.TrimSuffix(coordsData[2], "d"), 64)
	if err != nil {
		return Coordinates{}, err
	}

	coordinates := Coordinates{X: x, Y: y, Z: z}
	return coordinates, nil
}

func (a *RCONAdapter) GetPlayerHealth(p Player) (int, error) {
	result, err := a.client.Execute(fmt.Sprintf("/data get entity @p[name=%s] Health", p.Name))
	if err != nil {
		return 0, err
	}

	healthStr := strings.TrimPrefix(result, fmt.Sprintf("%s has the following entity data: ", p.Name))
	healthStr = strings.TrimSuffix(healthStr, "f") // Remove the trailing "f"
	healthStr = strings.Split(healthStr, ".")[0]   // Remove any extra data after the health value
	health, err := strconv.Atoi(healthStr)
	if err != nil {
		return 0, err
	}

	return health, nil
}

func (a *RCONAdapter) GivePlayerItem(p Player, item string, count int) error {
	_, err := a.client.Execute(fmt.Sprintf("/give %s %s %d", p.Name, item, count))
	if err != nil {
		return err
	}

	return nil
}

func (a *RCONAdapter) SetBlock(blockType string, location Coordinates) {
	a.client.SetBlock(int(location.X), int(location.Y), int(location.Z), blockType)
}

func (a *RCONAdapter) SpawnZombie(location Coordinates, zombieCount int) {
	a.client.SayMessage("Brace yourself, zombies are coming!!!")
	time.Sleep(3 * time.Second)
	for i := 0; i < zombieCount; i++ {
		summonResponse, err := a.client.SummonEntity("minecraft:zombie", location.X, location.Y, location.Z)
		if err != nil {
			log.Fatal("Error summoning entity:", err)
		}
		log.Println(summonResponse)
	}
}

func (a *RCONAdapter) SpawnVillager(location Coordinates) {
	a.client.SummonEntity("minecraft:villager", location.X, location.Y, location.Z)
}
