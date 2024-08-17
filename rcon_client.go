package main

import (
	"strconv"
	"sync"

	"github.com/gorcon/rcon"
)

type rconClient struct {
	mu   *sync.Mutex
	conn *rcon.Conn
}

func (c *rconClient) Close() {
	c.conn.Close()
}

func (c *rconClient) Execute(command string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.conn.Execute(command)
}

func (c *rconClient) Say(message string) (string, error) {
	return c.Execute("/say " + message)
}

func (c *rconClient) BanPlayer(player string) (string, error) {
	return c.Execute("/ban " + player)
}

func (c *rconClient) GrantAdvancement(target, advancement string) (string, error) {
	return c.Execute("/advancement grant " + target + " " + advancement)
}

func (c *rconClient) RevokeAdvancement(target, advancement string) (string, error) {
	return c.Execute("/advancement revoke " + target + " " + advancement)
}

func (c *rconClient) GetAttribute(target, attribute string) (string, error) {
	return c.Execute("/attribute " + target + " " + attribute + " get")
}

func (c *rconClient) SetBaseAttribute(target, attribute string, value float64) (string, error) {
	return c.Execute("/attribute " + target + " " + attribute + " base set " + strconv.FormatFloat(value, 'f', -1, 64))
}

func (c *rconClient) AddModifierAttribute(target, attribute string, modifierType string, value float64) (string, error) {
	return c.Execute("/attribute " + target + " " + attribute + " modifier add " + modifierType + " " + strconv.FormatFloat(value, 'f', -1, 64))
}

func (c *rconClient) RemoveModifierAttribute(target, attribute string, modifierType string) (string, error) {
	return c.Execute("/attribute " + target + " " + attribute + " modifier remove " + modifierType)
}

func (c *rconClient) ExecuteCommand(command string) (string, error) {
	return c.Execute("/execute run " + command)
}

func (c *rconClient) ExecuteIf(command string) (string, error) {
	return c.Execute("/execute if " + command)
}

func (c *rconClient) ExecuteUnless(command string) (string, error) {
	return c.Execute("/execute unless " + command)
}

func (c *rconClient) ExecuteAs(target, command string) (string, error) {
	return c.Execute("/execute as " + target + " run " + command)
}

func (c *rconClient) ExecuteAt(target, command string) (string, error) {
	return c.Execute("/execute at " + target + " run " + command)
}

func (c *rconClient) ExecuteStore(target, path, scale string, command string) (string, error) {
	return c.Execute("/execute store " + target + " " + path + " " + scale + " run " + command)
}

func (c *rconClient) ExecutePositioned(x, y, z float64, command string) (string, error) {
	return c.Execute("/execute positioned " + strconv.FormatFloat(x, 'f', -1, 64) + " " + strconv.FormatFloat(y, 'f', -1, 64) + " " + strconv.FormatFloat(z, 'f', -1, 64) + " run " + command)
}

func (c *rconClient) ExecuteRotated(yaw, pitch float64, command string) (string, error) {
	return c.Execute("/execute rotated " + strconv.FormatFloat(yaw, 'f', -1, 64) + " " + strconv.FormatFloat(pitch, 'f', -1, 64) + " run " + command)
}

func (c *rconClient) ExecuteFacing(facing string, command string) (string, error) {
	return c.Execute("/execute facing " + facing + " run " + command)
}

func (c *rconClient) ExecuteAlign(axis, alignMode string, command string) (string, error) {
	return c.Execute("/execute align " + axis + " " + alignMode + " run " + command)
}

func (c *rconClient) ExecuteAnchored(anchor string, command string) (string, error) {
	return c.Execute("/execute anchored " + anchor + " run " + command)
}

func (c *rconClient) ExecuteIn(dimensions string, command string) (string, error) {
	return c.Execute("/execute in " + dimensions + " run " + command)
}

func (c *rconClient) SummonEntity(entity string, x, y, z float64) (string, error) {
	return c.Execute("/summon " + entity + " " + strconv.FormatFloat(x, 'f', -1, 64) + " " + strconv.FormatFloat(y, 'f', -1, 64) + " " + strconv.FormatFloat(z, 'f', -1, 64))
}

