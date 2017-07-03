package main

import (
	"log"
	"net/http"
	"io"
	"os"
)

const (
	UPLOAD_DIR = "./uploads"
)

func errorHandler(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(),http.StatusInternalServerError)
	return
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		io.WriteString(w, "<form method=\"POST\" action=\"/upload\" "+
			" enctype=\"multipart/form-data\">"+
			"Choose an image to upload: <input name=\"image\" type=\"file\" />"+ "<input type=\"submit\" value=\"Upload\" />"+
			"</form>")
		return
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			errorHandler(w, err)
		}
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)

		if err != nil {
			errorHandler(w, err)
		}
		defer t.Close()

		if _, err := io.Copy(t,f); err != nil {
			errorHandler(w, err)
		}
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}

}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":9090", nil)
	if err !=nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
