package main

import "fmt"

type Route struct {
	routes  []int
	routesM map[int]bool
}

type track struct {
	bus               []int
	cbusNo            int
	haveReachedTarget bool
}

var (
	// k: bus v:stops
	BusRoutes map[int]*Route
	// k:busStop v:availableBus
	BusOpt map[int][]int

	vistedBus, vistedBusStop map[int]bool
)

func init() {
	BusRoutes = make(map[int]*Route)
	BusOpt = make(map[int][]int)

	vistedBus = make(map[int]bool)
	vistedBusStop = make(map[int]bool)
}

func main() {
	var (
		source, target int
	)

	// INPUT: [1,2,7],[3,6,7]], source = 1, target = 6
	source = 1
	target = 6
	// [7,12],[4,5,15],[6],[15,19],[9,12,13]] source = 15, target = 12
	// source = 15
	// target = 12
	inputRoutes := [][]int{
		0: {1, 2, 7},
		1: {3, 6, 7},

		// 0: {7, 12},
		// 1: {4, 5, 15},
		// 2: {6},
		// 3: {15, 19},
		// 4: {9, 12, 13},

		// 0: {1, 2, 7}, // source = 1
		// 1: {3, 6, 7},
		// 2: {9, 3, 4}, // target = 4
	}
	// source = 1
	// target = 4
	transalateInput(inputRoutes)

	//get available bus
	busOpts, ok := BusOpt[source]
	if !ok {
		fmt.Println(-1)
		return
	}

	routeTrack := &track{
		bus:               make([]int, 0),
		haveReachedTarget: false,
	}

	for _, busNo := range busOpts {
		vistedBus[busNo] = true
		routeTrack.bus = append(routeTrack.bus, busNo)
		routeTrack.cbusNo = busNo

		// take stopes
		busStopes := BusRoutes[busNo]
		// if the target within the same bus route
		// ie. check if this bus goes to the target
		if ok := busStopes.routesM[target]; ok {
			routeTrack.haveReachedTarget = true
			break
		}

		findRoute(target, routeTrack, busStopes.routes)
		if routeTrack.haveReachedTarget {
			break
		}
	}

	if !routeTrack.haveReachedTarget {
		fmt.Println(-1)
		return
	}
	fmt.Println("No of Bus i took> ", len(routeTrack.bus))
}

// ok, check if i can move to other bus
// so, far we know, the same "bus" can't go to "target"
func findRoute(target int, routeTrack *track, busStopes []int) {
	for _, nstop := range busStopes {
		vistedBusStop[nstop] = true
		busOpts := BusOpt[nstop]
		fmt.Printf("Stop: %d, buses: %+v\n", nstop, busOpts)
		if len(busOpts) == 1 {
			continue
		}

		// from where i can take next bus
		for _, nbus := range busOpts {
			if nbus != routeTrack.cbusNo && !vistedBus[nbus] {
				fmt.Printf("INSIDE> Stop: %d, bus: %+v\n", nstop, nbus)
				busStopes := BusRoutes[nbus]
				if ok := busStopes.routesM[target]; ok {
					fmt.Printf("REACHED> Stop: %d, bus: %+v\n", nstop, nbus)
					routeTrack.bus = append(routeTrack.bus, nbus)
					routeTrack.cbusNo = nbus
					routeTrack.haveReachedTarget = true
					return
				}

				// from the next stop
				fmt.Println("Checking all the busstop from this bus> ", nbus)
				routeTrack.bus = append(routeTrack.bus, nbus)
				routeTrack.cbusNo = nbus
				findRoute(target, routeTrack, busStopes.routes)
			}
		}

		if routeTrack.haveReachedTarget {
			return
		}
	}
	return
}

func transalateInput(inputRoutes [][]int) {
	for bus, routes := range inputRoutes {
		routeM := make(map[int]bool)
		for _, busStop := range routes {
			routeM[busStop] = true
			bo, ok := BusOpt[busStop]
			if !ok {
				BusOpt[busStop] = append(BusOpt[busStop], bus)
				continue
			}
			bo = append(bo, bus)
			BusOpt[busStop] = bo
		}

		BusRoutes[bus] = &Route{
			routes:  routes,
			routesM: routeM,
		}
	}
}

// [[1,2,7],[3,6,7]], source = 1, target = 6
// routesInput := []int
// map[busStop] = []Routes

/**
//
{
	0 : [1,2,7]
	1: [3,6,7]
}

//
{
	1: [0] -> start
	2: [0]
	3: [1]

	6: [1]
	7: [0, 1] {
		//
		for bus (
			for busRoutes {

			}
		)
		->
	}
}

// 1 > 2

// -----
// 1 > 2 > N

*/

// map[routes] -> [bus] -> findclosest?
// linked list of routes! ||

// source -> stop := []multiple
// find[]

// recursiveSearch() {

// }
