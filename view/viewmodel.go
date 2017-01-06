package view

import "fmt"
import "html/template"
import "io/ioutil"

// This class will deal with creating our HTML pages by adding the necessary
// assets (like CSS and Javascript) and facilitating the body customization.
type ViewModel struct {
    // This will describe the CSS style of the page
    Style template.CSS

    // This is the footer for the page, described at assets/html/footer.html
    Footer template.HTML

    // This mapping will relate the data produced by the model to the view.
    Body map[string]template.HTML

    // TODO Add Javacript assets
    Script template.JS
}

// Creates a new view model.
func NewViewModel() *ViewModel {
    vm := ViewModel {
        Style: template.CSS(loadCss()),
        Footer: template.HTML(loadFooter()),
        Body: make(map[string]template.HTML),
        Script: template.JS(loadJs()),
    }
    return &vm
}

// Extracts the CSS path
func loadCss() string {
    pwd := GetPwd()
    css, err := ioutil.ReadFile(pwd + "assets/css/app.css")

    if err != nil {
        fmt.Println(err)
        css = []byte { }
    }

    return string(css)
}

// Loads the footer HTML
func loadFooter() string {
    pwd := GetPwd()
    footer, err := ioutil.ReadFile(pwd + "assets/html/footer.html")

    if err != nil {
        fmt.Println(err)
        footer = []byte { }
    }

    return string(footer)
}

// Loads the Javascript asset
func loadJs() string {
    pwd := GetPwd()
    js, err := ioutil.ReadFile(pwd + "assets/js/app.js")

    if err != nil {
        fmt.Println(err)
        js = []byte { }
    }

    return string(js)
}

// TODO Create add body function
func (vm *ViewModel) AddBody(body map[string]template.HTML) {
    for key, value := range body {
        vm.Body[key] = value
    }
}
