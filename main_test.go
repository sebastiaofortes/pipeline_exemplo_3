package main

import "testing"

func TestSoma(t *testing.T){
	s := Soma(2, 5)
	if s != 7 {
		t.Error("Erro no teste")
	} 
}

func Test5oma2(t *testing.T){
	s := Soma(5, 5)
	if s != 10 {
		t.Error("Erro no teste")
	} 
}