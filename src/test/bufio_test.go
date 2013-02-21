package main

import (
	"fmt"
	"reflect"
	)
type T struct{
     d int
}


type Test interface{
	test(i int)
}

func (t *T) test(i int){
	t.d +=1
	fmt.Println(i)
}

func testInterface(t Test){
     t.test(1)
     fmt.Println(reflect.TypeOf(t))
}

func main() {
	//t1 := new(T)
	//t1.test(1)
	//(*t1).test(1)
	//fmt.Println(t1.d)
        //testInterface(t1)

	//var t2 T
	//t2.test(1) 
	//t2.test(1) 
	//(&t2).test(1)
	//(&t2).test(1)
	
	//fmt.Println(t2.d)
	//testInterface(t2)
	
	//t1 := new(T)
	//testInterface(t1)
	//testInterface(t1)
	
	//fmt.Println(t1.d)
	
	//t1.test(1)
	//t1.test(1)
	//fmt.Println(t1.d)

	var t2 T
	testInterface(t2)
	testInterface(t2)
	
}
