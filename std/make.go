package main
func main(){
	val := make([]byte,16)
	for idx := range 16{
		println(val[idx],"@",idx)
		idx++
	}
}