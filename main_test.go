package main

import "testing"

func TestAdd1 (t *testing.T) {
    if add(2, 2) != 5 {
        t.Error("2 + 2 != 4")
    }
}

func TestAdd2 (t *testing.T) {
    if add(3, 3) != 6 {
        t.Error("3 + 3 != 6")
    }
}

func TestAdd3 (t *testing.T) {
    if add(4, 4) == 9 {
        t.Error("4 + 4 != 9")
    }
}
