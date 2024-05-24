package main

import "testing"

func TestSoma(t *testing.T){
	s := Soma(2, 5)
	if s != 7 {
		t.Error("Erro no teste")
	} 
}