func (c *rconClient) ExecuteOn(target, command string) (string, error) {
	return c.Execute("/execute on " + target + " run " + command)
}

func (c *rconClient) AddBossBar(name, style, color string) (string, error) {
	return c.Execute("/bossbar add " + name + " " + style + " " + color)
}

func (c *rconClient) RemoveBossBar(name string) (string, error) {
	return c.Execute("/bossbar remove " + name)
}

func (c *rconClient) ListBossBars() (string, error) {
	return c.Execute("/bossbar list")
}

func (c *rconClient) SetBossBar(name, property, value string) (string, error) {
	return c.Execute("/bossbar set " + name + " " + property + " " + value)
}

func (c *rconClient) GetBossBar(name, property string) (string, error) {
	return c.Execute("/bossbar get " + name + " " + property)
}

func (c *rconClient) ClearInventory(targets string) (string, error) {
	return c.Execute("/clear " + targets)
}

func (c *rconClient) CloneBlocks(begin, end string, replaceMode bool) (string, error) {
	if replaceMode {
		return c.Execute("/clone " + begin + " " + end + " replace")
	}
	return c.Execute("/clone " + begin + " " + end)
}

func (c *rconClient) DamageEntity(target string, amount float64, damageType string) (string, error) {
	return c.Execute("/damage " + target + " " + strconv.FormatFloat(amount, 'f', -1, 64) + " " + damageType)
}

func (c *rconClient) MergeData(target, source string) (string, error) {
	return c.Execute("/data merge " + target + " " + source)
}

func (c *rconClient) GetData(target string) (string, error) {
	return c.Execute("/data get " + target)
}

func (c *rconClient) RemoveData(target string) (string, error) {
	return c.Execute("/data remove " + target)
}

func (c *rconClient) ModifyData(target, path, dataType, value string) (string, error) {
	return c.Execute("/data modify " + target + " " + path + " " + dataType + " " + value)
}

func (c *rconClient) EnableDatapack(name string) (string, error) {
	return c.Execute("/datapack enable " + name)
}

func (c *rconClient) DisableDatapack(name string) (string, error) {
	return c.Execute("/datapack disable " + name)
}

func (c *rconClient) ListDatapacks() (string, error) {
	return c.Execute("/datapack list")
}

func (c *rconClient) StartDebug() (string, error) {
	return c.Execute("/debug start")
}

func (c *rconClient) StopDebug() (string, error) {
	return c.Execute("/debug stop")
}

func (c *rconClient) StartFunction(name string) (string, error) {
	return c.Execute("/debug start " + name)
}

func (c *rconClient) StopFunction(name string) (string, error) {
	return c.Execute("/debug stop " + name)
}

func (c *rconClient) SetDefaultGamemode(gamemode string) (string, error) {
	return c.Execute("/defaultgamemode " + gamemode)
}

func (c *rconClient) SetDifficulty(difficulty string) (string, error) {
	return c.Execute("/difficulty " + difficulty)
}

func (c *rconClient) ClearEffect(target string) (string, error) {
	return c.Execute("/effect clear " + target)
}

func (c *rconClient) GiveEffect(target, effect string, duration, amplifier int) (string, error) {
	return c.Execute("/effect give " + target + " " + effect + " " + strconv.Itoa(duration) + " " + strconv.Itoa(amplifier))
}

func (c *rconClient) PerformAction(action string) (string, error) {
	return c.Execute("/me " + action)
}

func (c *rconClient) EnchantItem(targets, enchantment string, level int) (string, error) {
	return c.Execute("/enchant " + targets + " " + enchantment + " " + strconv.Itoa(level))
}

func (c *rconClient) AddExperience(targets string, amount int) (string, error) {
	return c.Execute("/experience add " + targets + " " + strconv.Itoa(amount))
}

func (c *rconClient) SetExperience(targets string, amount int) (string, error) {
	return c.Execute("/experience set " + targets + " " + strconv.Itoa(amount))
}

func (c *rconClient) QueryExperience(targets string) (string, error) {
	return c.Execute("/experience query " + targets)
}

