package main

import "fmt"

func main() {
	// For exercise 55:
	type engine struct {
		electric bool
	}

	type vehicle struct {
		engine
		maker string
		model string
		doors int
		color string
	}

	// Exercise 53:
	type person struct {
		firstname               string
		lastname                string
		favoriteIcecreamFlavors []string
	}

	p1 := person{
		firstname:               "Panos",
		lastname:                "Vasilopoulos",
		favoriteIcecreamFlavors: []string{"Vanilla", "Chocolate", "Oreo", "Buenno"},
	}

	p2 := person{
		firstname:               "Efraim",
		lastname:                "Vasilopoulos",
		favoriteIcecreamFlavors: []string{"Brownies", "Vanilla", "Oreo", "Buenno"},
	}

	fmt.Printf("%s's favorite flavors are: ", p1.firstname)
	for _, v := range p1.favoriteIcecreamFlavors {
		fmt.Printf("%s ", v)
	}
	fmt.Println()

	fmt.Printf("%s 's favorite flavors are: ", p2.firstname)
	for _, v := range p2.favoriteIcecreamFlavors {
		fmt.Printf("%s ", v)
	}
	fmt.Println()
	fmt.Printf("%T ", p1.favoriteIcecreamFlavors)
	fmt.Println()
	fmt.Println("---------")

	// Exercise 54:
	p1.favoriteIcecreamFlavors = append(p1.favoriteIcecreamFlavors, "Strawberry")

	personsMap := map[string]person{
		p1.firstname: p1,
		p2.firstname: p2,
	}

	for k, v := range personsMap {
		fmt.Println(k, v)
	}

	if _, ok := personsMap["Panos12"]; !ok {
		fmt.Print("Panos was not found.")
	} else {
		fmt.Print("Panos was found.")
	}
	fmt.Println()
	fmt.Println("---------")

	// Exercise 55:
	vehicle1 := vehicle{
		engine: engine{
			electric: true,
		},
		maker: "Alfa romeo",
		model: "147",
		doors: 2,
		color: "ciel",
	}
	fmt.Println(vehicle1)
	fmt.Println()
	fmt.Println("----------")

	// Exercise 56:
	anonymusStruct1 := struct {
		firstname           string
		favFruitsToPriceMap map[string]float64
		favDrinks           []string
	}{
		firstname: "Panos",
		favFruitsToPriceMap: map[string]float64{
			"Oranges":      2.55,
			"Apples":       3.1,
			"Strawberries": 4.86,
		},
		favDrinks: []string{"Coffee", "Beer", "Tea"},
	}

	for k, v := range anonymusStruct1.favFruitsToPriceMap {
		fmt.Printf("%s cost %.2f per kg.\n", k, v)
	}
	fmt.Println()

	fmt.Print("Panos's favorite drinks are: ")
	for _, v := range anonymusStruct1.favDrinks {
		fmt.Printf("%s ", v)
	}
}
