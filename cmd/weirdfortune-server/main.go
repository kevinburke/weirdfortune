package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"strings"
	"text/template"

	"github.com/kevinburke/handlers"
	weirdfortune "github.com/kevinburke/weirdfortune/lib"
)

const DefaultPort = "9608"

const tplString = `<!doctype html>
<html lang="en">
  <head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css" integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">

    <title>weirdfortune</title>
  </head>
  <body style="background-color: #0d47a1;">
	<div style="margin-top: 30px;" class="container">
		<br />
		<div class="row">
		<div style="background-color: white; padding-top: 10px; border-radius: 5px;" class="col-md-5 offset-md-3">
			<h3>{{ .Author }}</h3>
			<br />
			<p style="font-size: 1.5rem">{{ .Body }}</p>
		</div>
	</div>
  </body>
</html>
`

type tplData struct {
	Author string
	Body   string
}

func main() {
	flag.Parse()
	addr := net.JoinHostPort("localhost", DefaultPort)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	tpl, err := template.New("body").Parse(tplString)
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		choice, err := weirdfortune.Fortune()
		if err != nil {
			panic(err)
		}
		usertweet := strings.SplitN(choice, ": ", 2)
		tpl.Execute(w, tplData{Author: usertweet[0], Body: usertweet[1]})
	})
	handlers.Logger.Info("Started server", "addr", addr)
	http.Serve(ln, handlers.Log(handlers.Server(mux, "weirdfortune/0.1")))
}
