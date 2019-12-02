package main

import "fmt"
import "math"
import "os"
import "bufio"
import "strconv"
import "log"


func calculate_fuel(mass float64) int {
        fuel := int(math.Floor(mass/3) - 2)
	if fuel < 0 {
		fuel = 0
	}
	return fuel
}
func main() {
	input, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	total_fuel := 0
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		mass, _ := strconv.ParseFloat(line, 8)
		fuel := calculate_fuel(mass)
	 	total_fuel = total_fuel + fuel
		for {
			//fuel := calculate_fuel(float64(fuel))
			if fuel < 1 {
				break
			}
			fuel = calculate_fuel(float64(fuel))	
			total_fuel = total_fuel + fuel
		}
		fmt.Println(total_fuel)
	}
}
