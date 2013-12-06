// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements a stand-alone blog server.

package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

var (
	httpAddr     = flag.String("http", "localhost:9003", "HTTP listen address")
	contentPath  = flag.String("content", "content/", "path to content files")
	templatePath = flag.String("template", "template/", "path to template files")
	staticPath   = flag.String("static", "static/", "path to static files")
)

func main() {
	flag.Parse()
	s, err := NewServer(*contentPath, *templatePath)
	if err != nil {
		log.Fatal(err)
	}
	go reloadDocs(s, *contentPath)
	http.Handle("/", s)
	fs := http.FileServer(http.Dir(*staticPath))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("Now listening on", *httpAddr)
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}

func reloadDocs(s *Server, contentPath string) {

	var err error
	// sleep then reload docs
	for {
		time.Sleep(30 * time.Second)

		// Load content.
		err = s.loadDocs(filepath.Clean(contentPath))
		if err != nil {
			log.Println(err)
		}

		err = s.renderAtomFeed()
		if err != nil {
			log.Println(err)
		}

		err = s.renderJSONFeed()
		if err != nil {
			log.Println(err)
		}
	}

}