func (c *rconClient) FillArea(from, to, block string, mode string) (string, error) {
	return c.Execute("/fill " + from + " " + to + " " + block + " " + mode)
}

func (c *rconClient) FillBiome(from, to, biome string, mode string) (string, error) {
	return c.Execute("/fillbiome " + from + " " + to + " " + biome + " " + mode)
}

func (c *rconClient) AddForceLoad(chunkX, chunkZ int) (string, error) {
	return c.Execute("/forceload add " + strconv.Itoa(chunkX) + " " + strconv.Itoa(chunkZ))
}

func (c *rconClient) RemoveForceLoad(chunkX, chunkZ int) (string, error) {
	return c.Execute("/forceload remove " + strconv.Itoa(chunkX) + " " + strconv.Itoa(chunkZ))
}

func (c *rconClient) QueryForceLoad(chunkX, chunkZ int) (string, error) {
	return c.Execute("/forceload query " + strconv.Itoa(chunkX) + " " + strconv.Itoa(chunkZ))
}

func (c *rconClient) ExecuteFunction(name string, arguments string) (string, error) {
	return c.Execute("/function " + name + " " + arguments)
}

func (c *rconClient) SetGameMode(gamemode string, target string) (string, error) {
	return c.Execute("/gamemode " + gamemode + " " + target)
}

func (c *rconClient) SetGameRule(rule, value string) (string, error) {
	return c.Execute("/gamerule " + rule + " " + value)
}

func (c *rconClient) GiveItem(targets, item string, count int) (string, error) {
	return c.Execute("/give " + targets + " " + item + " " + strconv.Itoa(count))
}

func (c *rconClient) GetHelp(command string) (string, error) {
	helpCommand := "/help"
	if command != "" {
		helpCommand += " " + command
	}
	return c.Execute(helpCommand)
}

func (c *rconClient) ReplaceItem(targets, slot, item string, count int) (string, error) {
	return c.Execute("/item replace " + targets + " " + slot + " " + item + " " + strconv.Itoa(count))
}

func (c *rconClient) ModifyItem(targets, slot, modifier string, value int) (string, error) {
	return c.Execute("/item modify " + targets + " " + slot + " " + modifier + " " + strconv.Itoa(value))
}

func (c *rconClient) KickPlayer(targets, reason string) (string, error) {
	return c.Execute("/kick " + targets + " " + reason)
}

func (c *rconClient) KillPlayer(targets string) (string, error) {
	return c.Execute("/kill " + targets)
}

func (c *rconClient) ListPlayers(uuids bool) (string, error) {
	if uuids {
		return c.Execute("/list uuids")
	}
	return c.Execute("/list")
}

func (c *rconClient) LocateStructure(structure string) (string, error) {
	return c.Execute("/locate " + structure)
}

func (c *rconClient) LocateBiome(biome string) (string, error) {
	return c.Execute("/locate biome " + biome)
}

func (c *rconClient) LocatePOI(poi string) (string, error) {
	return c.Execute("/locate poi " + poi)
}

func (c *rconClient) ReplaceLoot(target, lootTable string) (string, error) {
	return c.Execute("/loot replace " + target + " " + lootTable)
}

func (c *rconClient) InsertLoot(target, lootTable string) (string, error) {
	return c.Execute("/loot insert " + target + " " + lootTable)
}

func (c *rconClient) GiveLoot(target, lootTable string) (string, error) {
	return c.Execute("/loot give " + target + " " + lootTable)
}

func (c *rconClient) SpawnLoot(target, lootTable string) (string, error) {
	return c.Execute("/loot spawn " + target + " " + lootTable)
}

func (c *rconClient) SendMessage(targets, message string) (string, error) {
	return c.Execute("/msg " + targets + " " + message)
}

func (c *rconClient) SendParticle(name, pos string) (string, error) {
	return c.Execute("/particle " + name + " " + pos)
}

func (c *rconClient) PlaceFeature(feature, pos string) (string, error) {
	return c.Execute("/place feature " + feature + " " + pos)
}

func (c *rconClient) PlaceJigsaw(jigsaw, pos string) (string, error) {
	return c.Execute("/place jigsaw " + jigsaw + " " + pos)
}

