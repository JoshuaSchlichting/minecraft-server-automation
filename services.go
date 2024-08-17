package main

import (
	"log"
	"time"

	"golang.org/x/exp/rand"
)

// Randomly distrubutes a diamond to players every rand.Intn(25) minutes
func (a *RCONAdapter) DiamondRoulette() {
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(25)) * time.Minute)
			players, err := a.GetPlayers()
			if err != nil {
				log.Fatal("Error getting players:", err)
			}

			for _, player := range players {
				if rand.Intn(2) == 0 {
					continue
				}
				a.client.SendMessage(player.Name, "For your troubles, here's a diamond. Don't spend it all in one place, kid!")
				a.GivePlayerItem(player, "minecraft:diamond", 1)
				log.Println(player.Name, " won a diamond!")
			}
		}
	}()
}

func (c *rconClient) PrintConnectedPlayers() {
	func() {
		for {
			players, err := c.ListPlayers(false)
			if err != nil {
				log.Println("Error getting players:", err)
				continue
			}

			log.Println(players)
			time.Sleep(5 * time.Minute)
		}
	}()
}

func SayRandomSnappleFacts(client *rconClient) {
	go func() {
		for {
			randomBool := rand.Intn(2) == 0
			if randomBool {
				continue
			}
			client.Say(getRandomSnappleFact())
			time.Sleep(15 * time.Minute)
		}
	}()
}
