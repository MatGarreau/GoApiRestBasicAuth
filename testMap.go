package main

import "fmt"


//func contains(slice []string, item string) bool {
//	set := make(map[string]struct{}, len(slice))
//	for _, s := range slice {
//		set[s] = struct{}{}
//	}
//
//	_, ok := set[item]
//	return ok
//}

// list of gpios available on the raspberry pi
var gpios = []int {0, 1, 2, 3, 4, 7, 8, 9, 10, 11, 17, 18, 22, 23, 24, 25}

func isAvailableGpio(gpio int) bool {
	gpioMap := make(map[int]struct{}, len(gpios))
	for _, pin := range gpios {
		gpioMap[pin] = struct{}{}
	}

	_, ok := gpioMap[gpio]
	return ok
}


func main() {

	//s := []string{"a", "b"}
	//int1 := 25
	fmt.Println(isAvailableGpio(0))
	fmt.Println(isAvailableGpio(12))
	fmt.Println(isAvailableGpio(5))
	fmt.Println(isAvailableGpio(26))
	fmt.Println(isAvailableGpio(25))
}
