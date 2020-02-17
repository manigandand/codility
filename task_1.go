/*
No of paths without temples
// Sample code to perform I/O:

fmt.Scanf("%s", &myname)            // Reading input from STDIN
fmt.Println("Hello", myname)        // Writing output to STDOUT
*/

// Write your code here
/*
6 3
1 2
1 6
2 3
2 4
2 5
2 3 4
*/
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type City struct {
	ID        int
	Roads     []int
	HasTemple bool
}

var totalCities, totalTemples int

var mu sync.RWMutex

var CitiesMap = make(map[int]*City)

func main() {
	var (
		inputs []string
		count  int
	)
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _ := reader.ReadString('\n')
		if input == "" {
			break
		}
		inputs = append(inputs, input)
	}

	if err := populateData(inputs); err != nil {
		fmt.Println("cities err: ", err)
		return
	}

	// fmt.Println("MAP: ", CitiesMap)
	// for _, city := range CitiesMap {
	//     fmt.Println(city.ID, city.Roads, city.HasTemple)
	// }

	city, ok := CitiesMap[1]
	if !ok {
		fmt.Println("Oops!")
		return
	}

	for _, road := range city.Roads {
		c := findCityCountHasNoTemple(road)
		count += c
	}

	fmt.Println(count)

}

func findCityCountHasNoTemple(cityID int) int {
	city, ok := CitiesMap[cityID]
	if !ok {
		fmt.Println("Oops!")
		return 0
	}
	if city.HasTemple {
		return 0
	}

	count := 1
	// NOTE: circular deps
	for _, road := range city.Roads {
		c := findCityCountHasNoTemple(road)
		count += c
	}

	return count
}

func populateData(inputs []string) error {
	last := len(inputs) - 1
	for i, s := range inputs {
		if i == 0 {
			if err := poulateCities(s); err != nil {
				return err
			}

			continue
		}

		if i == last {
			if err := updateTempaltes(s); err != nil {
				return err
			}
			continue
		}

		if err := updateRoads(s); err != nil {
			return err
		}
	}

	return nil
}

func poulateCities(s string) (err error) {
	inputs := strings.Split(s, " ")
	totalCities, err = strconv.Atoi(inputs[0])
	if err != nil {
		return err
	}

	totalTemples, err = strconv.Atoi(strings.Replace(inputs[1], "\n", "", -1))
	if err != nil {
		return err
	}

	for i := 1; i <= totalCities; i++ {
		city := &City{
			ID:        i,
			Roads:     make([]int, 0),
			HasTemple: false,
		}

		CitiesMap[i] = city
	}

	return nil
}

func updateRoads(s string) error {
	inputs := strings.Split(s, " ")
	cityID, err := strconv.Atoi(inputs[0])
	if err != nil {
		return err
	}

	nextCityID, err := strconv.Atoi(strings.Replace(inputs[1], "\n", "", -1))
	if err != nil {
		return err
	}

	city, ok := CitiesMap[cityID]
	if !ok {
		return errors.New(fmt.Sprintf("city not found %d", cityID))
	}

	if _, ok := CitiesMap[nextCityID]; !ok {
		return errors.New(fmt.Sprintf("city not found %d", nextCityID))
	}

	city.Roads = append(city.Roads, nextCityID)

	return nil
}

func updateTempaltes(s string) error {
	inputs := strings.Split(s, " ")
	if len(inputs) != totalTemples {
		return errors.New("wrong temples")
	}

	for _, i := range inputs {
		cityID, err := strconv.Atoi(strings.Replace(i, "\n", "", -1))
		if err != nil {
			return err
		}

		city, ok := CitiesMap[cityID]
		if !ok {
			return errors.New(fmt.Sprintf("city not found %d", cityID))
		}

		city.HasTemple = true
	}

	return nil
}
