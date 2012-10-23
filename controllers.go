package main

import (
	"code.google.com/p/goweb/goweb"
	"fmt"
	//"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type ApiController struct {
}

func (cr *ApiController) Create(cx *goweb.Context) {
	col := newModel(cx.PathParams["collection"])
	col.Fill(cx)
	//fmt.Fprintln(cx.ResponseWriter, col.C())
	col.C().Insert(col)
	fmt.Fprintln(cx.ResponseWriter /*col.c().Insert(cx.Request.Form)*/, col)
}

func (cr *ApiController) Delete(id string, cx *goweb.Context) {
	col := newModel(cx.PathParams["collection"])
	col.C().RemoveId(bson.ObjectIdHex(id))
	fmt.Fprintf(cx.ResponseWriter, "Deleted resource %s", id)
}

func (cr *ApiController) DeleteMany(cx *goweb.Context) {
	fmt.Fprintf(cx.ResponseWriter, "Delete many resources")
}

func (cr *ApiController) Read(id string, cx *goweb.Context) {
	col := newModel(cx.PathParams["collection"])
	err := col.C().FindId(bson.ObjectIdHex(id)).One(&col)
	fmt.Fprintf(cx.ResponseWriter, "Read resource %s \n %+v \n Error: %+v", id, col, err)
}

func (cr *ApiController) ReadMany(cx *goweb.Context) {
	switch cx.PathParams["collection"] {
	case "person":
		var result []*Person
	}

	//result := newModels(cx.PathParams["collection"])
	col := newModel(cx.PathParams["collection"])
	col.C().Find(nil).All(&result)
	//fmt.Fprintf(cx.ResponseWriter, "Result: %s", result[0].Id.String())
	//fmt.Fprintf(cx.ResponseWriter, "Result: %s", bson.ObjectIdHex(result[0].Id))
	fmt.Fprintf(cx.ResponseWriter, "Result: %#v", result)
	//fmt.Fprintln(cx.ResponseWriter, result)
	//mSession.DB("").C("Person")
}

func (cr *ApiController) Update(id string, cx *goweb.Context) {
	fmt.Fprintf(cx.ResponseWriter, "Update resource %s", id)
}

func (cr *ApiController) UpdateMany(cx *goweb.Context) {
	fmt.Fprintf(cx.ResponseWriter, "Update many resources")
}
