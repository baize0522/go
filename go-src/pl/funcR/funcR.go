package  funcR


func Reverse(r []int) []int{
	for a,b := 0,len(r)-1;a < len(r)/2;a,b = a+1,b-1{
		r[a],r[b] = r[b],r[a]
	}
	return []int(r)
}
