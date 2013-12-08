package main

import (
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/binding"
)

type BlogPost struct {
  Title   string `form:"title" json:"title" required`
  Content string `form:"content" json:"content"`
  Views   int    `form:"views" json:"views"`
}

// This method implements binding.Validator and is executed by the binding.Validate middleware
func (bp BlogPost) Validate(errors *Errors, req *http.Request) {
  if req.Header().Get("X-Custom-Thing") == "" {
    errors.Overall["x-custom-thing"] = "The X-Custom-Thing header is required"
  }
  if len(bp.Title) < 4 {
    errors.Fields["title"] = "Too short; minimum 4 characters"
  }
  else if len(bp.Title) > 120 {
    errors.Fields["title"] = "Too long; maximum 120 characters"
  }
  if bp.Views < 0 {
    errors.Fields["views"] = "Views must be at least 0"
  }
}

func main() {
  m := martini.Classic()

  //START OMIT
  m.Post("/blog", binding.Bind(BlogPost{}), func(post BlogPost) string {
    // This function won't execute if there were errors
    return post.Title
  })
  //END OMIT

  m.Get("/blog", binding.Form(BlogPost{}), binding.ErrorHandler, func(post BlogPost) string {
    // This function won't execute if there were errors
    return post.Title
  })

  m.Get("/blog", binding.Form(BlogPost{}), func(blogpost BlogPost, err binding.Errors, resp http.ResponseWriter) string {
    // This function WILL execute if there are errors because binding.Form doesn't handle errors
    if err.Count() > 0 {
      resp.WriteHeader(http.StatusBadRequest)
    }
    return blogpost.Title
  })

  m.Post("/blog", binding.Json(BlogPost{}), myOwnErrorHandler, func(blogpost BlogPost) string {
    // By this point, I assume that my own middleware took care of any errors
    return blogpost.Title
  })

  m.Run()
}
