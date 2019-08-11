package funcR

import (
	"fmt"
)

func  Test(){
	var (
		name  string
		score  int
	)
	name  =  "lt"
	score = 634
	fmt.Println(name,score)

	/*=====================================
	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)


	//=====================================*/

	a := true
	fmt.Printf("%T,%v\n",a,a)
	fmt.Printf("%T,%v\n",name,name)
	fmt.Printf("%T,%v\n",score,score)

	for tmp:= score;tmp > 0;tmp-- {
		fmt.Println(tmp)
		fmt.Println("\n---------------\n")
	}
}
