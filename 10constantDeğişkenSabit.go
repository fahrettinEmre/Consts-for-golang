//constants   programın çalışma süresi boyunca değişmeyen değerdir. Örneğin pi,ışık hızı faiz oranı gibi.
// const lara değer ataması bir kez yapılır.
package main

import (
	"fmt"
	"math"
)

func main() {

	/* 5
	3.14         // bunlar kendi başına bir const dır.
	"passed"
	true */

	r := 3.0

	const pi float64 = 3.14 // sabit oluşturdur.
	//areaOfCircle := 3.14 * (math.Pow(r, 2)) //üs alma işlemi yapar
	areaOfCircle := pi * (math.Pow(r, 2))
	fmt.Println(areaOfCircle)

	const x int = 2 // değişkenler ile sabitler aynı şekilde isimlendirilir.
	const y float32 = 3.2
	const z string = "test"
	const t bool = false

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", z, z)
	fmt.Printf("%T, %v\n", t, t)
	// const lar compile time a aittir. derleme zamanı
	// değişkenler ise run time aittir. çalışma zamanı

	//age:=40    burada age run time
	//const myAge = age burada const compile zamana ait değil önceden belirli olması lazım çalısmadan önce.

	var g = math.Pow(3, 4) // burada const g = math.Pow(3,4) yazamam. Compile edilmeden sabite değer veremem.
	fmt.Printf("%T, %v\n", g, g)
}
