package purchase
	const (
		car = "car"
    	truck = "truck"
    )

func NeedsLicense(kind string) bool {
    return kind == car || kind == truck
}

func ChooseVehicle(option1, option2 string) string {
    result := option1 < option2
    switch {
    case result:
        return option1 + " is clearly the better choice."
    default: 
        return option2 + " is clearly the better choice."
    } 
}

func CalculateResellPrice(originalPrice, age float64) float64 {
	switch  {
	    case age >= 3 && age < 10:
	    	return originalPrice * 0.7
        case age < 3:
        	return originalPrice * 0.8
        default:
        	return originalPrice * 0.5
	}
}
