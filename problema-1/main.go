package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// Valid FENs
	validFen1 := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	validFen2 := "r1bqkbnr/p1pp1ppp/1pn5/4p3/2B1P3/5N2/PPPP1PPP/RNBQK2R w KQkq - 2 4"

	// Invalid FENs
	invalidFen1 := "rnbqkbnr/pppppppp/7/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1" // Invalid rank sum
	invalidFen2 := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0"   // Not enough fields
	invalidFen3 := "rnbqkbnr/ppppppXp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1" // Invalid character

	fmt.Println("--- Valid FENs ---")
	fmt.Printf("'%s' -> %s\n", validFen1, resultToString(checkFen(validFen1)))
	fmt.Printf("'%s' -> %s\n", validFen2, resultToString(checkFen(validFen2)))

	fmt.Println("\n--- Invalid FENs ---")
	fmt.Printf("'%s' -> %s\n", invalidFen1, resultToString(checkFen(invalidFen1)))
	fmt.Printf("'%s' -> %s\n", invalidFen2, resultToString(checkFen(invalidFen2)))
	fmt.Printf("'%s' -> %s\n", invalidFen3, resultToString(checkFen(invalidFen3)))
}

func resultToString(b bool) string {
	if b {
		return "Valid"
	}
	return "Invalid"
}

func isValidPiece(c rune) bool {
	return strings.ContainsRune("prnbqkPRNBQK", c)
}

func checkPiecePlacement(piecePlacement string) bool {
	ranks := strings.Split(piecePlacement, "/")
	if len(ranks) != 8 {
		return false
	}

	for _, rank := range ranks {
		squareCount := 0
		for _, c := range rank {
			if unicode.IsDigit(c) {
				digit, _ := strconv.Atoi(string(c))
				if digit == 0 {
					return false
				}
				squareCount += digit
			} else if isValidPiece(c) {
				squareCount++
			} else {
				return false
			}
		}
		if squareCount != 8 {
			return false
		}
	}
	return true
}

func checkActiveColor(activeColor string) bool {
	return activeColor == "w" || activeColor == "b"
}

func checkCastling(castling string) bool {
	if castling == "-" {
		return true
	}
	if len(castling) > 4 {
		return false
	}
	k, q, K, Q := 0, 0, 0, 0
	for _, c := range castling {
		switch c {
		case 'K':
			K++
		case 'Q':
			Q++
		case 'k':
			k++
		case 'q':
			q++
		default:
			return false // Invalid char in castling string
		}
	}
	if k > 1 || q > 1 || K > 1 || Q > 1 {
		return false // Duplicate castling rights
	}
	return true
}

func checkEnPassant(enPassant string, activeColor string) bool {
	if enPassant == "-" {
		return true
	}
	if len(enPassant) != 2 {
		return false
	}
	file := enPassant[0]
	rank := enPassant[1]
	if file < 'a' || file > 'h' {
		return false
	}
	if !((rank == '3' && activeColor == "b") || (rank == '6' && activeColor == "w")) {
		return false
	}
	return true
}

func checkHalfmoveClock(halfmoveStr string) bool {
	halfmove, err := strconv.Atoi(halfmoveStr)
	if err != nil {
		return false
	}
	return halfmove >= 0
}

func checkFullmoveNumber(fullmoveStr string) bool {
	fullmove, err := strconv.Atoi(fullmoveStr)
	if err != nil {
		return false
	}
	return fullmove >= 1
}

func checkFen(fen string) bool {
	fields := strings.Fields(fen)
	if len(fields) != 6 {
		return false
	}

	activeColor := fields[1]

	return checkPiecePlacement(fields[0]) &&
		checkActiveColor(activeColor) &&
		checkCastling(fields[2]) &&
		checkEnPassant(fields[3], activeColor) &&
		checkHalfmoveClock(fields[4]) &&
		checkFullmoveNumber(fields[5])
}
