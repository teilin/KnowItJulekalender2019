package main

import (
	"container/ring"
	"fmt"
	"log"
	"math/big"
)

const (
	maxNumberOfSteps int = 1000740
	Elfs             int = 5
)

const Clockwise = +1
const Clockdumb = -1

func NextElf(a [Elfs]int) *ring.Ring {
	r := ring.New(Elfs)
	for i := 0; i < Elfs; i++ {
		r.Value = a[i]
		r = r.Next()
	}
	return r
}

func extreme(r *ring.Ring, max bool) (*ring.Ring, error) {
	flag := false
	needle := r
	for cur := needle.Next(); cur != r; cur = cur.Next() {
		switch {
		case cur.Value.(int) == needle.Value.(int):
			flag = true
		case max && cur.Value.(int) > needle.Value.(int):
			fallthrough
		case !max && cur.Value.(int) < needle.Value.(int):
			flag = false
			needle = cur
		}
	}
	if flag {
		return needle, fmt.Errorf("more than one candidate")
	}
	return needle, nil
}

func min(r *ring.Ring) (*ring.Ring, error) {
	return extreme(r, false)
}

func max(r *ring.Ring) (*ring.Ring, error) {
	return extreme(r, true)
}

func main() {
	work := NextElf([Elfs]int{0, 0, 0, 0, 0})
	last := work.Prev()

	direction := Clockwise

	for i := 1; i <= maxNumberOfSteps; i++ {
		if i == 1 {
			continue
		}

		work.Value = work.Value.(int) + 1

		if big.NewInt(int64(i)).ProbablyPrime(0) {
			possiblyNext, err := min(work)
			if err == nil {
				work = possiblyNext
				continue
			}
		}

		if i%28 == 0 {
			direction *= -1
			work = work.Move(direction)
			continue
		}

		if i%2 == 0 {
			possiblyNext := work.Move(direction)
			elf, err := max(work)
			if err == nil && elf == possiblyNext {
				work = possiblyNext
				work = work.Move(direction)
				continue
			}
		}

		if i%7 == 0 {
			work = last
			continue
		}

		work = work.Move(direction)
	}

	minWork, err := min(work)
	if err != nil {
		log.Fatal(err)
	}
	maxWork, err := max(work)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(maxWork.Value.(int) - minWork.Value.(int))
}
