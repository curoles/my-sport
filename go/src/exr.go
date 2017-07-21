// Copyright Igor Lesik.
// All Rights Reserved

// Main package.
package main

import (
    "os"
    "path"
    "fmt"
    "strings"
    "html/template"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type Exercise struct {
    Name        string   `json:"name"`
    Title       string   `json:"title"`
    Repetitions bool     `json:"repetitions"`
    Time        bool     `json:"time"`
    Link        []string `json:"link"`
}

type ExercisePage struct {
    *Exercise
    //Content string
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

func getExerciseStruct(jsonFileName string) (*Exercise,bool) {

    if _, err := os.Stat(jsonFileName); os.IsNotExist(err) {
        fmt.Println("Not found: ", jsonFileName)
        return nil,false
    }

    file, err := ioutil.ReadFile(jsonFileName)
    if err != nil {
        fmt.Printf("File error: %v\n", err)
        return nil, false
    }

    //fmt.Println("File: ", file)

    var exercise Exercise
    unmarshErr := json.Unmarshal(file, &exercise)
    if unmarshErr != nil {
        fmt.Println("JSON unmarshaling error: ", unmarshErr)
        return nil, false
    }


    return &exercise, true
}

func displayPage(
    writer   http.ResponseWriter,
    request *http.Request) {

    fmt.Println("Request: ", request.URL)

    reqURLStr := strings.TrimPrefix(request.URL.String(), "/display")

    jsonFileName := path.Join(getPathToExercises(),reqURLStr+".json")


    exercise, ok := getExerciseStruct(jsonFileName)
    if ok != true {
        http.NotFound(writer, request)
        return
    }

    fmt.Println("Exercise: ", exercise)

    page := &ExercisePage { exercise }
    /*    Title: "exercise",
        Content: "content todo",
    }*/

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
