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
	printCinema(seatMatrix)

	seatSelection := getSeatSelection()
	seatMatrix = updateSeatMatrix(seatMatrix, seatSelection)
	printTicketPrice(rowCount, seatsPerRow, seatSelection.rowNum)

	printCinema(seatMatrix)

	return
}

func printTicketPrice(rowCount, seatsPerRow, seatRow int) {
	totalSeats := rowCount * seatsPerRow

	var ticketPrice int
	if (totalSeats <= 60) || (seatRow <= rowCount/2) {
		ticketPrice = 10
	} else {
		ticketPrice = 8
	}

	fmt.Printf("Ticket price: $%d\n", ticketPrice)
	fmt.Println()
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
	fmt.Println("Cinema:")
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

func getSeatSelection() (seatSelection seat) {
	seatSelection = seat{
		rowNum:  readInteger("Enter a row number:"),
		seatNum: readInteger("Enter a seat number in that row:"),
	}
	fmt.Println()

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
