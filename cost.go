package main

import "fmt"

type Cost uint8

const (
	Zero Cost = iota
	One
	Two
	Three
	Four
	Five
)

func HeroByCost(heroList Heros) {
	var oneCostHero Heros
	var twoCostHero Heros
	var threeCostHero Heros
	var fourCostHero Heros
	var fiveCostHero Heros
	for _, hero := range heroList {
		switch hero.Cost {
		case One:
			oneCostHero = append(oneCostHero, hero)
		case Two:
			twoCostHero = append(twoCostHero, hero)
		case Three:
			threeCostHero = append(threeCostHero, hero)
		case Four:
			fourCostHero = append(fourCostHero, hero)
		case Five:
			fiveCostHero = append(fiveCostHero, hero)
		}
	}
	fmt.Println("1️⃣ 💰 :", len(oneCostHero), oneCostHero.Name())
	fmt.Println("2️⃣ 💰 :", len(twoCostHero), twoCostHero.Name())
	fmt.Println("3️⃣ 💰 :", len(threeCostHero), threeCostHero.Name())
	fmt.Println("4️⃣ 💰 :", len(fourCostHero), fourCostHero.Name())
	fmt.Println("5️⃣ 💰 :", len(fiveCostHero), fiveCostHero.Name())
}
