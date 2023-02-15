package main 
import(
	"fmt"
)

func main(){
	var n bool = true
	fmt.Printf("%v,%T\n",n,n)

	a:= 10  // 1010
	b:= 3   // 0011
	fmt.Println(a+b)
	fmt.Println(a-b)
	fmt.Println(a*b) 
	fmt.Println(a/b)
	fmt.Println(a%b)

	fmt.Println(a & b) // 0010
	fmt.Println(a | b) // 1011
	fmt.Println(a ^ b) // 1001
	fmt.Println(a &^ b) // 0100

	var k complex64 = 1 + 2i
	fmt.Printf("%v , %T" , k,k)

	var j complex128 = complex(5,12)
	fmt.Printf("%v , %T\n",j,j)

	s:= "this is string"
	s2:= "this is also a string"
	s3:= []byte(s)
	fmt.Printf("%v , %T\n", b,b)
	fmt.Printf("%v,%T\n" , s + s2 , s+s2)
	fmt.Printf("%v, %T\n",s[2], s[2])
	
}