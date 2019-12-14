package main

import (
	"fmt"
)

//Maps are references

var NationsCapitals map[string]string = make(map[string]string)

func main() {
	NationsCapitals["Spain"] = "Madrid"
	fmt.Println(NationsCapitals["Spain"])
}
