package main

import (
	"fmt"
	"strconv"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

type seat struct {
	rowNum  int
	seatNum int
}

func main() {
	rowCount := getRowCount()
	seatsPerRow := getSeatCount()
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
		default:
			fmt.Println(Red + "\nPlease enter an integer representing one of the options presented." + Reset)
		}
	}
}

func getRowCount() (rowCount int) {
	for {
		rowCount = readInteger("Enter the number of rows:")
		if rowCount <= 0 {
			fmt.Println(Red + "\nTheater must have at least 1 row!" + Reset)
		} else if rowCount >= 100 {
			fmt.Println(Red + "\nTheater rows can't exceed 99!" + Reset)
		} else {
			fmt.Println()
			return rowCount
		}
	}
}

func getSeatCount() (rowCount int) {
	for {
		rowCount = readInteger("Enter the number of seats in each row:")
		if rowCount <= 0 {
			fmt.Println(Red + "\nRows must have at least 1 seat!" + Reset)
		} else if rowCount >= 100 {
			fmt.Println(Red + "\nSeats per row can't exceed 99!" + Reset)
		} else {
			fmt.Println()
			return rowCount
		}
	}
}

func printStatistics(seatMatrix [][]string) {
	purchasedCount := countPurchasedTickets(seatMatrix)
	purchasedTicketPercentage := (float64(purchasedCount) / float64(len(seatMatrix)*len(seatMatrix[0]))) * 100.0

	fmt.Printf(Yellow + "\nNumber of purchased tickets: " + Reset + "%d\n", purchasedCount)
	fmt.Printf(Yellow + "Percentage: " + Reset + "%.2f%%\n", purchasedTicketPercentage)
	fmt.Printf(Yellow + "Current income: "+ Reset + "$%d\n", getCurrentIncome(seatMatrix))
	fmt.Printf(Yellow + "Potential income for full theater: " + Reset + "$%d\n\n", getTotalIncome(seatMatrix))

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
	fmt.Print("\nCinema:\n   ")
	for i := 1; i <= len(seatMatrix[0]); i++ {
		if i < 10 {
			fmt.Printf(" %d ", i)
		} else {
			fmt.Printf(" %d", i)
		}
	}
	fmt.Println()

	for rowIndex := range seatMatrix {
		if rowIndex < 9 {fmt.Printf("%d  ", rowIndex+1)
		}else {fmt.Printf("%d ", rowIndex+1)}

		for _, seatStatus := range seatMatrix[rowIndex] {
			if seatStatus == "S" {
				fmt.Printf(Blue + " %s " + Reset, seatStatus)
			} else {
				fmt.Printf(Red + " %s " + Reset, seatStatus)
			}
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
		if ticket.rowNum <= 0 || ticket.rowNum > len(seatMatrix) {fmt.Println(Red + "\nInvalid row number!" + Reset)
		} else{ break }
	}

	fmt.Println()
	for {
		ticket.seatNum = readInteger("Enter a seat number in that row:")
		if ticket.seatNum <= 0 || ticket.seatNum > len(seatMatrix[0]) {fmt.Println(Red + "\nInvalid seat number!" + Reset)
		} else{ break }
	} 

	return ticket
}

func updateSeatMatrix(seatMatrix [][]string, ticket seat) (updatedMatrix [][]string) {
	updatedMatrix = seatMatrix
	updatedMatrix[ticket.rowNum-1][ticket.seatNum-1] = "B" // The minus ones account for zero-indexing

	return updatedMatrix
}

func readInteger(prompt string) (integer int) {
	var userInput string
	for {
		fmt.Println(prompt)
		_, _ = fmt.Scan(&userInput)
		integer, err := strconv.Atoi(userInput)
		if err == nil {
			return integer
		} else {
			fmt.Println(Red + "\nInput must be an integer!" + Reset)
		}
	}
	return integer
}
