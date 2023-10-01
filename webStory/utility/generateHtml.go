package utility

import (
	"html/template"
	"log"
	"net/http"
)

const StoryPages = `
<!DOCTYPE html>
	<html>
	<head>
		<title>{{.Title}}</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		{{range .Story}}
			<p>{{.}}</p>
		{{end}}
		<h2>Options:</h2>
		<ul>	
			{{range .Options}}
				<li><a href="{{.Arc}}">{{.Text}}</a></li>
			{{end}}
		</ul>
	</body>
	</html>
	`

const IndexPage = `
<!DOCTYPE html>
	<html>
	<head>
		<title>Home</title>
	</head>
	<body>
		<h1>Welcome</h1>
		
		<h2>Click on link to start the story</h2>
		<ul>	
			
				<li><a href="./intro">Intro</a></li>
			
		</ul>
	</body>
	</html>
	`

func CreateTemplate(tpl string) *template.Template {

	htmlTemplate, err := template.New("webpage").Parse(tpl)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return htmlTemplate

}

func CreateHandler(htmlTemplate *template.Template, pageData Intro) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		err := htmlTemplate.Execute(w, pageData)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func check(err error) {

}
