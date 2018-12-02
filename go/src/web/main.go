package main

import(
	"os"
	"fmt"
	"strings"
	"net/http"
	"html/template"
	"web/model"
	"net"
	"log"
	"flag"
	"io/ioutil"
	"strconv"
	"path"
)

var (
	addr = flag.Bool("addr", false, "find open address and print to final-port.txt")
)
var mimeType map[string] string ;

func  main()  {
	//	page,_ :=model.LoadPage("./html/index.html", "Home");
	//	fmt.Println(page);

	// http.HandleFunc("/", handler)
	flag.Parse()
	http.HandleFunc("/", makeHandler(defaultHandler))

	loadCfg();

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
		fmt.Printf("%s %s\r\n",r.Method, r.URL.Path)
		var filePath = "./site"+r.URL.Path
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

		fileSuffix := path.Ext(filePath) //获取文件后缀
		var contentType = getContentType(fileSuffix)
		if contentType!= "" {
			w.Header().Set("Content-Type", contentType) 
		}

		fn(w, r, p)

	}
}

func defaultHandler (w http.ResponseWriter, r *http.Request, page *model.PageModel) {

	//t, _ := template.ParseFiles(filePath)
	//t.Execute(w, p)
	if r.Method=="POST"{
		r.ParseForm()  
		fmt.Println(r.Form)

		for k, v := range r.Form {
            fmt.Printf("%s: %s", k, strings.Join(v, "|"))
            fmt.Println("val:", )
        }

	}

	if strings.HasSuffix(page.Path, ".html"){
		renderTemplate(w, page.Path, page)
	}else{
		var htmlContent,err = ioutil.ReadFile(page.Path)
		if err != nil{
			http.Error(w, err.Error(), http.StatusNotFound)
			return;
		}
		fmt.Fprintf(w, string(htmlContent))
	}
}

// render a template
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


func loadCfg(){
	var mimePath = "./cfg/mime.cfg";
	fmt.Println("Loading mime config at",mimePath)
	mimeContent :=  readFile(mimePath)
	mimeContent= strings.Replace(mimeContent, "\r","", -1)
	mimeArr := strings.Split(mimeContent, "\n")
	fmt.Printf("Found %d items", len(mimeArr))
	mimeType = make(map[string]string,len(mimeArr))

	for _,v := range mimeArr{
		//fmt.Println(mimeItems[0],":",mimeItems[1])
		mimeItems := strings.Split(v, "\t")
		if len(mimeItems)==2{
			mimeType[ mimeItems[0]]=mimeItems[1]
			//fmt.Println(mimeItems[0],":",mimeItems[1])
		}
	}
}

func readFile(path string) string {
    fi, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer fi.Close()
    fd, err := ioutil.ReadAll(fi)
    return string(fd)
}
func getContentType(ext string) string{
	if ext[0]!='.' {
		ext = "."+ext;
	}
	
	var contentType, ok = mimeType[ext]
	if ok {
		return contentType
	}else{
		return ""
	}

}