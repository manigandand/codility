package main

// schdular worker model
// task -> worker(1-N)

// Registry ...
var Registry = make(map[string]Worker)

func main() {
	limit := 10000
	for _, w := range Registry {
		w.Do(limit)
	}
}

// Register ...
func Register(topic string, w Worker) {
	Registry[topic] = w
}
