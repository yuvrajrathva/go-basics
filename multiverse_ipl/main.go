package main

func main(){
	players := newPool()

	soldout_players, remaining_players := sold(players, 2)

	soldout_players.print()
	remaining_players.print()
}
