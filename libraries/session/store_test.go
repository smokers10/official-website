package session

import "testing"

func TestStoreMysql(t *testing.T) {
	if Store().MYSQL() == nil {
		t.Error("MYSQL session store FAILED, config return nil")
	} else {
		t.Log("MYSQL session store SUCCESS!")
	}
}
