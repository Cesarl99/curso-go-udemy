package main

import ("fmt"
"reflect")

func main(){
	fmt.Println("Arrays e Slice")

	var array1 [5]int
	array1[0]=1

	fmt.Println(array1)

	array2 := [5]string {"posição 2","posição 3","posição 4","posição 5","posição 6"}
	fmt.Println(array2)

	array3 := [...]int {3,5,10,25,36}
	fmt.Println(array3)

	slice1 := []int{2,6,14}
	fmt.Println(slice1)

	slice1 = append(slice1,18)

	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array1))
	fmt.Println(slice1)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "posição alterada"
	fmt.Println(slice2)

	//Arrays Internos

	slice3 := make([]float32,10,11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3,5)
	slice3 = append(slice3,7)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32,5)
	fmt.Println(slice4)
	slice4 = append(slice4,5)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))


}
