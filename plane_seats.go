package main

// plane seat cal
// you can also use imports, for example:
import (
	"fmt"
	"strings"
)

type seat struct {
	name       string
	isOccupied bool
}

func Solution(N int, S string) int {
	// fmt.Println(N, ">>> ", S)
	ps := constructPlaneSeats(N, S)

	maxFamily := 0
	for i, r := range ps {
		if r.G1.isFree() && r.G2.isFree() {
			maxFamily += 1
		}
		isG2G3Reserved := false
		if r.G2.isFree() && r.G3.isFree() {
			maxFamily += 1
			isG2G3Reserved = true
		}
		if r.G3.isFree() && r.G4.isFree() && !isG2G3Reserved {
			maxFamily += 1
		}
	}

	return maxFamily
}

func constructPlaneSeats(N int, S string) []*planeSeat {
	var ps = make([]*planeSeat, 0)

	os := strings.Split(S, " ")
	osMap := make(map[string]bool)
	for _, s := range os {
		osMap[s] = true
	}
	// 	fmt.Println(N, ">>> ", osMap)

	for i := 1; i <= N; i++ {
		s := &planeSeat{
			G1: &g1{
				aReserved: isReserved(i, "A", osMap),
				bReserved: isReserved(i, "B", osMap),
				cReserved: isReserved(i, "C", osMap),
			},
			G2: &g2{
				dReserved: isReserved(i, "D", osMap),
				eReserved: isReserved(i, "E", osMap),
			},
			G3: &g3{
				fReserved: isReserved(i, "F", osMap),
				gReserved: isReserved(i, "G", osMap),
			},
			G4: &g4{
				hReserved: isReserved(i, "H", osMap),
				jReserved: isReserved(i, "J", osMap),
				kReserved: isReserved(i, "K", osMap),
			},
		}

		ps = append(ps, s)
	}

	return ps
}

func isReserved(r int, seat string, osMap map[string]bool) bool {
	_, ok := osMap[fmt.Sprintf("%d%s", r, seat)]
	return ok
}

type planeSeat struct {
	G1 *g1
	G2 *g2
	G3 *g3
	G4 *g4
}

type g1 struct {
	aReserved bool
	bReserved bool
	cReserved bool
}

func (g *g1) isFree() bool {
	return !g.bReserved && !g.cReserved
}

type g2 struct {
	dReserved bool
	eReserved bool
}

func (g *g2) isFree() bool {
	return !g.dReserved && !g.eReserved
}

type g3 struct {
	fReserved bool
	gReserved bool
}

func (g *g3) isFree() bool {
	return !g.fReserved && !g.gReserved
}

type g4 struct {
	hReserved bool
	jReserved bool
	kReserved bool
}

func (g *g4) isFree() bool {
	return !g.hReserved && !g.jReserved
}
