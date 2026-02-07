//Package weather provides tools 
//to find weather condition of a location.
package weather

var (
    //CurrentCondition is a string variable that stores information of the current weather condition.
	CurrentCondition string 
	//CurrentLocation is a string variable that stores the name of the location.
    CurrentLocation  string 
)

//Forecast function returns the current weather of a location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
