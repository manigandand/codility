package main

import (
	"log"
	"os"
	"sync"
	"time"
)

const (
	Sleeping = "sleeping"
	Ready    = "ready"
	Checking = "checking"
	Cutting  = "cutting"
)

type BarberShop struct {
	name        string
	barber      *Barber
	roomSize    int
	waitingRoom chan *Customer
	wakeSig     chan int
	cuttingTime time.Duration
	mu          sync.Mutex
}

type Barber struct {
	name  string
	state string
	mu    sync.Mutex
}

type Customer struct {
	id int
	wg *sync.WaitGroup
}

func init() {
	log.SetPrefix("- ")
	log.SetOutput(os.Stdout)
	log.SetFlags(1)
}

func main() {
	waitingRoomSize := 5
	totalCutomer := 10
	barber := &Barber{
		name:  "Mani",
		state: Sleeping,
		mu:    sync.Mutex{},
	}
	shop := &BarberShop{
		name:        "Men's Fashion",
		barber:      barber,
		roomSize:    waitingRoomSize,
		waitingRoom: make(chan *Customer, waitingRoomSize),
		wakeSig:     make(chan int, 1),
		cuttingTime: 100 * time.Millisecond,
		mu:          sync.Mutex{}, // when working with mutiple barbers
	}
	// start the shop working
	go shop.Work()

	var wg sync.WaitGroup
	wg.Add(totalCutomer)
	for i := 1; i <= totalCutomer; i++ {
		go shop.Reception(&Customer{
			id: i,
			wg: &wg,
		})
	}
	wg.Wait()
	log.Println("done with all customers")
	time.Sleep(2 * time.Second)
}

// customers enters to the shop, checks the barber state.
// if barber sleeping, he wakes the barber and sits in the waiting room chair.
//
// if barber is busy and waiting room has available(empty) chair then he sits
// in the waiting room chair.
// else,
// he leaves the shop
func (s *BarberShop) Reception(cus *Customer) {
	time.Sleep((5 * time.Duration(cus.id)) * time.Millisecond)

	log.Printf("Customer [%d] arrived. [%s] is [%s]... WR: [%d]\n",
		cus.id, s.barber.name, s.barber.getState(), len(s.waitingRoom))
	switch s.barber.getState() {
	case Sleeping:
		// send wake signal
		select {
		case s.wakeSig <- cus.id:
		default:
			log.Println("cant wake barber")
		}
		// wait in the waiting room
	case Cutting:
		// wait in the waiting room
	case Checking:
		// wait in the waiting room
	}

	select {
	case s.waitingRoom <- cus:
		log.Printf("Customer [%d] is waiting... [%s] is [%s]... WR: [%d]\n",
			cus.id, s.barber.name, s.barber.getState(), len(s.waitingRoom))
	default:
		log.Printf("Customer [%d] is leaving :(. [%s] is [%s]... WR: [%d] full\n",
			cus.id, s.barber.name, s.barber.getState(), len(s.waitingRoom))
		cus.wg.Done()
	}
}

// barber shop with single barber starts working.
// he checks for the customers in the waiting room, if any customers available
// he goes to do haircut for that customer, else he goes to sleeping untill
// someone wakes up him.
func (s *BarberShop) Work() {
	log.Printf("[%s] started working with [%s]\n", s.name, s.barber.name)
	for {
		s.barber.setState(Checking)
		log.Printf("[%s] %s for customers...\n", s.barber.name, s.barber.state)

		select {
		case cus := <-s.waitingRoom:
			log.Printf("[%s] taking customer [%d]\n", s.barber.name, cus.id)
			// haircut
			s.barber.HairCut(cus, s.cuttingTime)
		default:
			// sleep untill someone wakes
			s.barber.setState(Sleeping)
			log.Printf("customers in the waiting room [%d]\n", len(s.waitingRoom))
			log.Printf("[%s] %s 😴Zzzzzzzzzzzzzzz...\n", s.barber.name, s.barber.state)
			c := <-s.wakeSig
			s.barber.setState(Ready)
			log.Printf("[%s] wokened by customer [%d]\n", s.barber.name, c)
		}
	}
}

func (b *Barber) HairCut(cus *Customer, t time.Duration) {
	b.mu.Lock()
	b.state = Cutting
	log.Printf("[%s] %s for customer [%d]\n", b.name, b.state, cus.id)
	time.Sleep(t)
	b.mu.Unlock()

	// finish customer
	log.Printf("[%s] %s finished customer [%d]\n", b.name, b.state, cus.id)
	cus.wg.Done()
}

func (b *Barber) getState() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.state
}

func (b *Barber) setState(state string) {
	b.mu.Lock()
	b.state = state
	b.mu.Unlock()
}
