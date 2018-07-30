package main

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

func main() {
	s := Request{}
	someDBStruct := SomeStruct{}
	state, _, _ := s.SELECT(someDBStruct).FROM("user").ToString()
	fmt.Println(state)
}

type SomeStruct struct {
	xdx    string `db:"Dont"`
	almost int
	xdx2    string `db:"drink"`
	almost2 int
	xdx3    string `db:"so"`
	almost3 int
	xdx4    string `db:"many"`
}

type Builder interface {
	SELECT(i interface{}) Builder
	FROM(tableName string) Builder
	ToString()(string, []interface{}, error)
}

type Request struct {
	statement string
	err       error
	arguments []interface{}
}

func (r *Request) SELECT(i interface{}) Builder {
	switch v := i.(type) {
	case string:
		r.statement = fmt.Sprintf("SELECT %s", v)
	default:
		if isStruct(i) {
			condition, err := structParserSelect(i)
			if err != nil{
				r.err = err
				break
			}
			r.statement = fmt.Sprintf("SELECT %s", condition)
		} else {
			r.err = errors.New("Cant compaile SELECT part")
		}
	}
	return r
}

func isStruct(i interface{})bool{
	x := reflect.TypeOf(i)
	return x.Kind() == reflect.Struct
}

func (r *Request) FROM(tableName string) Builder {
	r.statement = fmt.Sprintf("%s FROM %s", r.statement, tableName)
	return r
}

func (r *Request) ToString()(string, []interface{}, error) {
	return r.statement, r.arguments, r.err
}

func structParserSelect(st interface{}) (string, error) {
	var (
		statment     = "("
		v                  = reflect.TypeOf(st)
		countParams        = v.NumField()
	)

	for i := 0; i < countParams; i++ {
		field := v.Field(i)

		if data := field.Tag.Get("db"); data != "" {
			statment = fmt.Sprintf("%s%s,", statment, data)
		}
	}
	statment = strings.TrimSuffix(statment, ",")
	statment = statment + ")"
	if statment == "()"{
		return "", errors.New("Did not find any note db")
	}

	return statment, nil
}
