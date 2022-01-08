package main

import "fmt"

type person struct {
	name    string
	age     int
	weight  float32
	job     string
	hobbies []string
}

func superAddEven(numbers []int) int {
	total := 0
	for _, number := range numbers {
		if number%2 == 0 {
			total += number
		}
	}
	return total
}

func appendTen(numbers []int) []int {
	return append(numbers, 10)
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(superAddEven(input))

	mymap := map[string]string{"a": "b", "c": "d"}

	for key, value := range mymap {
		fmt.Printf("key is %s %s\n", key, value)
	}

	myinfo := person{
		name:    "Hojoon Son",
		age:     30,
		weight:  99.2,
		job:     "software engineer",
		hobbies: []string{"workout", "eating", "watching youtube"}}

	fmt.Println(myinfo.hobbies)
	fmt.Println(myinfo.job)

}
