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
	fmt.Println()
	seatMatrix := createSeatMatrix(rowCount, seatsPerRow)

	for {
		switch choice := readInteger("1. Show the seats\n2. Buy a ticket\n3. Statistics\n0. Exit"); choice {
		case 1:
			printCinema(seatMatrix)
		case 2:
			buyTicket(seatMatrix)
		case 3:
			printStatistics(seatMatrix)
		case 0:
			return
		}
	}
}

func printStatistics(seatMatrix [][]string) {
	purchasedCount := countPurchasedTickets(seatMatrix)
	purchasedTicketPercentage := (float64(purchasedCount) / float64(len(seatMatrix)*len(seatMatrix[0]))) * 100.0

	fmt.Printf("\nNumber of purchased tickets: %d\n", purchasedCount)
	fmt.Printf("Percentage: %.2f%%\n", purchasedTicketPercentage)
	fmt.Printf("Current income: $%d\n", getCurrentIncome(seatMatrix))
	fmt.Printf("Total income: $%d\n\n", getTotalIncome(seatMatrix))

	return
}

func getTotalIncome(seatMatrix [][]string) (totalIncome int) {
	rowCount := len(seatMatrix)
	seatsPerRow := len(seatMatrix[0])
	totalSeats := rowCount * len(seatMatrix[0])

	if totalSeats <= 60 {
		return 10 * totalSeats
	}
	if rowCount%2 == 0 { // Even number of rows
		return 9 * totalSeats
	}

	// Odd number of rows
	frontSeatsCount := (rowCount / 2) * seatsPerRow
	backSeatsCount := frontSeatsCount + seatsPerRow
	return (10 * frontSeatsCount) + (8 * backSeatsCount)
}

func getTicketPrice(seatMatrix [][]string, seatRow int) (ticketPrice int) {
	rowCount := len(seatMatrix)
	totalSeats := rowCount * len(seatMatrix[0])

	if (totalSeats <= 60) || (seatRow <= rowCount/2) {
		return 10
	}

	return 8
}

func getCurrentIncome(seatMatrix [][]string) (income int) {
	for rowIndex, row := range seatMatrix {
		for _, seat_ := range row {
			if seat_ == "B" {
				income += getTicketPrice(seatMatrix, rowIndex+1) // +1 to account for zero indexing
			}
		}
	}
	return income
}

func countPurchasedTickets(seatMatrix [][]string) (count int) {
	for _, row := range seatMatrix {
		for _, seat_ := range row {
			if seat_ == "B" {
				count++
			}
		}
	}

	return count
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
	fmt.Print("\nCinema:\n ")
	for i := 1; i <= len(seatMatrix[0]); i++ {
		fmt.Printf(" %d", i)
	}
	fmt.Println()

	for rowIndex := range seatMatrix {
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
	ticket := getTicket(seatMatrix)
	seatMatrix = updateSeatMatrix(seatMatrix, ticket)
	fmt.Printf("\nTicket price: $%d\n\n", getTicketPrice(seatMatrix, ticket.rowNum))

	return
}

func getTicket(seatMatrix [][]string) (ticket seat) {
	fmt.Println()
	for {
		ticket.rowNum = readInteger("Enter a row number:")
		ticket.seatNum = readInteger("Enter a seat number in that row:")

		if isTicketSelectionValid(seatMatrix, ticket) {
			return ticket
		}
	}
}

func isTicketSelectionValid(seatMatrix [][]string, ticket seat) (isValid bool) {
	isBadRow := (ticket.rowNum <= 0) || (ticket.rowNum > len(seatMatrix))
	isBadSeat := (ticket.seatNum <= 0) || (ticket.seatNum > len(seatMatrix[0]))
	if isBadRow || isBadSeat {
		fmt.Print("\nWrong input!\n\n")
		return false
	}

	isSeatTaken := seatMatrix[ticket.rowNum-1][ticket.seatNum-1] == "B" // The minus ones account for zero-indexing
	if isSeatTaken {
		fmt.Print("\nThat ticket has already been purchased!\n\n")
		return false
	}

	return true
}

func updateSeatMatrix(seatMatrix [][]string, ticket seat) (updatedMatrix [][]string) {
	updatedMatrix = seatMatrix
	updatedMatrix[ticket.rowNum-1][ticket.seatNum-1] = "B" // The minus ones account for zero-indexing

	return updatedMatrix
}

func readInteger(prompt string) (integer int) {
	fmt.Println(prompt)
	_, _ = fmt.Scan(&integer)

	return integer
}
