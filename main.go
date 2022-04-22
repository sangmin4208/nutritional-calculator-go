package main

import "fmt"

func main() {
	fmt.Println("Hello")
	ns := GetNutritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(300),
		Sugars:              SugarGram(1),
		SaturatedFattyAcids: SaturatedFattyAcids(3),
		Sodium:              SodiumMilligram(52),
		Fruits:              FruitsPercent(30),
		Fiber:               FiberGram(4),
		Protein:             ProteinGram(32),
	}, Food)

	fmt.Printf("Nutritional Score: %d\n", ns.Value)
	fmt.Printf("NutritScore: %v \n", ns.GetNutriScore())

}
