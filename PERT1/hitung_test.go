package main

import "testing"

// membuat use test untuk fungsi luas dan keliling 
func TestLuas(t *testing.T) {
	result := luas(3, 7)
	if result != 10.5 {
		t.Logf("Hasil perhitungan luas segitiga adalah %f, seharusnya %f", result, 10.5)
		t.Fail()
	}
}

// make a copy TestLuas & change Luas --> Keliling
func TestKeliling(t *testing.T) {
	result := keliling(6.5)
	if result != 19.5 {
		t.Logf("Hasil perhitungan keliling segitiga adalah %f, seharusnya %f", result, 19.5)
		t.Fail()
	}
}