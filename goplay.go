package main

import (
	"code.google.com/p/goweb/goweb"
	"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
	//"fmt"
)

var mHost = "localhost/goplay"
var mSession *mgo.Session

func main() {
	// Setup mongodb connection
	var err error
	mSession, err = mgo.Dial(mHost)
	if err != nil {
		panic(err)
	}
	defer mSession.Close()

	// Optional. Switch the session to a monotonic behavior.
	mSession.SetMode(mgo.Monotonic, true)

	// map the RESTful resources
	goweb.MapRest("/api/{collection}", new(ApiController))

	goweb.ConfigureDefaultFormatters()
	goweb.ListenAndServe(":8080")
}
