package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

//目前Go支持mongoDB最好的驱动就是mgo
//gopkg.in/mgo.v2

type Person struct {
	Name string //public
	Age  int
}

func main() {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	conn := session.DB("test").C("person")
	err = conn.Insert(&Person{"tom", 23}, &Person{"Lucy", 34})
	//{ "_id" : ObjectId("59ff16b4c1c7d0170b1341d1"), "name" : "tom", "age" : 23 }
	if err != nil {
		log.Fatal("insert ", err)
	}
	result := Person{}
	err = conn.Find(bson.M{"name": "tom"}).One(&result)
	if err != nil {
		log.Fatal("find ", err)
	}
	fmt.Println(result)
}
