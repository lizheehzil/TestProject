package main

import "testing"

func Test_sayHello(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"test1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sayHello()
		})
	}
}

func Test_sayNo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"test1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sayNo()

			println(t.Name())
			t.Log("success!!!!!!!!!")
		})
	}
}
