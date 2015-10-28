// web server
package main

import (
	"net/http"
	"github.com/kardianos/osext"
	"github.com/pkg/browser"
)

// main is the entry point for the application.
func main() {
	print("Starting - HTTP service on port 8676\n")
	print("Press 'ctrl + c' to quit\n")
	folderPath, err := osext.ExecutableFolder()
    if err != nil {
        print(err)
    }

    browser.OpenURL("http://localhost:8676/index.html")

    panic(http.ListenAndServe(":8676", http.FileServer(http.Dir(folderPath))))

}