package config

import "testing"

func TestHash(t *testing.T) {
	result, err := Bcrypt().Hash("admin123")
	if err != nil {
		t.Error("Hashing FAILED \n")
		t.Error(err)
	}

	if result == "" {
		t.Error("Hashing FAILED \n")
	}

	if result != "" {
		t.Logf("Hashing SUCCESS. Hashed text is %s", result)
	}
}

func TestCompare(t *testing.T) {
	if Bcrypt().Compare("$2a$10$4V1etw/qc5Evmd10sMJivu1hL241oWtjvjGmDGnUta9GbXCG3G40y", "admin123") {
		t.Logf("Compare SUCCESS. Comparison value %t", true)
	} else {
		t.Errorf("Compare FAILED. Comparison value %t", false)
	}
}
