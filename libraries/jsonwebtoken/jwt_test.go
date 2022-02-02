package jsonwebtoken

import "testing"

func TestSign(t *testing.T) {
	payload := &Payload{
		ID:       1,
		Email:    "admin@admin.com",
		IsLogged: true,
	}

	token, err := Init(payload).Sign()
	if err != nil {
		t.Errorf("Sign FAILED \n")
		t.Error(err)
	}

	if token == "" {
		t.Errorf("Sign FAILED. Token is empty string! \n")
	}

	if token != "" {
		t.Logf("Sign SUCCESS. Token value is %s \n", token)
	}
}

func TestVerificate(t *testing.T) {
	jwt := Init(&Payload{})
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImFkbWluQGFkbWluLmNvbSIsIklEIjoxLCJJc0xvZ2dlZCI6dHJ1ZX0.O1d6ndbpW8XUMXOqADBB8lTL5fbMU2irlbaOMqaMByc"

	claims, err := jwt.Verificate(token)
	if err != nil {
		t.Errorf("Sign FAILED \n")
		t.Error(err)
	}

	if claims.Email == "admin@admin.com" && claims.ID == 1 && claims.IsLogged == true {
		t.Logf("Verificate SUCCESS. The Payload is email = %s, id = %d, is logged = %t \n", claims.Email, claims.ID, claims.IsLogged)
	} else {
		t.Errorf("Varificate FAILED. The payload is not meet required payload \n")
	}
}

func BenchmarkTestVerificate(b *testing.B) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImFkbWluQGFkbWluLmNvbSIsIklEIjoxLCJJc0xvZ2dlZCI6dHJ1ZX0.O1d6ndbpW8XUMXOqADBB8lTL5fbMU2irlbaOMqaMByc"
	for i := 0; i < b.N; i++ {
		Init(&Payload{}).Verificate(token)
	}
}

func BenchmarkSign(b *testing.B) {
	payload := &Payload{
		ID:       1,
		Email:    "admin@admin.com",
		IsLogged: true,
	}

	for i := 0; i < b.N; i++ {
		Init(payload).Sign()
	}
}
