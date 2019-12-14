package main

import "fmt"

const (
	maxSekvensLengde int64 = 217532235
)

var (
	alfabet       []int64 = make([]int64, 5)
	sekvens       []int64
	prevAdded     int64 = 0
	index         int64 = 0
	sumOfSevens   int64 = 0
	sekvensLengde int64 = 0
)

func init() {
	alfabet[0] = 2
	alfabet[1] = 3
	alfabet[2] = 5
	alfabet[3] = 7
	alfabet[4] = 11
}

func add(number int64, repeated int64) {
	if (sekvensLengde + repeated) > maxSekvensLengde {
		repeated = maxSekvensLengde - sekvensLengde
	}
	var i int64
	for i = 0; i < repeated; i++ {
		sekvens = append(sekvens, number)
	}
	sekvensLengde += repeated
	if number == 7 {
		sumOfSevens += (number * repeated)
	}
}

func alveSekvens(lengde int64) int64 {
	for {
		if index > 4 {
			index = 0
		}
		if sekvens == nil {
			add(alfabet[index], alfabet[index])
		} else {
			add(alfabet[index], sekvens[prevAdded])
		}
		prevAdded += 1
		index += 1
		if sekvensLengde >= lengde {
			break
		}
	}
	return sumOfSevens
}

func main() {
	fmt.Println(alveSekvens(maxSekvensLengde))
}
