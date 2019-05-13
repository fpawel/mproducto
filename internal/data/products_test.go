package data

import (
	"encoding/json"
	"os"
	"testing"
)

func TestGetProductsCategoryNodes(t *testing.T) {
	db,err := NewConnectionDB(DefaultPgConfig())
	if err!=nil{
		t.Fatal(err)
	}
	xs := GetProductsCategoryNodes(db)


	b,err := json.MarshalIndent(xs, "", "    ")
	if err!=nil{
		t.Fatal(err)
	}

	f, err := os.Create("products-categories.ts")
	if err!=nil{
		t.Fatal(err)
	}
	defer f.Close()

	_,_ = f.WriteString("let productsCategoriesTree = \n")
	_,_ = f.Write(b)
	_,_ = f.WriteString(";")
}
