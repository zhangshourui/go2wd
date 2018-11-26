package main

import(
	"fmt"
	"io"
	"os"
)

type MyReader struct{
	Str string 
}
func (mr MyReader) Read(b []byte)( count int,  err error){
	var buffLen, strLen, readLen int = len(b),len(mr.Str), 0;
	readLen = buffLen;
	if strLen< readLen{
		readLen = strLen;
	}	

	for i:=0; i< readLen; i++{
		b[i] = mr.Str[i]
		count++
	}
	return count, nil
	
}


func Validate(r io.Reader) {
    b := make([]byte, 1024, 2048)
    i, o := 0, 0
    for ; i < 1<<20 && o < 1<<20; i++ { // test 1mb
        n, err := r.Read(b)
        for i, v := range b[:n] {
            if v != 'A' {
                fmt.Fprintf(os.Stderr, "got byte %x at offset %v, want 'A'\n", v, o+i)
                return
            }
        }
        o += n
        if err != nil {
            fmt.Fprintf(os.Stderr, "read error: %v\n", err)
            return
        }
    }
    if o == 0 {
        fmt.Fprintf(os.Stderr, "read zero bytes after %d Read calls\n", i)
        return
    }
    fmt.Println("OK!")
}

type rot13Reader struct {
	r io.Reader
}


func (rt rot13Reader) Read(b []byte)( int, error){

	var tempb []byte = make([] byte, 50)
	count, errs := rt.r.Read(tempb);
	if errs == nil{
		fmt.Println("tempb len",count)
		for i:=0; i<count;i++ {
			b[i] =rot13(tempb[i])
		}
		return count, errs
	} else{
		return 0, io.EOF
	}

	
}
func rot13(p byte) byte {
    switch {
    case p >= 'A' && p <= 'M':
        p += 13
    case p >= 'N' && p <= 'Z':
        p -= 13
    case p >= 'a' && p <= 'm':
        p += 13
    case p >= 'n' && p <= 'z':
        p -= 13
    }
    return p
}