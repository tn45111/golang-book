package main

import "fmt"

func main() {

	// Long delaration
	const x string = "Hello, World"
	fmt.Println(x)

	// cannot change
	// x = "change it"

	var y string
	y = "Hello, World"
	fmt.Println(y)

	// short declaration
	// Type inference
	z:="Hello, World"
	fmt.Println(z)
	fmt.Printf("Type: %T\n", z)

	z2:="abc"
	fmt.Println(z2)
	fmt.Printf("Type: %T\n", z2)

	z2="1"
	fmt.Println(z2)
	fmt.Printf("Type: %T\n", z2)


	z3:=true
	fmt.Println(z3)
	fmt.Printf("Type: %T\n", z3)


	// define multiple var
	var (
		a1=1
		b1=44
		c1=66
	)
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)

	v1,v2:="สวัสดีจ้า","ครับผม"
	fmt.Println(v1)
	fmt.Println(v2)
	
	t1,t2:=v2,v1
	fmt.Println(t1)
	fmt.Println(t2)

	// no need to use temp
	t1,t2=t2,t1
	fmt.Println(t1)
	fmt.Println(t2)
}

