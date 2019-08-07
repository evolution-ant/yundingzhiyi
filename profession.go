package main

import "fmt"

type ProfessionNames []ProfessionName
type ProfessionName string

type Profession struct {
	Name   ProfessionName
	Fetter []Fetter
}

const (
	AssassinName   ProfessionName = "刺客"
	SwordsmenName  ProfessionName = "剑士"
	WarriorName    ProfessionName = "斗士"
	ElementistName ProfessionName = "元素"
	GuardName      ProfessionName = "护卫"
	GunnerName     ProfessionName = "枪手"
	KnightName     ProfessionName = "骑士"
	RangerName     ProfessionName = "游侠"
	ChangerName    ProfessionName = "换形"
	MageName       ProfessionName = "法师"
)

func HeroByProfession(heroList Heros) {
	var assassinHero Heros
	var swordsmenHero Heros
	var warriorHero Heros
	var elementistHero Heros
	var guardHero Heros
	var gunnerHero Heros
	var knightHero Heros
	var rangerHero Heros
	var changerHero Heros
	var mageHero Heros
	for _, hero := range heroList {
		if isMatchProfessionName(AssassinName, hero.ProfessionNames) {
			assassinHero = append(assassinHero, hero)
		}
		if isMatchProfessionName(SwordsmenName, hero.ProfessionNames) {
			swordsmenHero = append(swordsmenHero, hero)
		}
		if isMatchProfessionName(WarriorName, hero.ProfessionNames) {
			warriorHero = append(warriorHero, hero)
		}
		if isMatchProfessionName(ElementistName, hero.ProfessionNames) {
			elementistHero = append(elementistHero, hero)
		}
		if isMatchProfessionName(GuardName, hero.ProfessionNames) {
			guardHero = append(guardHero, hero)
		}
		if isMatchProfessionName(GunnerName, hero.ProfessionNames) {
			gunnerHero = append(gunnerHero, hero)
		}
		if isMatchProfessionName(KnightName, hero.ProfessionNames) {
			knightHero = append(knightHero, hero)
		}
		if isMatchProfessionName(RangerName, hero.ProfessionNames) {
			rangerHero = append(rangerHero, hero)
		}
		if isMatchProfessionName(ChangerName, hero.ProfessionNames) {
			changerHero = append(changerHero, hero)
		}
		if isMatchProfessionName(MageName, hero.ProfessionNames) {
			mageHero = append(mageHero, hero)
		}
	}
	fmt.Println("=================职业=====================")
	fmt.Println(AssassinName, ":", len(assassinHero), assassinHero.Name())
	fmt.Println(SwordsmenName, ":", len(swordsmenHero), swordsmenHero.Name())
	fmt.Println(WarriorName, ":", len(warriorHero), warriorHero.Name())
	fmt.Println(ElementistName, ":", len(elementistHero), elementistHero.Name())
	fmt.Println(GuardName, ":", len(guardHero), guardHero.Name())
	fmt.Println(GunnerName, ":", len(gunnerHero), gunnerHero.Name())
	fmt.Println(KnightName, ":", len(knightHero), knightHero.Name())
	fmt.Println(RangerName, ":", len(rangerHero), rangerHero.Name())
	fmt.Println(ChangerName, ":", len(changerHero), changerHero.Name())
	fmt.Println(MageName, ":", len(mageHero), mageHero.Name())
	fmt.Println("==========================================")
}
func isMatchProfessionName(professionName ProfessionName, professionNames ProfessionNames) bool {
	for _, p := range professionNames {
		if p == professionName {
			return true
		}
	}
	return false
}

func initProfession() {
	Assassin := Profession{}
	Assassin.Name = AssassinName
	Assassin.Fetter = []Fetter{{3, "125 demage"}, {6, "350 demage"}}

	Swordsmen := Profession{}
	Swordsmen.Name = SwordsmenName
	Swordsmen.Fetter = []Fetter{{4, "125 speed"}, {6, "350 speed"}}

	Warrior := Profession{}
	Warrior.Name = WarriorName
	Warrior.Fetter = []Fetter{{2, "125 HP"}, {4, "350 HP"}}

	Elementist := Profession{}
	Elementist.Name = ElementistName
	Elementist.Fetter = []Fetter{{3, "Elementist HP"}}

	Guard := Profession{}
	Guard.Name = GuardName
	Guard.Fetter = []Fetter{{2, "GuardName HP"}, {4, "GuardName HP"}}

	Gunner := Profession{}
	Gunner.Name = GunnerName
	Gunner.Fetter = []Fetter{{2, "Gunner HP"}, {4, "Gunner HP"}}

	Knight := Profession{}
	Knight.Name = KnightName
	Knight.Fetter = []Fetter{{2, "Knight HP"}, {4, "Knight HP"}}

	Ranger := Profession{}
	Ranger.Name = RangerName
	Ranger.Fetter = []Fetter{{2, "Ranger HP"}, {4, "Ranger HP"}}

	Changer := Profession{}
	Changer.Name = ChangerName
	Changer.Fetter = []Fetter{{3, "Changer HP"}, {6, "Changer HP"}}

	Mage := Profession{}
	Mage.Name = MageName
	Mage.Fetter = []Fetter{{3, "Mage HP"}, {6, "Mage HP"}}

}
