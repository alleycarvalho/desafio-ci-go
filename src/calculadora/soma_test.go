package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(5, 5)
    resultado := 10
    if total != resultado {
       t.Errorf("A soma está errada: %d é diferente de %d!", total, resultado)
    }
}
