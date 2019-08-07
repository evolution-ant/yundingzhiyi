package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ByPrice Heros

func (a ByPrice) Len() int { return len(a) }
func (a ByPrice) Less(i, j int) bool {
	return a[i].Cost < a[j].Cost
}

func (a ByPrice) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type Hero struct {
	Name            string          `json:"name,omitempty"`
	RaceNames       RaceNames       `json:"races,omitempty"`
	ProfessionNames ProfessionNames `json:"professions,omitempty"`
	Cost            Cost            `json:"cost,omitempty"`
	Damage          int             `json:"damage,omitempty"`
	Arange          Aranged         `json:"aranged,omitempty"`
	Level           []Level         `json:"level,omitempty"`
}

type Level struct {
	Star   int     `json:"star,omitempty"`
	Aspeed float32 `json:"aspeed,omitempty"`
	HP     int     `json:"hp,omitempty"`
	MP     int     `json:"mp,omitempty"`
	Armor  int     `json:"armor,omitempty"`
}

type Aranged uint8

const (
	AZero Aranged = iota
	AOne
	ATwo
	AThree
	AFour
	AFive
)

type Heros []Hero

func (heros Heros) Name() (name string) {
	for i, hero := range heros {
		if i == len(heros)-1 {
			name += hero.Name
		} else {
			name += hero.Name + ","
		}
	}
	return
}

// func main() {
// heroList := GetHeroList()
// fmt.Println("heroList:", heroList)
// sort.Sort(ByPrice(heroList))
// HeroByCost(heroList)
// HeroByProfession(heroList)
// HeroByRace(heroList)
// }
func GetHeroList() Heros {
	var heroList = Heros{}
	b, err := ioutil.ReadFile("hero.json")
	if err != nil {
		fmt.Println("read file err:", err)
		return nil
	}
	err = json.Unmarshal(b, &heroList)
	if err != nil {
		fmt.Println("json unmarshal err:", err)
		return nil
	}
	return heroList
}
