// Copyright Igor Lesik.
// All Rights Reserved

// Main package.
package main

import (
    "os"
    "path"
    "fmt"
    "html/template"
    "net/http"
)

type Page struct {
    Title   string
    Content string
}

func getPathToTemplate() string {

    pathToBinary := path.Dir(os.Args[0])
    pathToTemplates := path.Join(pathToBinary,"html_templates")
    pathToTemplate := path.Join(pathToTemplates, "exr.html")

    return pathToTemplate
}

func getPathToExercises() string {
    pathToBinary := path.Dir(os.Args[0])
    pathToExercises := path.Join(pathToBinary,"exercise")

    return pathToExercises
}

// Store ready-to-use template in variable
var templateInstance = template.Must(template.ParseFiles(getPathToTemplate()))

// Runs before main and after all packages vars initialized.
func init() {
    fmt.Println("Initializing...")
}

func displayPage(
    writer   http.ResponseWriter,
    request *http.Request) {

    fmt.Println("Request: ", request.URL)

    page := &Page {
        Title: "exercise",
        Content: "content todo",
    }

    templateInstance.Execute(writer, page)
}

func main() {
    exercisesJsonDir := http.Dir(getPathToExercises())
    exrJsonHandler := http.StripPrefix("/exercise/", http.FileServer(exercisesJsonDir))
    http.Handle("/exercise/", exrJsonHandler)

    http.HandleFunc("/", displayPage)
    fmt.Println("Serving HTTP requests")
    http.ListenAndServe(":8080", nil)
}
