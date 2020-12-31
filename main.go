package main

import (
	"fmt"
	"html/template"
	"net/http"
	"log"
)

func IndexHandler(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("Views/index.html")
	if err != nil{
		fmt.Println("error")
	}
	 if err := t.Execute(w,nil); err != nil{
	 	fmt.Println("error")
	 }

}

func main()  {

	http.HandleFunc("/", IndexHandler)
	log.Fatalln(http.ListenAndServe(":8070", nil))

}