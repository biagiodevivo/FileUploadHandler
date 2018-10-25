package main 

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	
)

func main(){
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)

	log.Print("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}	

func foo(w http.ResponseWriter, req *http.Request){

	var s string
	if req.Method == http.MethodPost {
		f, _, err := req.FormFile("userFile")
		if err != nil {
			log.Println(err)
			http.Error(w, "UPLOAD_ERROR", http.StatusInternalServerError)
			return
		}
		defer f.Close()

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			log.Println(err)
			http.Error(w, "ERROR_READING_FILE", http.StatusInternalServerError)
			return 
		}

		s = string(bs)

	}

	w.Header().Set("CONTENT-TYPE", "text/html; charset=UTF-8")
	fmt.Fprintf(w,`<form action="/" method="post" enctype="multipart/form-data">
	Upoad a File<br>
	<input type="file" name="userFile"><br>
	<input type="submit">
	</form>
	<br>
	<br>
	<h1>%v</h1>`,s)
}