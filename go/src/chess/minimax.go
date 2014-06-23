package chess

import (
	//"fmt"
	"math/rand"
)

//Minimax-Decision(state) returns an action
//v = Max-Value(state)
//return the action in successors(state) with value v
func miniMaxDecision(state ChessNode) (move string, stats string) {
	//fmt.Println("miniMaxDecision(state=", state, ")")
	//fmt.Println("pre utility = ", state.utility_value)
	v := minValue(state)
	//fmt.Println("v=", v)
	moves := successors(state)
	var equalMoves []Move
	for i := 0; i < len(moves); i++ {
		//fmt.Println("moves[", i, "]", moves[i])
		newBoard := makeMove(state.board, moves[i])
		u := utility(state.active_color, newBoard)
		//fmt.Println("u=", u)
		if v == u {
			//fmt.Println("FOUND, move =", moves[i])
			equalMoves = append(equalMoves, moves[i])
		}
	}
	if equalMoves != nil {
		randMove := equalMoves[rand.Intn(len(equalMoves))]
		move = formatNextMove(randMove)
	} else {
		//somehow couldn't find a move that matched max, so just grab one
		randMove := moves[rand.Intn(len(moves))]
		move = formatNextMove(randMove)
	}
	stats = formatStats()
	return move, stats
}

//Max-Value(state) returns a utility value
//If Terminal-Test(state) then return Utility(state)
//v <= -infinity
//for a, s in Successors(state) do
//  v <= Max(v, Min-Value(s))
//return v
func maxValue(state ChessNode) int {
	updateStats(state, 1)
	if terminalTest(state) {
		return state.utility_value
	}

	v := -9999999
	moves := successors(state)
	for i := 0; i < len(moves); i++ {
		s := minValue(nextState(state, moves[i]))
		if s >= v {
			//fmt.Println("\n\n----->found larger value, ", s, ",", moves[i], "\n\n")
			v = s
		}
	}
	return v
}

//Min-Value(state) returns a utility value
//If Terminal-Test(state) then return Utility(state)
//v <= +infinity
//for a, s in Successors(state) do
//  v <= Min(v, Max-Value(s))
//return v
func minValue(state ChessNode) int {
	updateStats(state, 1)
	//fmt.Println("minValue, state=", state)
	if terminalTest(state) {
		return state.utility_value
	}

	v := 9999999
	moves := successors(state)
	for i := 0; i < len(moves); i++ {
		s := maxValue(nextState(state, moves[i]))
		if s <= v {
			//fmt.Println("---->found smaller value, ", s, ",", moves[i])
			v = s
		}
	}
	return v
}

func terminalTest(state ChessNode) bool {
	//TODO check for checkmate
	//stop at a certain depth, then return the utility
	if state.depth >= 2 {
		return true
	}
	return false
}

func successors(state ChessNode) []Move {
	var moves []Move
	moves = getLegalMoves(state.active_color, state.board)
	return moves
}

func utility(color string, board [8][8]string) int {
	//return calculatePointValue(color, board)// -
	return calculatePointValue(opposite(color), board)
}

func nextState(state ChessNode, move Move) ChessNode {
	newBoard := makeMove(state.board, move)
	color := opposite(state.active_color)
	return ChessNode{
		board:         newBoard,
		active_color:  color,
		depth:         state.depth + 1,
		prev_move:     move,
		utility_value: utility(color, newBoard),
	}
}
