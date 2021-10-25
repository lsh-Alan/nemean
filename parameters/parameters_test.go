package parameters

import (
	"testing"
)

func TestAccountCommitment(t *testing.T) {
	res, _ := Load()
	if x := res.AccountCommitment.Bases[0][0].X.String(); x != "3300925232884904069615742184086299065053114548682773855645322903933037060492" {
		t.Fatalf("got %s want %s", x, "3300925232884904069615742184086299065053114548682773855645322903933037060492")
	}
	if y := res.AccountCommitment.Bases[0][0].Y.String(); y != "2444473778486452524152098365831852310895020390984340393115418483776314361366" {
		t.Fatalf("got %s want %s", y, "2444473778486452524152098365831852310895020390984340393115418483776314361366")
	}
	if z := res.AccountCommitment.Bases[0][0].Z.String(); z != "5383863615898685798876068106340710353962812440584825112635071552164160655159" {
		t.Fatalf("got %s want %s", z, "5383863615898685798876068106340710353962812440584825112635071552164160655159")
	}
}