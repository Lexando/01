package main

import "fmt"

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var result []int
	var double bool //Буде показувати, чи знайдено дублікат.

	result = append(result, arr[0]) //Додаємо перший елемент в result безумовно.

	/*Ідея рішення - перебираємо елементи arr і порівнюємо кожен з них - з кожним
	  елементом зрізу result. Якщо співпадінь не знайдено - додаємо цей елемент до result,
	  і повторюємо знову. Таким чином, в result потраплять лише унікальні числа.*/

	for i := 1; i < len(arr); i++ { //Перебор arr.
		double = false                     //Скидаємо на false при новому заході у result.
		for j := 0; j < len(result); j++ { //Перебор result.
			if result[j] == arr[i] {
				double = true //Якщо знайшли дубль - запам'ятали і вийшли з циклу.
				break
			}
		}
		if !double { //Якщо немає дублів, то дописуємо новий елемент у колекцію result.
			result = append(result, arr[i])
		}
	}
	fmt.Println(result)
}
