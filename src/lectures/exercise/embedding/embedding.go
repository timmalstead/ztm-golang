//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import "fmt"

// okay, generics are the missing bit I was looking for
type DashboardUnit interface {
	int | float32
}

// can do as an interface or as a literal in the function
// can use a tilde to approximate it. so if I have a type aliased int or float 32, it will be accepted

func findAverage[num ~int | ~float32](arrOfNums []num) num {
	var total num
	for _, numberToAverage := range arrOfNums {
		total += numberToAverage
	}
	return total / num(len(arrOfNums))
}

type BandwidthUsage struct {
	amount []int
}

func (b *BandwidthUsage) findAverageBandwithUse() int {
	return findAverage([]int(b.amount))
}

type CpuTemp struct {
	temp []float32
}

func (c *CpuTemp) findAverageCpuTemp() float32 {
	return findAverage(c.temp)
}

type MemoryUsage struct {
	amount []int
}

func (b *MemoryUsage) findAverageMemoryUse() int {
	return findAverage(b.amount)
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func main() {
	bandwidth := BandwidthUsage{[]int{50000, 100000, 130000, 80000, 1000}}
	temp := CpuTemp{[]float32{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]int{800000, 800000, 810000, 820000, 800000}}

	var dash = Dashboard{BandwidthUsage: bandwidth, CpuTemp: temp, MemoryUsage: memory}

	fmt.Println(dash.findAverageBandwithUse())
	fmt.Println(dash.findAverageCpuTemp())
	fmt.Println(dash.findAverageMemoryUse())

	// looks like the go compiler won't even let me use .amount at the top level, as it is used in two different structs
	fmt.Println(dash.BandwidthUsage.amount)
	fmt.Println(dash.MemoryUsage.amount)

}
