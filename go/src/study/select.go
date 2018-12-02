package main

import "fmt"


func __main(){
	ch := make(chan int)
	cr := 0
	go func(){
		for{
		 	select {
				case ch <- cr:
				cr++
				fmt.Println("index=",cr)
			}
		}
	}()
	for{
		var r int = 0;
		select {
		case  r=<-ch:		   
		   fmt.Println("result=",r)
	   }
   }
	
}