package main

import "fmt"

type vehical struct{
	doors int
	color string
}

type truck struct{
	vehical
	fourwheel bool
}

type sedan struct{
	vehical
	luxury bool
}

func main(){

	t1 := truck{
		vehical : vehical{
			doors: 2,
			color: "White",
		},
		fourwheel: true,
	}

	s1 := sedan{
		vehical : vehical{
			doors : 4,
			color:"Gold",
		},
		luxury : false,
	}
 
	fmt.Println(t1)
	fmt.Println(s1)

	fmt.Println(t1.doors)
	fmt.Println(s1.doors)
		

	
}