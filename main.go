package main

import (
	"fmt"
	_ "kiva/cmd/app"
)

func hashCode(a []byte){
	h := int64(0)
	for i := 0; i < len(a); i++{
		h = h * 31 + int64(a[i])
		h%=1000000007
//		p = p  * 31
		fmt.Println(int64(a[i]))
	}
	fmt.Println(h)
}

func main() {
	fmt.Println("Hello I am start of Kiva Project")

}