func (c *rconClient) PlaceStructure(structure, pos string) (string, error) {
	return c.Execute("/place structure " + structure + " " + pos)
}

func (c *rconClient) PlaceTemplate(template, pos string) (string, error) {
	return c.Execute("/place template " + template + " " + pos)
}

func (c *rconClient) PlaySound(sound, category string) (string, error) {
	return c.Execute("/playsound " + sound + " " + category)
}

func (c *rconClient) RandomValue(min, max int) (string, error) {
	return c.Execute("/random value " + strconv.Itoa(min) + " " + strconv.Itoa(max))
}

func (c *rconClient) RandomRoll(chance float64) (string, error) {
	return c.Execute("/random roll " + strconv.FormatFloat(chance, 'f', -1, 64))
}

func (c *rconClient) RandomReset() (string, error) {
	return c.Execute("/random reset")
}

func (c *rconClient) Reload() (string, error) {
	return c.Execute("/reload")
}

func (c *rconClient) GiveRecipe(targets, recipe string) (string, error) {
	return c.Execute("/recipe give " + targets + " " + recipe)
}

func (c *rconClient) TakeRecipe(targets, recipe string) (string, error) {
	return c.Execute("/recipe take " + targets + " " + recipe)
}

func (c *rconClient) ReturnValue(value string) (string, error) {
	return c.Execute("/return " + value)
}

func (c *rconClient) Fail() (string, error) {
	return c.Execute("/return fail")
}

func (c *rconClient) RunFunction(name string) (string, error) {
	return c.Execute("/return run " + name)
}

func (c *rconClient) RideEntity(target, mount string) (string, error) {
	return c.Execute("/ride " + target + " mount " + mount)
}

func (c *rconClient) DismountEntity(target string) (string, error) {
	return c.Execute("/ride " + target + " dismount")
}

func (c *rconClient) SayMessage(message string) (string, error) {
	return c.Execute("/say " + message)
}

func (c *rconClient) ScheduleFunction(name string) (string, error) {
	return c.Execute("/schedule function " + name)
}

func (c *rconClient) ClearSchedule() (string, error) {
	return c.Execute("/schedule clear")
}

func (c *rconClient) CreateObjective(name, criteria string) (string, error) {
	return c.Execute("/scoreboard objectives add " + name + " " + criteria)
}

func (c *rconClient) RemoveObjective(name string) (string, error) {
	return c.Execute("/scoreboard objectives remove " + name)
}

func (c *rconClient) ListObjectives() (string, error) {
	return c.Execute("/scoreboard objectives list")
}

func (c *rconClient) ListPlayersInObjective(objective string) (string, error) {
	return c.Execute("/scoreboard players list " + objective)
}

func (c *rconClient) SetSpawnPoint(targets string) (string, error) {
	return c.Execute("/spawnpoint " + targets)
}

func (c *rconClient) SetWorldSpawn(pos string) (string, error) {
	return c.Execute("/setworldspawn " + pos)
}

func (c *rconClient) Spectate(target string) (string, error) {
	return c.Execute("/spectate " + target)
}

func (c *rconClient) SpreadPlayers(center string, spreadDistance, maxRange float64, respectTeams bool) (string, error) {
	if respectTeams {
		return c.Execute("/spreadplayers " + center + " " + strconv.FormatFloat(spreadDistance, 'f', -1, 64) + " " + strconv.FormatFloat(maxRange, 'f', -1, 64) + " respectTeams")
	}
	return c.Execute("/spreadplayers " + center + " " + strconv.FormatFloat(spreadDistance, 'f', -1, 64) + " " + strconv.FormatFloat(maxRange, 'f', -1, 64))
}

func (c *rconClient) SetBlock(x, y, z int, block string) (string, error) {
	return c.Execute("/setblock " + strconv.Itoa(x) + " " + strconv.Itoa(y) + " " + strconv.Itoa(z) + " " + block)
}

func newRCONClient(hostname string, port int, password string) (*rconClient, error) {
	conn, err := rcon.Dial(hostname+":"+strconv.Itoa(port), password)

	if err != nil {
		return nil, err
	}
	return &rconClient{conn: conn, mu: &sync.Mutex{}}, nil
}
