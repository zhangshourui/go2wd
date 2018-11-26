package model

import(
	"fmt"
	"io/ioutil"
)

type PageModel struct{
	Title string
	Body []byte
	Path string
}
func (p PageModel) String() string{
	return fmt.Sprintf("- %s -\r\n%s\r\n---------\r\n%s", p.Title, string(p.Body), p.Path )
}

func LoadPage(path string, title string) (*PageModel, error)   {

	body, err := ioutil.ReadFile(path);
	if err!=nil{
		return nil, err
	}else{
		return &PageModel{Title: title, Body: body, Path: path},nil
	}
}