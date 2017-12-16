package archeage

import (
	"net/http"
	"testing"
)

func TestAction(t *testing.T) {
	aa := New(&http.Client{})
	_, err := aa.Auction("TOTAL", "목재", 1)
	if err != nil {
		t.Error(err)
	}
}
