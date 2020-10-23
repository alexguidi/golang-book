package main

import "fmt"

func main() {
	var x [6]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	x[5] = 0

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))

	total = 0

	//first element _ is related with the index (like the i in the previous for)
	for _, value := range x {
		total += value
	}

	fmt.Println(total / float64(len(x)))

	y := [5]float64{98, 93, 77, 82, 83}

	for _, value := range y {
		fmt.Println(value)
	}

	z := make([]float64, 5)
	z[0] = 11
	fmt.Println("Using make: ", z)

	w := make([]float64, 5, 10)
	fmt.Println("Using make with capacity extended", w, "capacity: ", cap(w), "lenght: ", len(w))

	//get from index 0 until 3 (excluding the 3) x[0:] x[:5] x[:] are also valid
	arr := x[0:3]
	fmt.Println(arr)

	copy(w, z)
	fmt.Println("Slice after copy ", w)

	fmt.Println("-------------Maps-------------")

	m := make(map[string]int)
	m["key"] = 10
	fmt.Println(m["key"])
	delete(m, "key")
	fmt.Println(m["key"])

	elements := make(map[string]string) //short way map[string] string{"H": "Hydrogen", "He": "Helium",}
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"

	if name, ok := elements["H"]; ok {
		fmt.Println(name, " was found")
	}

	if name, ok := elements["Un"]; ok {
		fmt.Println(name, " was found")
	}
}
