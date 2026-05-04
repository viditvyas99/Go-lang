package variadic

import "fmt"


func main(){	

	sum := sumup(1,2,3,4)
	fmt.Println(sum)


}

func sumup(startingValue int, number ...int)int  {

	updatedNumber := append([]int{startingValue},number...)

	sum := 0 
	for _,val := range updatedNumber {
		sum += val
	}
	return sum
	
}
