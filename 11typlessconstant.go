//typeless constant : her hangi bir veri tipine sahip olmayan sabit
package main

import (
	"fmt"
)

func main() {

	// const x int16 = 5 //burada veri tipi belli
	const x = 5 // burada veri tipi belli değil ama go varsayılan veri tipi olarak değişkene int tipini verir.
	fmt.Printf("%T, %v\n", x, x)

	const y = 3 // burada const y int8=3 yazarsam 16. satır hata verir. Çünkü veri tipi fakrlı.
	//const y yazarak typelees duruma getirir ve hata ortadan kalkar.
	var y1 int16 = 12
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", y1, y1)
	fmt.Printf("%T, %v\n", y+y1, y+y1)

	const a = 5.2 + 4 // 9.2 float64 döner
	fmt.Printf("%T, %v\n", a, a)
	const b = 5.2 + 4.8 // float64 döner
	fmt.Printf("%T, %v\n", b, b)
	const c = int16(5.2 + 4.8) // type convertion
	fmt.Printf("%T, %v\n", c, c)

	//const j int32 = 3
	//const i float32 = 5.6
	const j = 3
	const i = 5.6
	fmt.Printf("%T, %v\n", j, j)
	fmt.Printf("%T, %v\n", i, i)
	fmt.Printf("%T, %v\n", i+j, i+j) // iki const ' da typeless olduğu için varsayılan üzerinden toplar ve değer döner. float64

}
