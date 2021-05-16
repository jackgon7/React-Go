package main

import (
	// "fmt"
	"net/http"

	department "attendance-react/server"

	"github.com/gorilla/mux"
)

func main(){
	r := mux.NewRouter()

	// fs := http.FileServer(http.Dir("./static"))

	// r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "./static/homepage.html")
	// })

	// //r.Handle("/", fs)
	// r.PathPrefix("/js/").Handler(fs)
	dl := &department.DepartmentList{}
	
	r.HandleFunc("/department", dl.AddEmployee).Methods("POST")

	


	


	//s := r.PathPrefix("/").Subrouter()
	

	svc := &http.Server{
		Handler: r,
		Addr: "127.0.0.1:3000",
	}

	svc.ListenAndServe()
}