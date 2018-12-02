package main

import (
	"fmt"
	"math"
	"time"
	//"strings"
	//"io"
	//"os"
	//"image"
)
// import "./src/tools/kit/getBigger"
 
func _main(){
	//var a =123
	
	// math
	//fmt.Println("srq(5) =",sqrt(5))
/*
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
	*/
	/*
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
	*/
	/*
	var loc1 = Location{X: 1.23, Y: 4.56}
	fmt.Println(loc1)
	*/
	/*
	var s []int
	printSlice("s",s);
	s = append(s, 123, 3,4,5,6,7)
	printSlice("s-a",s)
	fmt.Println(s)

	for i, v := range s {
		fmt.Printf("[%d] = %d\n", i, v)
	}
	fmt.Println("----------")

	for i := range s {
		fmt.Printf("[%d] = %d\n", i, s[i])
	}
	fmt.Println("-2---------")

	for i,_ := range s {
		fmt.Printf("[%d] = %d\n", i, s[i])
	}
	fmt.Println("-2---------")

	for _,v := range s {
		fmt.Printf("[-] = %d\n", v)
	}
	*/
/*
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}

*/
/*
	var line Line = Line{P1: Location { X: 1, Y: 1}, P2: Location{X: 4, Y:5} }	
	fmt.Println("line len: ", line.len())
	*/
/*
	v := Vertex{3, 4}
	//v.Scale(10)
	Scale(v, 10);
	fmt.Println(Abs(v))
	*/

	/*
	var line Line = Line{P1: Location { X: 1, Y: 1}, P2: Location{X: 4, Y:5} }	
	fmt.Println( line)

	var sq, err= Sqrt(-7)
	if err == nil{
		fmt.Println("sqrt over: ", sq);
	}else
	{
		fmt.Println(err)
	}
*/

/*
Validate(MyReader{Str: "AAAB"})
*/

/*
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
		r := rot13Reader{s}
		io.Copy(os.Stdout, &r)
*/
/*
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println("bounds", m.Bounds())
	fmt.Println("rgba: ")
	fmt.Println(m.At(0, 0).RGBA())
*/		
/*
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	fmt.Println("running")	
	go sum(s[:len(s)/2], c)

	fmt.Println("running")
	go sum(s[len(s)/2:], c)

	
	x, y, z:= <- c, <- c, <- c
	fmt.Println	("result is ", x, y, z)
*/

tick := time.Tick(1000 * time.Millisecond)
	boom := time.After(5000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Printf(" . ")
			time.Sleep(50 * time.Millisecond)
		}
	}
}


func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		time.Sleep(1000 * time.Millisecond)
		sum += v
	}
	c <- sum // 将和送入 c
}

/*********functions ********/
type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (v* Vertex) Scale( f float64){
	v.X *= f;
	v.Y *= f;
}
func Scale(v Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}


func fibonacci() func( i int) int {
	var pre1 = 0;
	var pre2 = 0;
	return func(i int) int{
		if i==0{
			return 0
		} else if i== 1{
			pre1 =0
			pre2 =1
			return 1
		}else{
			var r =  pre1+pre2;
			pre1 = pre2;
			pre2 = r;
			return r;
		}
	}
}


type Location struct{
	X float32
	Y float32
}

type Line struct{
	P1 Location
	P2 Location
}
func (loc Location) String() string{
	return fmt.Sprintf("[%v, %v]", loc.X, loc.Y)
}
func (line Line) String() string{
	return fmt.Sprintf("[%v->%v]", line.P1.String(), line.P2.String() )
}

func (line Line) len() float64{
	var v, h = line.P1.X - line.P2.X, line.P1.Y - line.P2.Y;
	return math.Sqrt( float64(v*v +h*h ))
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
func getBigger(a int, b int) int {
	//return a>b? a: b
	if a>b{
		return a;
	} else{
		return b;
	}

}
type DefaultErr struct{
	Time time.Time
	Msg string
} 
func (e DefaultErr) Error() string{
	
	return fmt.Sprintf("[%v]%v", e.Time.Format("2018-01-02 15:04:01"), e.Msg)
	
}
func Sqrt(s int) (float64, error) {
	if s < 0{
		var e DefaultErr = DefaultErr{Time: time.Now(), Msg: "sqrt error: Number must not be negative."};
		return 0, &e
	}
	src := float64(s)
	r := src/2
	//secLen := src/2

	for math.Abs(r*r - src)> 0.000001 {
		var tmp = r*r;
		if tmp == src{
			return r , nil;
		} else if tmp > src{
			//r -= secLen/2	
			fmt.Println(r, " is too big")
			
		} else{
			//r += secLen/2
			fmt.Println(r, " is too small")
		}
		r -= (r*r - src) / (2*r)
	//	secLen /= 2;
	}
	fmt.Println(r, " is OK!")
	return r, nil;
}