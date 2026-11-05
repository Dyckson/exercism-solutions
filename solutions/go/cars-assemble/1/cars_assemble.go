package cars

const individual, group = 10000, 95000 

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    successRate = successRate / 100
	return float64(productionRate) * successRate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    productionRate = productionRate / 60
    successRate = successRate / 100
	result := float64(productionRate) * successRate
    
	return int(result)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	if carsCount >= 10 {
		individualCost := individual * (carsCount % 10)
		groupCost := group * (int(carsCount / 10))
		totalCost := individualCost + groupCost
		return uint(totalCost)
	}
	individualCost := individual * carsCount
	return uint(individualCost)
}