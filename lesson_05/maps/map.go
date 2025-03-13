package main

import "fmt"

func main(){
	//working with map
	//Note: a map should be initialize before use to aviod runtime error
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x)

	//a function available to map
	delete(x,"key")
	fmt.Println(x)

	//making a periodic table
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"])
	
	//clean way of checking for an key and value that doesn't exist
	name,ok := elements["Ne"]
	fmt.Println(name,ok)

	//short way of creating a map
	cars := map[string]string{
		"Toyota":"4runner",
		"Lexus":"LX",
		"Land Curiseer":"Prado",
	}
	fmt.Println("Cars in my garrage: ",cars)

	//making a mp inside of a map
	person := map[string]map[string]string{
		"Person 1":{
			"name":"Walon",
			"age":"45",
		},
		"Person 2" :{
			"name":"Lamin",
			"age":"67",
		},
	}
	if el,ok := person["Person 3"]; ok {
		fmt.Println(el["name"],el["age"])
	}
}