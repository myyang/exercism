/*
Package queenattack provides notification verification.

Chess notification is briefly described here:
Using A-H and 1-8 represents chess movement on chessboard for recording games.

For more detail, please visit: See http://en.wikipedia.org/wiki/Algebraic_notation_(chess)

*/
package queenattack

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

// QAtkError defines custom error
type QAtkError struct{}

// Error returns error msg string
func (e QAtkError) Error() string {
	return fmt.Sprintf("Queen attack error")
}

// CanQueenAttack check two input position and return a (attackable, valid) boolean pair
func CanQueenAttack(QueenPosition1, QueenPosition2 string) (bool, error) {

	row1, col1 := getPos(QueenPosition1)
	row2, col2 := getPos(QueenPosition2)

	switch {
	case row1 < 0 || row2 < 0 || col1 < 0 || col2 < 0:
		return false, QAtkError{}
	case row1 == row2 && col1 == col2:
		return false, QAtkError{}
	case row1 == row2 || col1 == col2:
		return true, nil
	case AbsInt(row1-row2) == AbsInt(col1-col2):
		return true, nil
	}
	return false, nil
}

func getPos(pos string) (row, col int) {
	posRe, err := regexp.Compile(`^(?P<row>[a-h]{1})(?P<col>[1-8]{1})$`)
	if err != nil {
		log.Fatal(`Regex fail`)
		return -1, -1
	}

	p := posRe.FindStringSubmatch(pos)
	if p == nil {
		return -1, -1
	}

	row, err = strconv.Atoi(fmt.Sprintf("%v", p[1][0]-"a"[0]))
	if err != nil {
		return -1, -1
	}

	col, err = strconv.Atoi(p[2])
	if err != nil {
		return -1, -1
	}

	return row, col
}

// AbsInt return absolute value int
func AbsInt(i int) int {
	if i < 0 {
		return i * (-1)
	}
	return i
}
