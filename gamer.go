package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Man int

// 1L  劫8，慎8，凯南7.3，阿卡丽12，狮子狗7.3，寡妇7.3，派克8
// 7L  劫8，慎8，凯南7.3，阿卡丽12，狮子狗7.3，寡妇7.3，派克8
// 7L  劫8，慎8，凯南7.3，阿卡丽12，狮子狗7.3，寡妇7.3，派克8
// 7L  劫8，慎8，凯南7.3，阿卡丽12，狮子狗7.3，寡妇7.3，派克8

// 4L 1* 4.8 2* 7
// 5L 1* 6.5 2* 7
const (
	CardNumber = 5
)
const (
	zeroL Man = iota
	oneM
	twoM
	threeM
	fourM
	fiveM
	sixM
	sevenM
	eightM
	nineM
)

var onePool = make(map[string]Hero)
var twoPool = make(map[string]Hero)
var threePool = make(map[string]Hero)
var fourPool = make(map[string]Hero)
var fivePool = make(map[string]Hero)
var globalHerosPool = make(map[string]Hero)
var gamer *Gamer

type Gamer struct {
	Man      Man
	hasHeros Heros
}

func main() {
	initHerosPool()
	var level int
	for {
		fmt.Println("please enter level")
		fmt.Scanln(&level)
		if level > 1 && level < 10 {
			break
		}
		fmt.Println("level must in [2,10)")
	}
	fmt.Println("press d refresh")
	gamer := &Gamer{}
	for {
		var cmd string
		fmt.Scanln(&cmd)
		if cmd == "d" {
			gamer.Man = Man(level)
			Start(globalHerosPool, gamer)
			fmt.Println("pool left ", len(globalHerosPool))
		}
		if cmd == "s" {
			fmt.Println(gamer.hasHeros.Name())
		}
	}
}

