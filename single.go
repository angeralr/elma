package main

import "fmt"

/*
Чудные вхождения в массив
Дан не пустой массив A состоит из N целых чисел.
Массив содержит некоторое число элементов и каждый элемент может быть парой для другого элемента с таким же значением,
кроме одного элемента, который не содержит пары. Необходимо найти этот элемент.

К примеру:

  A[0] = 9  A[1] = 3  A[2] = 9
  A[3] = 3  A[4] = 9  A[5] = 7
  A[6] = 9

Элементы по индексам 0 и 2 имеют значение 9,
а элемент по индексу 5 имеет значение 7 и не имеет пары.

Необходимо написать функцию, которая вернет элемент без пары

func Solution(A []int) int

N четное число в диапазоне [1..1,000,000];
каждый элемент массива A целое число в диапазоне [1..1,000,000,000];
*/

func main() {
	nums := []int{9, 3, 9, 3, 7, 9}
	fmt.Println(Solution(nums))

}

func Solution(nums []int) []int {
	counter := map[int]int{}
	for _, item := range nums {
		counter[item]++
	}
	resp := make([]int, 0, len(counter))
	for key, value := range counter {
		if value == 1 {
			item := key
			resp = append(resp, item)
		}
	}
	return resp
}