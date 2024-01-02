package main

import "github.com/parquet-go/parquet-go"

type User struct {
	Name string
	Age  int
}

func ParquetWrite() {

	users := []User{
		{
			Name: "John",
			Age:  30,
		}, {
			Name: "Jane",
			Age:  25,
		},
	}

	if err := parquet.WriteFile("users.parquet", users); err != nil {
		panic(err)
	}
}
