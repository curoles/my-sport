// Copyright Igor Lesik.
// All Rights Reserved

// Main package.
package main

import (
    "os"
    "path"
    "html/template"
    "net/http"
)

type Page struct {
    Title   string
    Content string
}

func displayPage(
    writer   http.ResponseWriter,
    request *http.Request) {

    page := &Page {
        Title: "exercise",
        Content: "content todo",
    }

    pathToBinary := path.Dir(os.Args[0])
    pathToTemplates := path.Join(pathToBinary,"html_templates")
    pathToTemplate := path.Join(pathToTemplates, "exr.html")

    template := template.Must(template.ParseFiles(pathToTemplate))

    template.Execute(writer, page)
}

func main() {
    http.HandleFunc("/", displayPage)
    http.ListenAndServe(":8080", nil)
}
