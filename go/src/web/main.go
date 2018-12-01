package main

import(
	"fmt"
	"net/http"
	"html/template"
	"web/model"
	"net"
	"log"
	"flag"
	"io/ioutil"
	"strconv"
)

var (
	addr = flag.Bool("addr", true, "find open address and print to final-port.txt")
)

func  main()  {
//	page,_ :=model.LoadPage("./html/index.html", "Home");
//	fmt.Println(page);

	// http.HandleFunc("/", handler)
	flag.Parse()
	http.HandleFunc("/", makeHandler(defaultHandler))

	if *addr {
		fmt.Println("Listen 127.0.0.1:0")
		l, err := net.Listen("tcp", "127.0.0.1:8081")
   		if err != nil {
   			log.Fatal(err)
   		}
   		err = ioutil.WriteFile("final-port.txt", []byte(l.Addr().String()), 0644)
   		if err != nil {
   			log.Fatal(err)
   		}
   		s := &http.Server{}
   		s.Serve(l)
   		return
   	}

	root := "/"
	port := 8080
	fmt.Println("Initializing With Root", root)
	fmt.Println("Listening :"+strconv.Itoa(port))

	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}


func makeHandler(fn func (http.ResponseWriter, *http.Request, *model.PageModel)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var filePath = "."+r.URL.Path
		fmt.Printf("%s %s\r\n",r.Method, filePath)

		p, err := model.LoadPage(filePath, "Index")
		if err != nil {
			p = &model.PageModel{Title: "Index", Path: filePath}
			/*
			http.Redirect(w, r, "/edit/"+title, http.StatusFound)
			return
		*/
		} else{
			p.Body=[]byte("testasdfasdf");
		}


		fn(w, r, p)

	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *model.PageModel) {
   	t, err := template.ParseFiles(tmpl)
	//t.Execute(w, p)
	
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
//	err = t.Execute(w, p)
	log.Println(t.Execute(w, p))
   // if err != nil {
   //     http.Error(w, err.Error(), http.StatusInternalServerError)
   // }
}

func defaultHandler (w http.ResponseWriter, r *http.Request, page *model.PageModel) {
	

	//t, _ := template.ParseFiles(filePath)
	//t.Execute(w, p)
	renderTemplate(w, page.Path, page)
}