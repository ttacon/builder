package insertablebuffer

import "testing"

func TestInsert(t *testing.T) {
	buf := NewBufString("hello")
	buf.InsertString(0, "world")
	if buf.String() != "worldhello" {
		t.Error("didn't get back expected string")
	}
}

func TestInsert_InMiddle(t *testing.T) {
	buf := NewBufString("hello")
	buf.InsertString(1, "world")
	if buf.String() != "hworldello" {
		t.Error("didn't get back expected string, got: ", buf.String())
	}
}
