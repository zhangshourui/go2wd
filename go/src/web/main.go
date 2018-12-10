package main

import(
	"os"
	"fmt"
	"strings"
	"net/http"
	"html/template"
	"runtime"
	"reflect"
	"net"
	"log"
	"flag"
	"io/ioutil"
	"strconv"
	"path"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"

)

var (
	addr = flag.Bool("addr", false, "find open address and print to final-port.txt")
)
type viewData map[string] interface{} ;

type PageInfo struct{
	Path string
	FilePath string
	ViewData viewData
	Response *http.ResponseWriter
	Request *http.Request
}

var mimeType map[string] string ;
var onActionMap map[string] func(pageInfo *PageInfo);

func Page_Index(pageInfo *PageInfo) {
	var w = * pageInfo.Response;
	var r = pageInfo.Request;
	
	if r.Method=="POST"{

		r.ParseForm()  
		for k, v := range r.Form {
		   // fmt.Fprintf(w, "%s: %s <br/>", k, strings.Join(v, "|"))
		   pageInfo.ViewData[k]=strings.Join(v, "|");
        }

	}else{
		db, err := sql.Open("mysql", "test:123456@localhost/school?charset=utf8")
		checkErr(err)
		defer db.Close()
		rows, err := db.Query("SELECT StudentId, Name FROM sc_student")
        checkErr(err)
	
        for rows.Next() {
            var studentId int
            var name string
          
            err = rows.Scan(&studentId, &name)
            checkErr(err)
          
            fmt.Fprint(w, studentId)
            fmt.Fprint(w, name)
		}
		  
       

    }


}

func checkErr(err error) {
	if err != nil {
		fmt.Println("error found");
		panic(err)
	
	}
}
func RegistPages(){
	onActionMap["/html/index.html"]=Page_Index;
}

func  main()  {
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


func makeHandler(fn func (http.ResponseWriter, *http.Request, *PageInfo)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var filePath = "./site"+r.URL.Path
		/*
		p, err := model.LoadPage(filePath, "Index")
		if err != nil {
			p = &model.PageModel{Title: "Index", Path: filePath
		} else{
			p.Body=[]byte("testasdfasdf");
		}
		*/

		fileSuffix := path.Ext(filePath) //获取文件后缀
		var contentType = GetContentType(fileSuffix)
		if contentType!= "" {
			w.Header().Set("Content-Type", contentType) 
		}

		var onAction = onActionMap[r.URL.Path];
		var vd =make(viewData, 10);
		pageInfo := PageInfo{Path: r.URL.Path, ViewData: vd, Response: &w, Request: r,FilePath: filePath }
		if onAction != nil {
			var funName= GetFunctionName(onAction)
			fmt.Println("onAction:",funName)
			onAction(&pageInfo)
		}else{
			fmt.Println("Static Page: "+r.URL.Path)
		}

		fn(w, r, &pageInfo)

	}
}
func GetFunctionName(i interface{}) string {
    return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func defaultHandler (w http.ResponseWriter, r *http.Request, pageInfo *PageInfo) {

	/*t, _ := template.ParseFiles(pageInfo.Path)
	t.Execute(w, pageInfo)
	*/	
	if strings.HasSuffix(pageInfo.FilePath, ".html"){
		renderTemplate(w, pageInfo.FilePath, pageInfo)
	}else{
		var htmlContent,err = ioutil.ReadFile(pageInfo.FilePath)
		if err != nil{
			http.Error(w, "Not Found: "+r.URL.Path, http.StatusNotFound)
			return;
		}
		fmt.Fprintf(w, string(htmlContent))
	}
}

// render a template
func renderTemplate(w http.ResponseWriter, tmpl string, p *PageInfo) {
	fmt.Printf("%s %s\r\n", p.Request.Method, p.Request.URL.Path)

	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	 log.Println(t.Execute(w, p))

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

	
	fmt.Println("RegistPages :")
	onActionMap=make(map[string] func(pageInfo *PageInfo), 256)
	RegistPages();
	/*
	for k,v :=range onActionMap {
		fmt.Printf("action: %s, %s\r\n", k, GetFunctionName(v))
	}
	*/
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

// get mime for content by extension of a file
// ext with dot of not is ok.
// if no mathes by extension, "" is returned.
func GetContentType(ext string) string{
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