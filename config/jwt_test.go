package config

import "testing"

func TestSign(t *testing.T) {
	payload := &JsonWebToken{
		ID:       1,
		Email:    "admin@admin.com",
		IsLogged: true,
	}

	token, err := payload.Sign()
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
	payload := &JsonWebToken{
		ID:       1,
		Email:    "admin@admin.com",
		IsLogged: true,
	}

	result := &JsonWebToken{}

	token, err := payload.Sign()
	if err != nil {
		t.Errorf("Verificate FAILED. Error when sign token \n")
		t.Error(err)
	}

	if token == "" {
		t.Errorf("Verificate FAILED. Token is empty string! \n")
	}

	dumpResult, err := result.Verificate(token)
	if err != nil {
		t.Errorf("Sign FAILED \n")
		t.Error(err)
	}

	if dumpResult.Email == "admin@admin.com" && dumpResult.ID == 1 && dumpResult.IsLogged == true {
		t.Logf("Verificate SUCCESS. The Payload is email = %s, id = %d, is logged = %t \n", dumpResult.Email, dumpResult.ID, dumpResult.IsLogged)
	} else {
		t.Errorf("Varificate FAILED. The payload is not meet required payload \n")
	}
}
