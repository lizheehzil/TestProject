package main

import "testing"

func TestHelloWorld(t *testing.T) {

	for i := 0; i < 10; i++ {
		t.Log("hello world!")
	}
}

func TestMethod1(t *testing.T) {
	for i := 0; i < 20; i++ {
		t.Log("This is method1 ")
	}
}
