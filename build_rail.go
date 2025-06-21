package main

import (
	"math"

	log "github.com/JoshuaSchlichting/minecraft-server-automation/logger"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

// String method to convert Direction to a human-readable string
func (d Direction) String() string {
	switch d {
	case North:
		return "north_south"
	case East:
		return "east_west"
	case South:
		return "north_south"
	case West:
		return "east_west"
	default:
		return "InvalidDirection"
	}
}

func BuildRail(rcon RCONAdapter, direction Direction, start BlockCoordinates, end BlockCoordinates) {
	railBlocks := buildRail(start, end)
	log.Info("Setting rails")
	blockRangeIndex := 0
	for _, block := range railBlocks {
		var blockLeft, blockRight BlockCoordinates
		switch direction {
		case North, South:
			blockLeft = BlockCoordinates{X: block.X - 1, Y: block.Y, Z: block.Z}
			blockRight = BlockCoordinates{X: block.X + 1, Y: block.Y, Z: block.Z}
		case East, West:
			blockLeft = BlockCoordinates{X: block.X, Y: block.Y, Z: block.Z - 1}
			blockRight = BlockCoordinates{X: block.X, Y: block.Y, Z: block.Z + 1}
		}
		blockbeneath := BlockCoordinates{X: block.X, Y: block.Y - 1, Z: block.Z}
		rcon.SetBlock(blockbeneath, "minecraft:glass")
		rcon.SetBlock(blockLeft, "minecraft:glass")
		rcon.SetBlock(blockRight, "minecraft:glass")
		blockRightOneUp := BlockCoordinates{X: blockRight.X, Y: blockRight.Y + 1, Z: blockRight.Z}
		rcon.SetBlock(blockRightOneUp, "minecraft:glass")
		blockLeftOneUp := BlockCoordinates{X: blockLeft.X, Y: blockLeft.Y + 1, Z: blockLeft.Z}
		rcon.SetBlock(blockLeftOneUp, "minecraft:glass")
		twoBlocksAbove := BlockCoordinates{X: block.X, Y: block.Y + 2, Z: block.Z}
		rcon.SetBlock(twoBlocksAbove, "minecraft:glass")
		if blockRangeIndex%5 == 0 {
			rcon.SetBlock(blockLeft, "minecraft:torch")
			rcon.SetBlock(block, "minecraft:powered_rail[shape="+direction.String()+"]")
			rcon.SetBlock(blockRight, "minecraft:redstone_torch")
			blockRangeIndex++
			continue
		}
		blockAbove := BlockCoordinates{X: block.X, Y: block.Y + 1, Z: block.Z}
		rcon.SetBlock(blockAbove, "minecraft:air")
		rcon.SetBlock(block, "minecraft:air")
		rcon.SetBlock(block, "minecraft:rail[shape="+direction.String()+"]")
		blockRangeIndex++
	}
}

func buildRail(start BlockCoordinates, end BlockCoordinates) []BlockCoordinates {
	var railPath []BlockCoordinates

	// Calculate the direction vector
	dx := end.X - start.X
	dy := end.Y - start.Y
	dz := end.Z - start.Z

	// Determine the number of steps needed (longest axis determines the steps)
	steps := int(math.Max(math.Abs(float64(dx)), math.Max(math.Abs(float64(dy)), math.Abs(float64(dz)))))

	// Calculate the step increments for each axis
	stepX := float64(dx) / float64(steps)
	stepY := float64(dy) / float64(steps)
	stepZ := float64(dz) / float64(steps)

	// Generate the rail path
	for i := 0; i <= steps; i++ {
		x := start.X + int(math.Round(float64(i)*stepX))
		y := start.Y + int(math.Round(float64(i)*stepY))
		z := start.Z + int(math.Round(float64(i)*stepZ))
		railPath = append(railPath, BlockCoordinates{X: x, Y: y, Z: z})
	}

	return railPath
}
