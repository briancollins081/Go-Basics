/*
 * Polymorphism in Go is achieved with the help of interfaces.
 *
 * A variable of type interface can hold any value which implements the interface.
 * This property of interfaces is used to achieve polymorphism in Go.
 *
 */
package main

import "fmt"

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName  string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

type Advertisement struct {
	adName     string
	CPC        int
	noOfClicks int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

// Adding a new income stream to the above program
func (a Advertisement) calculate() int {
	return a.noOfClicks * a.CPC
}
func (a Advertisement) source() string {
	return a.adName
}

func calculateNetIncome(ic []Income) {
	var netincome int = 0
	for _, income := range ic {
		fmt.Printf("Income from %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net Income of the Organization = $%d\n", netincome)
}

func main() {
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 50000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 100000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 500, hourlyRate: 30}
	// Adding a new income stream to the above program
	bannerAd := Advertisement{adName: "Banner Ad", CPC: 2, noOfClicks: 500}
	popupAd := Advertisement{adName: "Popup Ad", CPC: 5, noOfClicks: 750}
	incomeStreams := []Income{project1, project2, project3, bannerAd, popupAd}
	calculateNetIncome(incomeStreams)
}