func Start(herosPool map[string]Hero, gamer *Gamer) {
	rand.Seed(time.Now().UnixNano())
	chooseHero := Heros{}
	switch gamer.Man {
	case twoM:
		for {
			hero := ChooseFiveCards(onePool)
			chooseHero = append(chooseHero, hero)
			if len(chooseHero) == CardNumber {
				break
			}
		}
		fmt.Println(chooseHero.Name())
	case threeM:
		for {
			rn := rand.Intn(100) + 1
			if rn <= 5 {
				hero := ChooseFiveCards(threePool)
				chooseHero = append(chooseHero, hero)
			} else if 5 < rn && rn <= 65 {
				hero := ChooseFiveCards(onePool)
				chooseHero = append(chooseHero, hero)
			} else {
				hero := ChooseFiveCards(twoPool)
				chooseHero = append(chooseHero, hero)
			}
			if len(chooseHero) == CardNumber {
				break
			}
		}
		fmt.Println(chooseHero.Name())
	case fourM:
		// 50,35,15
		// 0-15 15-50 50-100
		for {
			rn := rand.Intn(100) + 1
			if rn <= 15 {
				hero := ChooseFiveCards(threePool)
				chooseHero = append(chooseHero, hero)
			} else if 15 < rn && rn <= 50 {
				hero := ChooseFiveCards(twoPool)
				chooseHero = append(chooseHero, hero)
			} else {
				hero := ChooseFiveCards(onePool)
				chooseHero = append(chooseHero, hero)
			}
			if len(chooseHero) == CardNumber {
				break
			}
		}
		fmt.Println(chooseHero.Name())
	case fiveM:
		// 3,25,35,37
		// 0-3 3-28 28-63 63-100
		for {
			rn := rand.Intn(1000) + 1
			if rn <= 3 {
				hero := ChooseFiveCards(fourPool)
				chooseHero = append(chooseHero, hero)
			} else if 3 < rn && rn <= 28 {
				hero := ChooseFiveCards(threePool)
				chooseHero = append(chooseHero, hero)
			} else if 28 < rn && rn <= 63 {
				hero := ChooseFiveCards(twoPool)
				chooseHero = append(chooseHero, hero)
			} else {
				hero := ChooseFiveCards(onePool)
				chooseHero = append(chooseHero, hero)
			}
			if len(chooseHero) == CardNumber {
				break
			}
		}
		fmt.Println(chooseHero.Name())
	case sixM:

		// 5,100,300,350,245
		// 0-5 5-105 105-405 405-755
		for {
			rn := rand.Intn(1000) + 1
			if rn <= 5 {
				hero := ChooseFiveCards(fivePool)
				chooseHero = append(chooseHero, hero)
			} else if 5 < rn && rn <= 105 {
				hero := ChooseFiveCards(fourPool)
				chooseHero = append(chooseHero, hero)
			} else if 105 < rn && rn <= 405 {
				hero := ChooseFiveCards(threePool)
				chooseHero = append(chooseHero, hero)
			} else if 405 < rn && rn <= 755 {
				hero := ChooseFiveCards(twoPool)
				chooseHero = append(chooseHero, hero)
			} else {
				hero := ChooseFiveCards(onePool)
				chooseHero = append(chooseHero, hero)
			}
			if len(chooseHero) == CardNumber {
				break
			}
		}
		fmt.Println(chooseHero.Name())
	case sevenM:
		// 2,15,33,30,20
		// 0-2 2-17 17-50 50-80 80-100
		for {
			rn := rand.Intn(100) + 1
			if rn <= 2 {
				hero := ChooseFiveCards(fivePool)
				chooseHero = append(chooseHero, hero)
			} else if 2 < rn && rn <= 17 {
				hero := ChooseFiveCards(fourPool)
				chooseHero = append(chooseHero, hero)
			} else if 17 < rn && rn <= 50 {
				hero := ChooseFiveCards(threePool)
				chooseHero = append(chooseHero, hero)
			} else if 50 < rn && rn <= 80 {
				hero := ChooseFiveCards(twoPool)
				chooseHero = append(chooseHero, hero)
			} else {
				hero := ChooseFiveCards(onePool)
				chooseHero = append(chooseHero, hero)
			}
			if len(chooseHero) == CardNumber {
				break
			}
		}
		fmt.Println(chooseHero.Name())
	case eightM:
		// 15,25,35,20,5
		// 0-15 15-40 40-75 75-95 95-100
		for {
			rn := rand.Intn(100) + 1
			if rn <= 15 {
				hero := ChooseFiveCards(fivePool)
				chooseHero = append(chooseHero, hero)
			} else if 15 < rn && rn <= 40 {
				hero := ChooseFiveCards(fourPool)
				chooseHero = append(chooseHero, hero)
			} else if 40 < rn && rn <= 75 {
				hero := ChooseFiveCards(threePool)
				chooseHero = append(chooseHero, hero)
			} else if 75 < rn && rn <= 95 {
				hero := ChooseFiveCards(twoPool)
				chooseHero = append(chooseHero, hero)
			} else {
				hero := ChooseFiveCards(onePool)
				chooseHero = append(chooseHero, hero)
			}
			if len(chooseHero) == CardNumber {
				break
			}
		}
		fmt.Println(chooseHero.Name())
	case nineM:
		// 10,15,35,30,10
		// 0-10 10-25 25-60 60-90 90-100
		for {
			rn := rand.Intn(100) + 1
			if rn <= 10 {
				hero := ChooseFiveCards(fivePool)
				chooseHero = append(chooseHero, hero)
			} else if 10 < rn && rn <= 25 {
				hero := ChooseFiveCards(fourPool)
				chooseHero = append(chooseHero, hero)
			} else if 25 < rn && rn <= 60 {
				hero := ChooseFiveCards(threePool)
				chooseHero = append(chooseHero, hero)
			} else if 60 < rn && rn <= 90 {
				hero := ChooseFiveCards(twoPool)
				chooseHero = append(chooseHero, hero)
			} else {
				hero := ChooseFiveCards(onePool)
				chooseHero = append(chooseHero, hero)
			}
			if len(chooseHero) == CardNumber {
				break
			}
		}
		fmt.Println(chooseHero.Name())
	}
	gamer.hasHeros = append(gamer.hasHeros, chooseHero...)
}
func ChooseFiveCards(heroMap map[string]Hero) (chooseHero Hero) {
	var keySlice []string
	for k, _ := range heroMap {
		keySlice = append(keySlice, k)
	}
	rand.Seed(time.Now().UnixNano())
	rn := rand.Intn(len(heroMap))
	chooseHero = heroMap[keySlice[rn]]
	delete(heroMap, keySlice[rn])
	delete(globalHerosPool, keySlice[rn])
	return chooseHero
}

func initHerosPool() {
	allHero := GetHeroList()
	HeroByCost(allHero)
	for _, hero := range allHero {
		switch hero.Cost {
		case One:
			for a := 0; a < 39; a++ {
				key := hero.Name + strconv.Itoa(a+1)
				onePool[key] = hero
			}
		case Two:
			for b := 0; b < 26; b++ {
				key := hero.Name + strconv.Itoa(b+1)
				twoPool[key] = hero
			}
		case Three:
			for c := 0; c < 21; c++ {
				key := hero.Name + strconv.Itoa(c+1)
				threePool[key] = hero
			}
		case Four:
			for d := 0; d < 13; d++ {
				key := hero.Name + strconv.Itoa(d+1)
				fourPool[key] = hero
			}
		case Five:
			for e := 0; e < 10; e++ {
				key := hero.Name + strconv.Itoa(e+1)
				fivePool[key] = hero
			}
		}
	}
	for k, v := range onePool {
		globalHerosPool[k] = v
	}
	for k, v := range twoPool {
		globalHerosPool[k] = v
	}
	for k, v := range threePool {
		globalHerosPool[k] = v
	}
	for k, v := range fourPool {
		globalHerosPool[k] = v
	}
	for k, v := range fivePool {
		globalHerosPool[k] = v
	}
}
func isExist(n int, slice []int) bool {
	for _, s := range slice {
		if s == n {
			return true
		}
	}
	return false
}
