package main

import (
	"time"

	log "github.com/JoshuaSchlichting/minecraft-server-automation/logger"
	"golang.org/x/exp/rand"
)

type Service struct {
	rcon *RCONAdapter
}

func NewService(r *RCONAdapter) *Service {
	return &Service{rcon: r}
}

// StartZombieRaidBase spawns a zombie horde at the base every 20 minutes
func (gl *Service) StartZombieHordeRaid(c Coordinates) {
	go func() {
		time.Sleep(time.Duration(20) * time.Minute)
		gl.rcon.client.Say("A zombie horde is attacking the base!")
		// Spawn a zombie horde at the base
		for i := 0; i < 10; i++ {
			gl.rcon.SpawnZombie(c, 5)
		}
	}()
}

// Randomly distrubutes a diamond to players every rand.Intn(25) minutes
func (gl *Service) StartDiamondRoulette() {
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(25)) * time.Minute)
			players, err := gl.rcon.GetPlayers()
			if err != nil {
				log.Fatal("Error getting players:", err)
			}

			for _, player := range players {
				if rand.Intn(2) == 0 {
					continue
				}
				gl.rcon.client.SendMessage(player.Name, "For your troubles, here's a diamond. Don't spend it all in one place, kid!")
				gl.rcon.GivePlayerItem(player, "minecraft:diamond", 1)
				log.Info(player.Name, " won a diamond!")
			}
		}
	}()
}

// StartRandomSnappleFacts sends a random Snapple fact to the server every 15 minutes based on a coin toss each time
func (gl *Service) StartRandomSnappleFacts() {
	go func() {
		for {
			randomBool := rand.Intn(2) == 0
			if randomBool {
				continue
			}
			gl.rcon.client.Say(getRandomSnappleFact())
			time.Sleep(15 * time.Minute)
		}
	}()
}

// StartPrintConnectedPlayers prints the connected players every 5 minutes
func (gl *Service) StartPrintConnectedPlayers() {
	go func() {
		for {
			players, err := gl.rcon.client.ListPlayers(false)
			if err != nil {
				log.Error("Error getting players:", err)
				continue
			}

			log.Info(players)
			time.Sleep(5 * time.Minute)
		}
	}()
}

func (s *Service) StartLightningStorms() {
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(60)) * time.Minute)

			players, err := s.rcon.GetPlayers()
			if len(players) == 0 {
				if err != nil {
					log.Error("Error getting players:", err)
				}
				continue
			}

			for playerName, player := range players {
				s.rcon.client.SendMessage(playerName, "A lightning storm is approaching!")
				time.Sleep(time.Duration(rand.Intn(15)) * time.Second)
				playerLocation, err := s.rcon.GetPlayerLocation(player)
				if err != nil {
					log.Error("Error getting player location:", err)
					continue
				}

				s.rcon.client.Say("A lightning storm is approaching!")
				s.rcon.client.Say("A lighting storm!!!")

				// Generate random offsets for lightning strikes
				offsetX := float64(rand.Intn(15))
				offsetZ := float64(rand.Intn(10))

				// Calculate the coordinates near the player location
				lightningX := playerLocation.X + offsetX
				lightningZ := playerLocation.Z + offsetZ

				// Summon lightning strikes near the player location
				s.rcon.SummonLightning(Coordinates{X: lightningX, Y: playerLocation.Y, Z: lightningZ})
				time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

				s.rcon.SummonLightning(Coordinates{X: lightningX, Y: playerLocation.Y, Z: lightningZ})
				time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

				s.rcon.SummonLightning(Coordinates{X: lightningX, Y: playerLocation.Y, Z: lightningZ})
				time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

				s.rcon.SummonLightning(Coordinates{X: lightningX, Y: playerLocation.Y, Z: lightningZ})
				time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

				s.rcon.SummonLightning(Coordinates{X: lightningX, Y: playerLocation.Y, Z: lightningZ})
				time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

				s.rcon.SummonLightning(Coordinates{X: lightningX, Y: playerLocation.Y, Z: lightningZ})
				time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

				s.rcon.SummonLightning(Coordinates{X: lightningX, Y: playerLocation.Y, Z: lightningZ})
				time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

				s.rcon.SummonLightning(Coordinates{X: lightningX, Y: playerLocation.Y, Z: lightningZ})
				time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

				s.rcon.SummonLightning(Coordinates{X: lightningX, Y: playerLocation.Y, Z: lightningZ})
				time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			}
		}
	}()
}
