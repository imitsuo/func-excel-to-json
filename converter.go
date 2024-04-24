package converter

import (
	// "fmt"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/FerdinaKusumah/excel2json"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("ConvertHTTP", ConvertHTTP)
}

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func ConvertHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	const maxMemory = 2 * 1024 * 1024 // 2 megabytes.

	// if r.ContentLength <= maxMemory {
	// 	http.Error(w, "File size is bigger than 2mb", http.StatusBadRequest)
	// 	return
	// }
	file, _ := os.CreateTemp("", "cnv")
	defer file.Close()

	io.Copy(file, r.Body)
	// bytes, _ := io.ReadAll(r.Body)
	// log.Printf("%d bytes: %s\n", 100, string(bytes[:100]))

	result, _ := excel2json.GetExcelFilePath(file.Name(), "Sheet1", nil)

	for _, val := range result {
		result, _ := json.Marshal(val)
		fmt.Println(string(result))
	}
	// if err := r.ParseMultipartForm(maxMemory); err != nil {
	//         'http.Error(w, "Unable to parse form", http.StatusBadRequest)'
	//         'log.Printf("Error parsing form: %v", err)'
	//         'return'
	// }

	// defer func() {
	//         if err := r.MultipartForm.RemoveAll(); err != nil {
	//                 http.Error(w, "Error cleaning up form files", http.StatusInternalServerError)
	//                 log.Printf("Error cleaning up form files: %v", err)
	//         }
	// }()

	// // r.MultipartForm.File contains *multipart.FileHeader values for every
	// // file in the form. You can access the file contents using
	// // *multipart.FileHeader's Open method.
	// for _, headers := range r.MultipartForm.File {
	//         for _, h := range headers {
	//                 fmt.Fprintf(w, "File uploaded: %q (%v bytes)", h.Filename, h.Size)
	//                 // Use h.Open() to read the contents of the file.
	//                 bytes := make([]byte, h.Size)
	//                 fp, _ := h.Open()
	//                 n,_ := fp.Read(bytes)
	//                 fmt.Printf("%d bytes: %s\n", n, string(bytes[:n]))
	//         }
	// }

	// var d struct {
	//         Name string `json:"name"`
	// }
	// if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
	//         fmt.Fprint(w, "Hello, World!")
	//         return
	// }
	// if d.Name == "" {
	//         fmt.Fprint(w, "Hello, World!")
	//         return
	// }
	// fmt.Fprintf(w, "Hello, %s!", html.EscapeString(d.Name))
}
