package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"labix.org/v2/mgo"
)

//START0 OMIT
type Wish struct {
	Name        string `form:"name"`
	Description string `form:"description"`
}
//END0 OMIT

//START1 OMIT
func DB() martini.Handler {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	return func(c martini.Context) {
		s := session.Clone()
		c.Map(s.DB("advent"))
		defer s.Close()
		c.Next()
	}
}
//END1 OMIT

//START2 OMIT
func GetAll(db *mgo.Database) []Wish {
	var wishlist []Wish
	db.C("wishes").Find(nil).All(&wishlist)
	return wishlist
}
//END2 OMIT

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
  //START3 OMIT
	m.Use(DB())
  //END3 OMIT

  //START4 OMIT
	m.Get("/wishes", func(r render.Render, db *mgo.Database) {
		r.HTML(200, "list", GetAll(db))
	})
  //END4 OMIT

  //START5 OMIT
	m.Post("/wishes", binding.Form(Wish{}), func(wish Wish, r render.Render, db *mgo.Database) {
		db.C("wishes").Insert(wish)
		r.HTML(200, "list", GetAll(db))
	})
  //END5 OMIT

	m.Run()
}
