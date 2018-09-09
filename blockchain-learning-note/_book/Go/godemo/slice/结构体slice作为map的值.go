package main

import "fmt"

type player struct {
	name string
	age  int
}

func main() {
	allPlayers := make(map[string][]player)
	chinaPlayers := make([]player, 0)
	chinaPlayers = append(chinaPlayers, player{"uzi", 18})
	chinaPlayers = append(chinaPlayers, player{"clearlove7", 27})
	koreaPlayers := []player{{"faker", 20}, {"deft", 19}}
	allPlayers["China"] = chinaPlayers
	allPlayers["Korea"] = koreaPlayers
	for k, v := range allPlayers {
		fmt.Printf("%s的选手有：\n", k)
		for _, play := range v {
			fmt.Println(play.name, "年龄为", play.age)
		}
	}
}
