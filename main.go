package main

import "fmt"

func main() {

	ns := GetNutritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(0),
		Sugar:               SugarGram(10),
		SaturatedFattyAcids: SaturatedFattyAcids(2),
		Sodium:              SodiumMilligram(500),
		Fruits:              FruitsPercent(60),
		Fibre:               FiberGram(4),
		Protien:             ProtienGram(2),
	}, Food)

	fmt.Printf("Your Nutritional Score:%d\n", ns)
	fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())
}
