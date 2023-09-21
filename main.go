package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

var (
    line_wise []string
)

type Quotes struct{
    Quote string
    Author string
}

type handler_log struct{

  httpHandler http.Handler  

}

func (lh handler_log) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming connection: RemoteAddr=%s, Method=%s, URL=%s", r.RemoteAddr, r.Method, r.URL.Path)
	lh.httpHandler.ServeHTTP(w, r)
}

func homePage(w http.ResponseWriter,r *http.Request){
    
    tmp_arr:=strings.Split(line_wise[rand.Intn(100)],".")

    var quotes Quotes
    quotes.Quote,quotes.Author=tmp_arr[0],tmp_arr[1]
    // fmt.Fprintf(w,"<h1>%s %s</h1>",quote_author[0],quote_author[1])
    // fmt.Fprintf(w,"<h1>%s</h1>",line_wise[rand.Intn(99)])
    templ,_:=template.ParseFiles("./template/template.html")
    templ.Execute(w,quotes)
	log.Printf("Incoming connection: RemoteAddr=%s, Method=%s, URL=%s", r.RemoteAddr, r.Method, r.URL.Path)

    // fmt.Fprintf(w,"welcome to homepage\n")
    // fmt.Fprintln(w,"Endpoint Hit:homepage")
    // fmt.Fprintln(w,r.URL.Path)

}

// func notFound(w http.ResponseWriter,r *http.Request)  {
// 	log.Printf("Incoming connection: RemoteAddr=%s, Method=%s, URL=%s", r.RemoteAddr, r.Method, r.URL.Path)
//     w.WriteHeader(http.StatusNotFound)
//     fmt.Fprintf(w,"404 - Page not found")
// }

func handleRequests(){

    // http.HandleFunc("/*",notFound)
    http.HandleFunc("/quote",homePage)

    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./template/css/"))))

    bind_address:="0.0.0.0:8080" 
    fmt.Println("listening on port",bind_address) 
    log.Fatal(http.ListenAndServe(bind_address,nil))
}

func loadFile(){

    byte_arry,err:=ioutil.ReadFile("./template/stoic-quotes.csv")
    if(err != nil){
        log.Fatal(err.Error())
    }

    line_wise=strings.Split(string(byte_arry), "\n")
    line_wise=line_wise[:len(line_wise)-1]

}

func main()  {

    loadFile()
    handleRequests()

}
