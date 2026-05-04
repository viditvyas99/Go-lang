package anonmyous

import "fmt"


func main(){

numbers := make([]int, 0) 

numbers = append(numbers, 1,2,3)
double:=  createTansformer(2)
transedNumber := transNumber(&numbers,double)

fmt.Println(transedNumber)

}

func transNumber (numbers *[]int,transform func(int) int) []int {

	newNumbers := make([]int,0)

	for  _ , val :=range *numbers {
		newNumbers = append(newNumbers, transform(val))
	}

	return newNumbers
	}


func createTansformer(i int)  func(int) int {
	return func (j int)  int{
		return  j*i
	}
}
