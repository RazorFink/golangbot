package main

import "fmt"

func salReport(personSalary map[string]int) {
	for emp, sal := range personSalary {
		fmt.Printf("%s makes $%.2f\n", emp, float64(sal))
	}
}

func salaryEqual(salx map[string]int, saly map[string]int) bool {
	if len(salx) != len(saly) {
		fmt.Println("salary len ne")
		return false
	}

	for i, v := range salx {
		if saly[i] != v {
			return false
		}
	}
	return true
}

func main() {

	personSalary := map[string]int{
		"steve": 200,
		"henry": 1200,
		"carly": 2200,
		"alice": 5200,
		"carol": 6200,
	}

	personSalary2 := map[string]int{
		"steve": 200,
		"henry": 1200,
		"alice": 5200,
		"carol": 6200,
	}

	personSalary3 := map[string]int{
		"steve": 200,
		"henry": 1200,
		"carly": 2200,
		"alice": 5200,
		"carol": 6200,
	}

	fmt.Printf("personSalary == personSalary2: %t\n", salaryEqual(personSalary, personSalary2))
	fmt.Printf("personSalary == personSalary3: %t\n", salaryEqual(personSalary, personSalary3))
	fmt.Printf("personSalary2 == personSalary3: %t\n", salaryEqual(personSalary2, personSalary3))

	fmt.Println(personSalary)
	fmt.Println(personSalary["steve"])
	value, ok := personSalary["carol"]
	fmt.Printf("%t %d\n", ok, value)
	salReport(personSalary)

	delete(personSalary, "henry")
	fmt.Println("deleted henry")
	salReport(personSalary)

}
