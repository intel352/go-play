package main

import (
	"code.google.com/p/goweb/goweb"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	//"reflect"
)

type Model interface {
	C() *mgo.Collection
	Fill(*goweb.Context)
}

type Person struct {
	Id    bson.ObjectId `bson:"_id,omitempty"`
	Name  string
	Phone string
}

func (p Person) C() *mgo.Collection {
	return Coll("Person")
}

func (p *Person) Fill(cx *goweb.Context) {
	cx.Fill(&p)
}

func newPerson() *Person {
	return &Person{}
}

type Product struct {
	Id    bson.ObjectId `bson:"_id,omitempty"`
	Name  string
	Price float32
}

func (p Product) C() *mgo.Collection {
	return Coll("Product")
}

func (p *Product) Fill(cx *goweb.Context) {
	cx.Fill(&p)
}

func (p Product) TestFunc(s string) string {
	return s
}

func newProduct() *Product {
	return &Product{}
}

func newModel(c string) Model {
	switch c {
	case "person":
		return newPerson()
	}
	return nil
}

func Coll(c string) *mgo.Collection {
	return mSession.DB("").C(c)
}
