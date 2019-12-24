package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	pointsArray := readPoints("./turer.txt")

	for i, arr := range pointsArray {
		p, pErr := plot.New()
		if pErr != nil {
			fmt.Println("An error creating a new plot")
		}

		p.Title.Text = "KnowIT julekalender 24. desember"
		p.X.Label.Text = "X"
		p.Y.Label.Text = "Y"
		err := plotutil.AddLinePoints(p, fmt.Sprintf("Skitur %d", i), arr)
		if err != nil {
			fmt.Println("Error reading points")
		}
		if err := p.Save(4*vg.Inch, 4*vg.Inch, fmt.Sprintf("./Skitur/Skitur_%d.png", i)); err != nil {
			fmt.Println("Saveing grapth plot error")
		}
	}
}

func readPoints(path string) []plotter.XYs {
	var retArray []plotter.XYs
	file, fileError := os.Open(path)
	if fileError != nil {
		fmt.Println("Error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var xs []int
	var ys []int
	for scanner.Scan() {
		str := scanner.Text()
		if str == "---" {
			pts := make(plotter.XYs, len(xs))
			for i := range pts {
				pts[i].X = float64(xs[i])
				pts[i].Y = float64(ys[i])
			}
			xs = make([]int, 0)
			ys = make([]int, 0)
			retArray = append(retArray, pts)
		} else {
			tmp := strings.Split(str, ",")
			x, xParseErr := strconv.Atoi(tmp[0])
			y, yParseErr := strconv.Atoi(tmp[1])
			if xParseErr != nil || yParseErr != nil {
				fmt.Println("String to int parse error")
			}
			xs = append(xs, x)
			ys = append(ys, y)
		}
	}
	return retArray
}
