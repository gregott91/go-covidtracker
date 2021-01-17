package main

import (
	"bufio"
	"html/template"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	example()
}

func example() {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>`

	f, err := os.Create(os.Args[1])
	check(err)

	defer f.Close()

	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}

	w := bufio.NewWriter(f)
	err = t.Execute(w, data)
	check(err)

	noItems := struct {
		Title string
		Items []string
	}{
		Title: "My another page",
		Items: []string{},
	}

	err = t.Execute(w, noItems)
	check(err)

	w.Flush()
}
