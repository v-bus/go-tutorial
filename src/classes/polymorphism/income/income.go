/*Package income depicts current situation
 Imaginary organisation has income from two kinds of projects viz. fixed billing and time and material.
 The net income of the organisation is calculated by the sum of the incomes from these projects.
 Assume that the currency is dollars and we will not deal with cents. It will be represented using int.
/**/
package income

import "fmt"

//Income interface contains two methods calculate() which calculates and returns the income from the source and source() which returns the name of the source.
type Income interface {
	calculate() int //calculates and returns the income from the source
	source() string //returns the name of the source
}

//FixedBilling project has two fields "projectName" which represents the name of the project and "biddedAmount" which is the amount that the organisation has bid for the project.
type FixedBilling struct {
	projectName  string
	biddedAmount int
}

//Advertisement new income stream
type Advertisement struct {
	adName     string
	cPC        int
	noOfClicks int
}

//NewAdvertisement new func of Advertisement struct
func NewAdvertisement(adName string, CPC, noOfClicks int) *Advertisement {
	return &Advertisement{adName, CPC, noOfClicks}
}

//NewFixedBilling new func of FixedBilling struct
func NewFixedBilling(projectName string, biddedAmount int) *FixedBilling {
	return &FixedBilling{projectName, biddedAmount}
}

//NewTimeAndMaterial new func of TimeAndMaterial struct
func NewTimeAndMaterial(projectName string, noOfHours, hourlyRate int) *TimeAndMaterial {
	return &TimeAndMaterial{projectName, noOfHours, hourlyRate}
}

//TimeAndMaterial represent projects of Time and Material type.
type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tam TimeAndMaterial) calculate() int {
	return tam.noOfHours * tam.hourlyRate
}

func (tam TimeAndMaterial) source() string {
	return tam.projectName
}

func (a Advertisement) calculate() int {
	return a.cPC * a.noOfClicks
}

func (a Advertisement) source() string {
	return a.adName
}

//CalculateNetIncome calculates and prints the total income
func CalculateNetIncome(ic []Income) {
	var netincome int
	for _, v := range ic {
		fmt.Printf("Income from %s is %d\n", v.source(), v.calculate())
		netincome += v.calculate()
	}
	fmt.Println("Total net income is ", netincome)
}
