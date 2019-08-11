//类型变换 数据截断代码

package book

import (
	"fmt"
	"math"
)

func TestMain(){

	fmt.Println("int8 range:",math.MinInt8,math.MaxInt8)
	fmt.Println("int16 range:",math.MinInt16,math.MaxInt16)
	fmt.Println("int32 range:",math.MinInt32,math.MaxInt32)
	fmt.Println("int64 range:",math.MinInt64,math.MaxInt64)
	var  number int32 = 1047483647
	fmt.Printf("nummber====%d\n",number)
	fmt.Printf("int32:0x%x %d\n",number,number)


	a := int16(number)
	fmt.Printf("int16:0x%x %d",a,a)
}
