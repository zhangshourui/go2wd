package main

import(
	"fmt"
	"net/http"
	"html/template"
	"web/model"
)

func  main()  {
//	page,_ :=model.LoadPage("./html/index.html", "Home");
//	fmt.Println(page);

	// http.HandleFunc("/", handler)
	root := "/"
	port := 8080
	fmt.Println("Initializing With Root", root)
	fmt.Println("Listening At ", port)

	http.HandleFunc("/", makeHandler(defaultHandler))
	http.ListenAndServe(":8080", nil)
}


func makeHandler(fn func (http.ResponseWriter, *http.Request, *model.PageModel)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var filePath = "."+r.URL.Path
		fmt.Println("GET", filePath)

		p, err := model.LoadPage(filePath, "")
		if err != nil {
			p = &model.PageModel{Title: "", Path: "N/A"}
			/*
			http.Redirect(w, r, "/edit/"+title, http.StatusFound)
			return
		*/
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
    err = t.Execute(w, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func defaultHandler (w http.ResponseWriter, r *http.Request, page *model.PageModel) {
	

	//t, _ := template.ParseFiles(filePath)
	//t.Execute(w, p)
	renderTemplate(w, page.Path, page)
}