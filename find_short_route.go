// two cities in file(csv), with distance
// take a city name,
// []A , B , Distance

package main

import (
	"fmt"
	"os"
)

type Graph struct {
	From     string
	TO       string
	Distance int
}

var GeoGraph []Graph

// check with map data
type GeoGraphMap map[string]map[string]int

func (g GeoGraphMap) set(city string, ip Graph) {
	data, ok := g[city]
	if !ok {
		g[city] = map[string]int{
			ip.TO: ip.Distance,
		}
		return
	}
	data[ip.TO] = ip.Distance
	g[city] = data
}

func main() {
	graphData := make(GeoGraphMap)
	for _, d := range GeoGraph {
		graphData.set(d.From, d)
	}

	// check is connected

	// {mumbia: 550, ahjvs:22, as}
	aruguments := os.Args
	city1, city2 := aruguments[1], aruguments[2]

	availablePaths, ok := graphData[city1]
	if !ok {
		panic("invalid city1")
	}

	// isDirectlyConnected := false
	if distance, ok := availablePaths[city2]; ok {
		fmt.Println(city1, " is connected directly to ", city2, " distance is: ", distance)
		return
	}
	// loop over to all the possible paths untill to reach city2, recursive one!
	// A ->
	// 1. paths, a, b, T
	// ------
	// 2. paths, a, b, c,... T
	isConnected := false
	connectedPath := make(map[string]int)

	for nextCity, distance := range availablePaths {
		availablePaths, _ := graphData[nextCity]
		// 1. check if target is connected directly
		if dis, ok := availablePaths[city2]; ok {
			isConnected = true
			connectedPath[nextCity] = distance
			connectedPath[city2] = dis
			fmt.Println(city1, " is connected to ", city2, "via ", nextCity, " distance is: ", distance+dis)
			return
		}
	}
	if isConnected {
		// calculate the distance...
		return
	}

	// 2. check from the all the available paths
	// TODO: to be checked/validate this logic

	for nextCity, _ := range availablePaths {
		nextCityAvailablePaths, _ := graphData[nextCity]

		// A
		for nextCity, distance := range nextCityAvailablePaths {

			// somthing here...

			// 1. check if target is connected directly
			if dis, ok := availablePaths[city2]; ok {
				isConnected = true
				connectedPath[nextCity] = distance
				connectedPath[city2] = dis
				fmt.Println(city1, " is connected directly to ", city2, " distance is: ", distance)
				return
			}
		}
	}
}

//
// func isconnected(nextCity string) {
// 	for {
// 		isConnected = false && !isLast {
// 			isconnected(nextCityA)
// 		}
// 		// wif
// 		if isConnected{

// 		}
// 	}
// }

func init() {
	GeoGraph = []Graph{
		{"mumbai", "delhi", 1421},
		{"mumbai", "bangalore", 982},
		{"pune", "bangalore", 840},
		{"pune", "guwahati", 2620},
		{"bangalore", "panaji", 595},
		{"bangalore", "jammu", 2761},
		{"srinagar", "jammu", 264},
		{"srinagar", "guwahati", 2747},
		{"kolkata", "guwahati", 1046},
		{"kolkata", "bangalore", 1881},
		{"panaji", "kochi", 779},
		{"kochi", "chennai", 692},
		{"chennai", "srinagar", 3020},
		{"srinagar", "leh", 418},
		{"kolkata", "leh", 2492},
		{"delhi", "chandigarh", 244},
		{"chandigarh", "leh", 724},
		{"panaji", "indore", 1026},
		{"indore", "bhopal", 1510},
		{"bhopal", "bangalore", 791},
	}
}
