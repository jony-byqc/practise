package main

import (
	"bytes"
	"github.com/tinylib/msgp/msgp"
)

type Person struct {
	Name       string `msg:"name" msgpack:"name"`
	Address    string `msg:"address" msgpack:"address"`
	Age        int    `msg:"age" msgpack:"age"`
	Hidden     string `msg:"-" msgpack:"_"` // this field is ignored
	unexported bool   // this field is also ignored
}

func main() {
	var buf bytes.Buffer
	myPerson := Person{
		Name:    "荣锋亮",
		Address: "beijing",
		Age:     33,
	}
	err := msgp.Encode(&buf, &myPerson)
	if err != nil {
		panic("encode some wrong" + err.Error())
	}
	var dstPerson Person
	var dstPerson2 = &Person{}
	err = msgpack.Unmarshal(buf.Bytes(), &dstPerson)
	datas, err := dstPerson2.UnmarshalMsg(buf.Bytes())
	if err != nil {
		panic("uncode:" + err.Error())
	}
	if len(datas) > 0 {
		log.Panicf("%d bytes left over after UnmarshalMsg(): %q", len(datas), datas)
	}
	log.Println("from msgp: ", string(buf.Bytes()))
	log.Printf("msgpack:%v,msgp: %v", dstPerson, *dstPerson2)
	log.Println("from msgpack:", dstPerson.Name)
}
