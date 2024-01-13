package main

import "fmt"

func main() {
	rowCount := getRowCount()
	seatsPerRow := getSeatsPerRow()
	totalProfit := calculateTotalProfit(rowCount, seatsPerRow)

	fmt.Println("Total income:")
	fmt.Printf("$%d\n", totalProfit)
	return
}

func calculateTotalProfit(rowCount int, seatsPerRow int) (totalProfit int) {
	totalSeats := rowCount * seatsPerRow

	if totalSeats <= 60 {
		totalProfit = 10 * totalSeats
	} else if rowCount%2 == 0 { // Even number of rows
		totalProfit = 9 * totalSeats
	} else { // Odd number of rows
		frontSeatsCount := (rowCount / 2) * seatsPerRow
		backSeatsCount := frontSeatsCount + seatsPerRow
		totalProfit = (10 * frontSeatsCount) + (8 * backSeatsCount)
	}

	return totalProfit
}

func getRowCount() (numRows int) {
	fmt.Println("Enter the number of rows:")
	fmt.Scan(&numRows)

	return numRows
}

func getSeatsPerRow() (seatsPerRow int) {
	fmt.Println("Enter the number of seats in each row:")
	fmt.Scan(&seatsPerRow)

	return seatsPerRow
}
