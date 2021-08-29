package main

import (
	"fmt"
	"testing"
)


func TestPool(t *testing.T) {
	conn := CreateConnection()
	defer conn.Close()
	// Open doesn't open a connection. Validate DSN data:
	err := conn.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	t.Log("Ping OK!")
}


func TestTagDAO_PrepareInsert(t *testing.T) {
	dao := TagDAO{
		conn: CreateConnection(),
	}
	for i := 0; i < 5; i++ {
		dao.PrepareInsert(fmt.Sprintf("Test%d", i))
	}
}

func TestTagDAO_PrepareSelect(t *testing.T) {
	dao := TagDAO{
		conn: CreateConnection(),
	}
	data, err := dao.PrepareSelect()
	if err != nil {
		fmt.Println(err)
	}
	size := len(data)
	for i := 0; i < size; i++ {
		fmt.Println(data[i])
	}
}