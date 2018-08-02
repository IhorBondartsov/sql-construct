package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSELECT(t *testing.T){
	a := assert.New(t)
	expected := "SELECT *"

	sql, args, err := NewConstructor().SELECT("*").ToString()
	a.NoError(err)
	a.Nil(args)
	a.Equal(expected,sql)
}

func TestSELECTStruct(t *testing.T){
	a := assert.New(t)
	expected := "SELECT (id,name)"
	expected3 := "SELECT (id,name,Age)"

	type myStruct struct{
		Id int `db:"id"`
		Name string `db:"name"`
	}

	type myStruct2 struct {
		Id int `db:"id"`
		Name string `db:"name"`
		Login string
	}

	type myStruct3 struct {
		Id int `db:"id"`
		Name string `db:"name"`
		Login string
		Age int `db:"Age"`
	}

	user := myStruct{}
	user2 := myStruct2{}
	user3 := myStruct3{}

	sql, args, err := NewConstructor().SELECT(user).ToString()
	a.NoError(err)
	a.Nil(args)
	a.Equal(expected,sql)

	sql, args, err = NewConstructor().SELECT(user2).ToString()
	a.NoError(err)
	a.Nil(args)
	a.Equal(expected,sql)

	sql, args, err = NewConstructor().SELECT(user3).ToString()
	a.NoError(err)
	a.Nil(args)
	a.Equal(expected3,sql)
}

func TestSELECTError(t *testing.T){
	a := assert.New(t)

	_, args, err := NewConstructor().SELECT("").ToString()
	a.Error(err)
	a.Nil(args)

	_, args, err = NewConstructor().SELECT(345).ToString()
	a.Error(err)
	a.Nil(args)

	_, args, err = NewConstructor().SELECT(345.03).ToString()
	a.Error(err)
	a.Nil(args)
}

func TestFROM(t *testing.T){
	a := assert.New(t)
	expected := "SELECT * FROM table"

	sql, args, err := NewConstructor().SELECT("*").FROM("table").ToString()
	a.NoError(err)
	a.Nil(args)
	a.Equal(expected,sql)
}

func TestFROMError(t *testing.T){
	a := assert.New(t)

	_, args, err := NewConstructor().SELECT("*").FROM("").ToString()
	a.Error(err)
	a.Nil(args)
}