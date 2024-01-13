package main

import (
	"fmt"
)

type seat struct {
	rowNum  int
	seatNum int
}

func main() {
	rowCount := readInteger("Enter the number of rows:")
	seatsPerRow := readInteger("Enter the number of seats in each row:")
	seatMatrix := createSeatMatrix(rowCount, seatsPerRow)
	fmt.Println()

	for {
		switch userAction := readInteger("1. Show the seats\n2. Buy a ticket\n0. Exit"); userAction {
		case 1:
			printCinema(seatMatrix)
		case 2:
			buyTicket(seatMatrix)
		case 0:
			return
		}
	}
}

func printTicketPrice(seatMatrix [][]string, seatRow int) {
	rowCount := len(seatMatrix)
	totalSeats := rowCount * len(seatMatrix[0])

	var ticketPrice int
	if (totalSeats <= 60) || (seatRow <= rowCount/2) {
		ticketPrice = 10
	} else {
		ticketPrice = 8
	}

	fmt.Printf("\nTicket price: $%d\n\n", ticketPrice)
	return
}

func createSeatMatrix(rowCount, seatsPerRow int) (seatMatrix [][]string) {
	seatMatrix = make([][]string, rowCount)

	for i := 0; i < rowCount; i++ {
		seatMatrix[i] = make([]string, seatsPerRow)
		for j := 0; j < seatsPerRow; j++ {
			seatMatrix[i][j] = "S"
		}
	}

	return seatMatrix
}

func printCinema(seatMatrix [][]string) {
	fmt.Println("\nCinema:")
	fmt.Print(" ")
	for i := 1; i <= len(seatMatrix[0]); i++ {
		fmt.Printf(" %d", i)
	}
	fmt.Println()

	for rowIndex, _ := range seatMatrix {
		fmt.Printf("%d", rowIndex+1)
		for _, seatStatus := range seatMatrix[rowIndex] {
			fmt.Printf(" %s", seatStatus)
		}
		fmt.Println()
	}
	fmt.Println()

	return
}

func buyTicket(seatMatrix [][]string) {
	ticket := getSeatSelection()
	seatMatrix = updateSeatMatrix(seatMatrix, ticket)
	printTicketPrice(seatMatrix, ticket.rowNum)

	return
}

func getSeatSelection() (seatSelection seat) {
	seatSelection = seat{
		rowNum:  readInteger("Enter a row number:"),
		seatNum: readInteger("Enter a seat number in that row:"),
	}

	return seatSelection
}

func updateSeatMatrix(seatMatrix [][]string, seatSelection seat) (updatedMatrix [][]string) {
	updatedMatrix = seatMatrix
	updatedMatrix[seatSelection.rowNum-1][seatSelection.seatNum-1] = "B" // The minus ones account for zero-indexing

	return updatedMatrix
}

func readInteger(prompt string) (integer int) {
	fmt.Println(prompt)
	fmt.Scan(&integer)

	return integer
}
