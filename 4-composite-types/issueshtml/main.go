package main

import (
	"html/template"
	"log"
	"os"
	"time"

	"github.com/hanchiang/the-go-programming-language/4-composite-types/github/githubissues"
)

const templ = `
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Items}}
<tr>
<td><a href='{{.HtmlUrl}}'>{{.Number}}</td>
<td>{{.State}}</td>
<td><a href='{{.User.HtmlUrl}}'>{{.User.Login}}</a></td>
<td><a href='{{.HtmlUrl}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`

var issueList = template.Must(template.New("issuelist").Parse(templ))

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	result, err := githubissues.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
