package archeage

import (
	"net/http"
	"testing"
)

func TestCharactor(t *testing.T) {
	aa := New(&http.Client{})
	_, err := aa.SearchCharactor("KYPROSA", "러블링")
	if err != nil {
		t.Error(err)
	}
}
