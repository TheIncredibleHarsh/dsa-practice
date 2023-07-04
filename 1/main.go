package main

import "fmt"

func main(){
	fmt.Println(CheckDuplicate([]int{1,2,3,1}))
}

func CheckDuplicate(nums []int) bool{
	hash := make(map[int]bool)
	for _, value := range nums {
		if (hash[value] == true){
			return true
		} else {
			hash[value] = true
		}

	}
	return false
}