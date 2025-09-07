package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorcon/rcon"
)

type RCONAdapter struct {
	rcon *rcon.Conn
}

func NewRCONAdapter(hostname string, port int, password string) (*RCONAdapter, error) {
	conn, err := rcon.Dial(hostname+":"+strconv.Itoa(port), password)
	if err != nil {
		return nil, err
	}
	return &RCONAdapter{rcon: conn}, nil
}

func (a *RCONAdapter) Close() error {
	return a.rcon.Close()
}

type Player struct {
	Name string    `json:"name"`
	UUID uuid.UUID `json:"uuid"`
}

func (a *RCONAdapter) ListPlayerNames() (map[string]Player, error) {

	playerListSentence, err := a.rcon.Execute("list uuids")
	if err != nil {
		return nil, err
	}
	// remove everything before the first :
	playerListSentence = strings.Split(playerListSentence, ": ")[1]
	players := make(map[string]Player)
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
		players[name] = player
	}

	return players, nil
}

type Coordinates struct {
	X float64
	Y float64
	Z float64
}

type BlockCoordinates struct {
	X int
	Y int
	Z int
}

func convertCoordinatesToBlockCoordinates(c Coordinates) BlockCoordinates {
	return BlockCoordinates{
		X: int(math.Floor(c.X)),
		Y: int(math.Floor(c.Y)),
		Z: int(math.Floor(c.Z)),
	}
}

func (a *RCONAdapter) GetPlayerLocation(p Player) (Coordinates, error) {

	result, err := a.rcon.Execute(fmt.Sprintf("data get entity @p[name=%s] Pos", p.Name))
	if err != nil {
		return Coordinates{}, fmt.Errorf("error with rcon execution: %w", err)
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
	result, err := a.rcon.Execute(fmt.Sprintf("/data get entity @p[name=%s] Health", p.Name))
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
	_, err := a.rcon.Execute(fmt.Sprintf("give %s %s %d", p.Name, item, count))
	if err != nil {
		return err
	}

	return nil
}

func (a *RCONAdapter) SpawnZombie(location Coordinates, zombieCount int) {
	a.Say("Brace yourself, zombies are coming!!!")
	time.Sleep(3 * time.Second)
	for i := 0; i < zombieCount; i++ {
		summonResponse, err := a.SummonEntity("minecraft:zombie", location.X, location.Y, location.Z)
		if err != nil {
			log.Fatal("Error summoning entity:", err)
		}
		log.Println(summonResponse)
	}
}

func (a *RCONAdapter) SpawnVillager(location Coordinates) {
	a.SummonEntity("minecraft:villager", location.X, location.Y, location.Z)
}

func (a *RCONAdapter) SummonLightning(location Coordinates) {
	a.SummonEntity("minecraft:lightning_bolt", location.X, location.Y, location.Z)
}

func (a *RCONAdapter) SendMessage(targets, message string) (string, error) {
	return a.rcon.Execute("msg " + targets + " " + message)
}

func (a *RCONAdapter) SummonEntity(entity string, x, y, z float64) (string, error) {
	return a.rcon.Execute("summon " + entity + " " + strconv.FormatFloat(x, 'f', -1, 64) + " " + strconv.FormatFloat(y, 'f', -1, 64) + " " + strconv.FormatFloat(z, 'f', -1, 64))
}

func (a *RCONAdapter) SetBlock(location BlockCoordinates, block string) (string, error) {
	return a.rcon.Execute("setblock " + strconv.Itoa(int(location.X)) + " " + strconv.Itoa(int(location.Y)) + " " + strconv.Itoa(int(location.Z)) + " " + block)
}

func (a *RCONAdapter) Fill(start, end BlockCoordinates, block string) (string, error) {
	return a.rcon.Execute("fill " + strconv.Itoa(int(start.X)) + " " + strconv.Itoa(int(start.Y)) + " " + strconv.Itoa(int(start.Z)) + " " + strconv.Itoa(int(end.X)) + " " + strconv.Itoa(int(end.Y)) + " " + strconv.Itoa(int(end.Z)) + " " + block)
}

func (a *RCONAdapter) Say(message string) (string, error) {
	return a.rcon.Execute("say " + message)
}

func (a *RCONAdapter) GiveItem(targets, item string, count int) (string, error) {
	return a.rcon.Execute("give " + targets + " " + item + " " + strconv.Itoa(count))
}

func (a *RCONAdapter) CreativeMode(player string) (string, error) {
	return a.rcon.Execute("/gamemode creative " + player)
}

func (a *RCONAdapter) SurvivalMode(player string) (string, error) {
	return a.rcon.Execute("/gamemode survival " + player)
}
