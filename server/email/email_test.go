package email

import "testing"

func TestSend(t *testing.T) {
	err := Send("linder2015@outlook.com", "test subject", "test body")
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}
