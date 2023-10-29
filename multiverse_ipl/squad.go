package main
import "fmt"

type pool []string;

func newPool() pool{
	players := pool{}

	playerTypes := []string{"RH Batsman", "LH Batsman", "S Bowler", "F Bowler", "All rounder"}
	playerNames := []string{"KL Rahul", "Y Chahal", "R Jadeja"}

	for _, pType := range playerTypes{
		for _, pName := range playerNames{
			players = append(players, pName+" is a "+pType)
		}
	}

	return players
}

func sold(p pool, number_of_players int) (pool, pool){
	return p[:number_of_players], p[number_of_players:]
}

func (p pool) print(){
	for i, player := range p{
		fmt.Println(i, player)
	}
}