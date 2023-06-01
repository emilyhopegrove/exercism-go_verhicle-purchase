package purchase

import "fmt"

// variables for CalculateResellPrice
const (
	lessThan3YearsPercentage     = 0.8
	between3And10YearsPercentage = 0.7
	tenYearsOrOlder              = 0.5
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	bestOption := option1
	if option2 < option1 {
		bestOption = option2
	}

	return fmt.Sprintf("%s is clearly the better choice.", bestOption)

}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age <= 0 {
		return -1
	}

	resellPrice := originalPrice

	if age < 3 {
		resellPrice *= lessThan3YearsPercentage
	} else if age >= 10 {
		resellPrice *= tenYearsOrOlder
	} else {
		resellPrice *= between3And10YearsPercentage
	}

	return resellPrice
}
