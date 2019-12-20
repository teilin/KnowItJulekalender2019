package main

import "fmt"

const (
	maxNumberOfSteps int64 = 1000740
)

var (
	direction    int64           = 1
	currentIndex int64           = 0
	elves        map[int64]int64 = make(map[int64]int64)
)

func isEven(num int64) bool {
	if (num & 1) == 0 {
		return true
	} else {
		return false
	}
}

func getMin() (int64, int64) {
	var elv, numTasks int64
	for k, v := range elves {
		if elv == 0 && numTasks == 0 {
			elv = k
			numTasks = k
		}
		if v < numTasks {
			elv = k
			numTasks = v
		}
	}
	return elv, numTasks
}

func getMax() (int64, int64) {
	var elv, numTasks int64
	for k, v := range elves {
		if elv == 0 && numTasks == 0 {
			elv = k
			numTasks = k
		}
		if v > numTasks {
			elv = k
			numTasks = v
		}
	}
	return elv, numTasks
}

func numElves(numTasks int64) int64 {
	var count int64
	for _, v := range elves {
		if numTasks == v {
			count += 1
		}
	}
	return count
}

func getNextIndex(inc int64) int64 {
	next := currentIndex + (inc * direction)
	if next > 5 {
		return 5 - next
	}
	if next < 1 {
		if next == 0 {
			return 5
		} else {
			return 4
		}
	}
	return next
}

func getNext(step int64) {
	nextStep := step + 1
	elvMin, minCount := getMin()
	if isEven(nextStep) == false && numElves(minCount) == 1 {
		// Alven med færrest oppgaver
		currentIndex = elvMin
	} else if nextStep%28 == 0 {
		// Retningen snur
		direction *= -1
		currentIndex = getNextIndex(1)
	} else if isEven(nextStep) && numElves(minCount) == 1 {
		// Neste steg er partall og neste alv har gjort flest oppgaver, hopp over (så +/-2)
		currentIndex = getNextIndex(2)
	} else if nextStep%7 == 0 {
		// Alv 5 skal gjøre neste oppgave
		currentIndex = 5
	} else {
		// Gi gaven videre til neste alv i klokkeretingen
		currentIndex = getNextIndex(1)
	}
	elves[currentIndex]++
}

func main() {
	var i int64
	for i = 1; i <= maxNumberOfSteps; i++ {
		getNext(i)
	}
	_, minCount := getMin()
	_, maxCount := getMax()

	fmt.Println(minCount - maxCount)
	fmt.Println(maxCount - minCount)
}
