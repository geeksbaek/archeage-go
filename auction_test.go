package archeage

import (
	"net/http"
	"testing"
)

func TestAction(t *testing.T) {
	aa := ArcheAge(&http.Client{})
	_, err := aa.Auction("TOTAL", "목재")
	if err != nil {
		t.Error(err)
	}
}
