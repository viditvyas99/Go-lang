package main

import "fmt"

type foaltMap map[string]float64

func (m foaltMap) outPrintMap() {
	fmt.Println(m)
}

func main() {

	nums := make([]float64, 2, 8)

	for i := 0; i <= len(nums) && i <= 11; i++ {
		nums = append(nums, float64(i))
		fmt.Println("len:", len(nums), "cap:", cap(nums))
	}

	testObject := make(foaltMap, 3)

	testObject["test"] = 7.0

	testObject["num"] = 9.9

	testObject["test1"] = 6.6

	testObject.outPrintMap()

}
