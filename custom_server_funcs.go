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
func (s *Service) StartZombieHordeRaid(c Coordinates) {
	log.Debug("Starting zombie horde raid service...")
	go func() {
		time.Sleep(time.Duration(20) * time.Minute)
		s.rcon.client.Say("A zombie horde is attacking the base!")
		// Spawn a zombie horde at the base
		for i := 0; i < 10; i++ {
			s.rcon.SpawnZombie(c, 5)
		}
	}()
}

// Randomly distrubutes a diamond to players every rand.Intn(25) minutes
func (s *Service) StartDiamondRoulette() {
	log.Debug("Starting diamond roulette service...")
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(25)) * time.Minute)
			players, err := s.rcon.GetPlayers()
			if err != nil {
				log.Fatal("Error getting players:", err)
			}

			for _, player := range players {
				if rand.Intn(2) == 0 {
					continue
				}
				s.rcon.client.SendMessage(player.Name, "For your troubles, here's a diamond. Don't spend it all in one place, kid!")
				s.rcon.GivePlayerItem(player, "minecraft:diamond", 1)
				log.Info(player.Name, " won a diamond!")
			}
		}
	}()
}

// StartRandomSnappleFacts sends a random Snapple fact to the server every 15 minutes based on a coin toss each time
func (s *Service) StartRandomSnappleFacts() {
	log.Debug("Starting random Snapple fact service...")
	go func() {
		for {
			randomBool := rand.Intn(2) == 0
			if randomBool {
				continue
			}
			s.rcon.client.Say(getRandomSnappleFact())
			time.Sleep(15 * time.Minute)
		}
	}()
}

// StartPrintConnectedPlayers prints the connected players every 5 minutes
func (s *Service) StartPrintConnectedPlayers() {
	go func() {
		for {
			players, err := s.rcon.client.ListPlayers(false)
			if err != nil {
				log.Error("Error getting players:", err)
				continue
			}

			log.Info(players)
			time.Sleep(5 * time.Minute)
		}
	}()
}

// StartLightningStorms sends a message to players that a lightning storm is approaching and then summons lightning strikes near the player location
func (s *Service) StartLightningStorms() {
	log.Debug("Starting lightning storm service...")
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(60)) * time.Minute)
			// simulate coin toss
			if rand.Intn(2) == 0 {
				continue
			}
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
				// Summon lightning strikes near the player location
				for range 50 {
					s.rcon.SummonLightning(Coordinates{X: playerLocation.X + float64(rand.Intn(20)), Y: playerLocation.Y, Z: playerLocation.Z + float64(rand.Intn(20))})
					time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
				}

			}
		}
	}()
}

// StartMineRailGiveaway randomly selects players to receive minecart and rail equipment
func (s *Service) StartMineRailGiveaway() {
	log.Debug("Starting minecart giveaway service...")
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(60)) * time.Minute)
			players, err := s.rcon.GetPlayers()
			if err != nil {
				log.Error("Error getting players:", err)
				continue
			}

			for _, player := range players {
				if rand.Intn(2) == 0 {
					continue
				}
				s.rcon.client.SendMessage(player.Name, "You have been selected to receive free minecart equipment! I suggest you built a crate, soon!!!")
				time.Sleep(time.Duration(20) * time.Second)
				s.rcon.GivePlayerItem(player, "minecraft:minecart", 3)
				s.rcon.GivePlayerItem(player, "minecraft:rail", 128)
				s.rcon.GivePlayerItem(player, "minecraft:powered_rail", 64)
				s.rcon.GivePlayerItem(player, "minecraft:redstone_torch", 64)
				log.Info(player.Name, " won minecart equipment!")
			}
		}
	}()
}
