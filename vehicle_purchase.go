package purchase
import "fmt"


// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	// if (kind == "car" || kind == "truck"){
	// 	return true
	// }else{
	// 	return false
	// }
	
	return (kind == "car" || kind == "truck")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if (option1 < option2){
		return fmt.Sprintf("%s is clearly the better choice.", option1)
	}else{
		return fmt.Sprintf("%s is clearly the better choice.", option2)
	}
}

//optimized solution -- make it scalabe (defining constants so it's more flexible)
const (
	lessThan3YearsPercentage = 0.8
	between3And10YearsPercentage = 0.7
	olderThan10YearsPercentage = 0.5
)

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	//if the vehicle is less than 3 years old, it costs 80% of the original price
	//If it is at least 10 years old, it costs 50%
	//If the vehicle is at least 3 years old but less than 10 years, 
		//it costs 70% of the original price.

	// Original solution -- make it work
	// highPrice := float64(originalPrice * .8)
	// mediumPrice := float64(originalPrice * .7)
	// lowPrice := float64(originalPrice * .5)

	// if (age < 3) {
	// 	return highPrice
	// } else if (age >= 3 && age < 10){
	// 	return mediumPrice
	// } else if (age >= 10){
	// 	return lowPrice
	// }
	// return -1

	//optimized solution -- make it scalabe (remove the hardcoding so it's more flexible)
	
	if age < 3 {
		return originalPrice * lessThan3YearsPercentage
	} else if age >= 3 && age < 10 {
		return originalPrice * between3And10YearsPercentage
	} else if age >= 10 {
		return originalPrice * olderThan10YearsPercentage
	}
	return -1
}
