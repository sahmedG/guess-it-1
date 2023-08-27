package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/montanaflynn/stats"
)

var numbers []float64
func main() {
    reader := bufio.NewReader(os.Stdin)
	fileScanner := bufio.NewScanner(reader)
	for fileScanner.Scan() {
		number, _ := strconv.Atoi(fileScanner.Text())
		predictRange(number)
	}
}

func predictRange(number int) {
	if number > 0 && number < 250{
		numbers = append(numbers, float64(number))
	}

	mean := Average(numbers)
	//variance,_ := stats.Median(numbers)
	stdDev, _ := stats.StandardDeviation(numbers)
	// lin_reg,_ := stats.LinearRegression(stats.Series{})
	multiplier := 1.0

	lowerLimit := mean - (stdDev*multiplier)
	upperLimit := mean + (stdDev*multiplier)
	fmt.Printf("%d %d", int(math.Round(lowerLimit)), int(math.Round(upperLimit)))
	fmt.Println()
}

// Calculate the average of an array of float64 numbers
func Average(array []float64) float64 {
	n := len(array)
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += array[i]
	}
	return (sum) / float64(n)
}
