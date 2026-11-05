// Package weather é responsável por prever a condição climática atual de várias cidades de Goblinocus.
package weather

var (
	// CurrentCondition representa a condição atual que a cidade se encontra.
	CurrentCondition string

	// CurrentLocation representa a atual cidade de Goblinocus.
	CurrentLocation string
)

// Forecast tem como função receber a cidade e a condição, salvar elas respectivamente em CurrentLocation e CurrentCondition e retornar uma concatenação da mensagem formatada.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
