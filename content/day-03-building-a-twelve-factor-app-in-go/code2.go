package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

var configHTML = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>MyApp Config</title>
  <link href="//netdna.bootstrapcdn.com/bootstrap/3.0.2/css/bootstrap.min.css" rel="stylesheet">
</head>
<body>
  <h2>Configuration Settings</h2>
  <table class="table table-bordered table-striped table-condensed">
    <tr><th>Name</th><th>Value</th></tr>
    <tr><td>Debug</td><td>{{ .Debug }}</td></tr>
    <tr><td>Host</td><td>{{ .Host }}</td></tr>
    <tr><td>Password</td><td>{{ .Password }}</td></tr>
    <tr><td>Port</td><td>{{ .Port }}</td></tr>
    <tr><td>Timeout</td><td>{{ .Timeout }}</td></tr>
    <tr><td>Username</td><td>{{ .Username }}</td></tr>
  </table>
</body>
</html>`

// Spec represents the myapp configuration.
type Spec struct {
	Debug    bool
	Host     string
	Password string
	Port     string
	Timeout  uint
	Username string
}

// spec holds the myapp configuration.
var spec Spec

func ConfServer(w http.ResponseWriter, req *http.Request) {
	t := template.New("configHTML")
	t.Parse(configHTML)
	t.Execute(w, spec)
}

func main() {
	err := envconfig.Process("myapp", &spec)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/config", ConfServer)
	listenAddr := net.JoinHostPort(spec.Host, spec.Port)
	if err = http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatal(err)
	}
}
