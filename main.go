package main

import ( 
	"fmt"
	"net/http" 
)

type UploadBuild struct {}


func (handler UploadBuild) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := "Hallo world"
	fmt.Printf(message)
	fmt.Fprintf(w, message)
}


func main() {
	panic(http.ListenAndServe(":8080", UploadBuild{}))
}
