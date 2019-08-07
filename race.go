package main

import "fmt"

type RaceNames []RaceName

type RaceName string

type Race struct {
	Name   RaceName
	Fetter []Fetter
}

const (
	DevilName  RaceName = "恶魔"
	DragonName RaceName = "龙族"
	RoninName  RaceName = "浪人"
	IceName    RaceName = "极地"
	EmpireName RaceName = "帝国"
	NinjaName  RaceName = "忍者"
	RoyalName  RaceName = "贵族"
	ShadowName RaceName = "暗影"
	PirateName RaceName = "海盗"
	RobotName  RaceName = "机器"
	VoidName   RaceName = "虚空"
	WildName   RaceName = "狂野"
	JodlName   RaceName = "约德"
)

func HeroByRace(heroList Heros) {
	var DevilRace Heros
	var DragonRace Heros
	var RoninRace Heros
	var IceRace Heros
	var EmpireRace Heros
	var NinjaRace Heros
	var RoyalRace Heros
	var ShadowRace Heros
	var PirateRace Heros
	var RobotRace Heros
	var VoidRace Heros
	var WildRace Heros
	var JodlRace Heros
	for _, hero := range heroList {
		if isMatchRaceName(DevilName, hero.RaceNames) {
			DevilRace = append(DevilRace, hero)
		}
		if isMatchRaceName(DragonName, hero.RaceNames) {
			DragonRace = append(DragonRace, hero)
		}
		if isMatchRaceName(RoninName, hero.RaceNames) {
			RoninRace = append(RoninRace, hero)
		}
		if isMatchRaceName(IceName, hero.RaceNames) {
			IceRace = append(IceRace, hero)
		}
		if isMatchRaceName(EmpireName, hero.RaceNames) {
			EmpireRace = append(EmpireRace, hero)
		}
		if isMatchRaceName(NinjaName, hero.RaceNames) {
			NinjaRace = append(NinjaRace, hero)
		}
		if isMatchRaceName(RoyalName, hero.RaceNames) {
			RoyalRace = append(RoyalRace, hero)
		}
		if isMatchRaceName(ShadowName, hero.RaceNames) {
			ShadowRace = append(ShadowRace, hero)
		}
		if isMatchRaceName(PirateName, hero.RaceNames) {
			PirateRace = append(PirateRace, hero)
		}
		if isMatchRaceName(RobotName, hero.RaceNames) {
			RobotRace = append(RobotRace, hero)
		}
		if isMatchRaceName(VoidName, hero.RaceNames) {
			VoidRace = append(VoidRace, hero)
		}
		if isMatchRaceName(WildName, hero.RaceNames) {
			WildRace = append(WildRace, hero)
		}
		if isMatchRaceName(JodlName, hero.RaceNames) {
			JodlRace = append(JodlRace, hero)
		}
	}
	fmt.Println("=================种族=====================")
	fmt.Println(DevilName, ":", len(DevilRace), DevilRace.Name())
	fmt.Println(DragonName, ":", len(DragonRace), DragonRace.Name())
	fmt.Println(RoninName, ":", len(RoninRace), RoninRace.Name())
	fmt.Println(IceName, ":", len(IceRace), IceRace.Name())
	fmt.Println(EmpireName, ":", len(EmpireRace), EmpireRace.Name())
	fmt.Println(NinjaName, ":", len(NinjaRace), NinjaRace.Name())
	fmt.Println(RoyalName, ":", len(RoyalRace), RoyalRace.Name())
	fmt.Println(ShadowName, ":", len(ShadowRace), ShadowRace.Name())
	fmt.Println(PirateName, ":", len(PirateRace), PirateRace.Name())
	fmt.Println(RobotName, ":", len(RobotRace), RobotRace.Name())
	fmt.Println(VoidName, ":", len(VoidRace), VoidRace.Name())
	fmt.Println(WildName, ":", len(WildRace), WildRace.Name())
	fmt.Println(JodlName, ":", len(JodlRace), JodlRace.Name())
	fmt.Println("==========================================")
}

func isMatchRaceName(raceName RaceName, raceNames RaceNames) bool {
	for _, r := range raceNames {
		if r == raceName {
			return true
		}
	}
	return false
}

func initRaces() {

	Devil := Race{}
	Devil.Name = DevilName
	Devil.Fetter = []Fetter{{3, "125 demage"}, {6, "350 demage"}}

	Dragon := Race{}
	Dragon.Name = DragonName
	Dragon.Fetter = []Fetter{{3, "125 demage"}, {6, "350 demage"}}

	Ronin := Race{}
	Ronin.Name = RoninName
	Ronin.Fetter = []Fetter{{3, "125 demage"}, {6, "350 demage"}}

	Ice := Race{}
	Ice.Name = IceName
	Ice.Fetter = []Fetter{{3, "125 demage"}, {6, "350 demage"}}

	Empire := Race{}
	Empire.Name = EmpireName
	Empire.Fetter = []Fetter{{3, "125 demage"}, {6, "350 demage"}}

	Ninja := Race{}
	Ninja.Name = NinjaName
	Ninja.Fetter = []Fetter{{3, "125 demage"}, {6, "350 demage"}}

	Royal := Race{}
	Royal.Name = RoyalName
	Royal.Fetter = []Fetter{{3, "125 demage"}, {6, "350 demage"}}

	Shadow := Race{}
	Shadow.Name = ShadowName
	Shadow.Fetter = []Fetter{{3, "125 demage"}, {6, "350 demage"}}

	Pirate := Race{}
	Pirate.Name = PirateName
	Pirate.Fetter = []Fetter{{3, "125 demage"}, {6, "350 demage"}}

	Robot := Race{}
	Robot.Name = RobotName
	Robot.Fetter = []Fetter{{3, "125 demage"}, {6, "350 demage"}}

	Void := Race{}
	Void.Name = VoidName
	Void.Fetter = []Fetter{{3, "125 demage"}, {6, "350 demage"}}

	Wild := Race{}
	Wild.Name = WildName
	Wild.Fetter = []Fetter{{3, "125 demage"}, {6, "350 demage"}}

	Jodl := Race{}
	Jodl.Name = JodlName
	Jodl.Fetter = []Fetter{{3, "125 demage"}, {6, "350 demage"}}

}
