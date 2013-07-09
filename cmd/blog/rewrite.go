// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "net/http"

// Register HTTP handlers that redirect old blog paths to their new locations.
func init() {
	for p := range urlMap {
		dest := "/" + urlMap[p]
		http.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, dest, http.StatusMovedPermanently)
		})
	}
}

// keeping these as examples of rewrites should we need them
var urlMap = map[string]string{
	"/2010/03/go-whats-new-in-march-2010.html":               "go-whats-new-in-march-2010",
	"/2013/05/advanced-go-concurrency-patterns.html":         "advanced-go-concurrency-patterns.article",
}
