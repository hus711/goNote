package main

import "fmt"

/*
delete() 函数
delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key。
*/
func main9() {
	/* 创建 map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}

	fmt.Println("原始 map")

	/* 打印 map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* 删除元素 */
	delete(countryCapitalMap, "France")
	fmt.Println("Entry for France is deleted")

	fmt.Println("删除元素后 map")

	/* 打印 map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
}

/*
原始 map
   Capital of France is Paris
   Capital of Italy is Rome
   Capital of Japan is Tokyo
   Capital of India is New Delhi
   Entry for France is deleted
   删除元素后 map
   Capital of India is New Delhi
   Capital of Italy is Rome
   Capital of Japan is Tokyo
*/
