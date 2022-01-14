package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
        "html/template"



)

type Page struct {
	Title string
	Body []byte
}

type Todo struct {
    Title string
    Done  bool
}



type PageData struct {

}


func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error){
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil,err
	}
	
	return &Page{Title: title, Body: body}, nil
}



func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func string2Hash(){
	fmt.Printf("Hello");      
}




func main() {

    tmpl0 := template.Must(template.ParseFiles("index.html"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	  
	data := PageData{}
        tmpl0.Execute(w, data)
    })
    

    tmpl3main := template.Must(template.ParseFiles("blockchain.html"))

    http.HandleFunc("/blockchain", func(w http.ResponseWriter, r *http.Request) {
	  
	data := PageData{}
        tmpl3main.Execute(w, data)
    })
    


    tmpl1 := template.Must(template.ParseFiles("hash.html"))
    http.HandleFunc("/hash", func(w http.ResponseWriter, r *http.Request) {
	  
	data := PageData{}
        tmpl1.Execute(w, data)
    })
    
    tmpl2 := template.Must(template.ParseFiles("block.html"))
    http.HandleFunc("/block", func(w http.ResponseWriter, r *http.Request) {
	  
	data := PageData{}
        tmpl2.Execute(w, data)
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}

