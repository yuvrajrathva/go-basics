package main
import "fmt"

type squad []string;

func newSquad() squad{
	players := squad{}

	playerTypes := []string{"RH Batsman", "LH Batsman", "S Bowler", "F Bowler", "All rounder"}
	playerNames := []string{"KL Rahul", "Y Chahal", "R Jadeja"}

	for _, pType := range playerTypes{
		for _, pName := range playerNames{
			players = append(players, pName+" is a "+pType)
		}
	}

	return players
}

func (s squad) print(){
	for _, player := range s{
		fmt.Println(player)
	}
}