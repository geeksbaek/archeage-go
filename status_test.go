package archeage

import (
	"net/http"
	"testing"
)

func TestFetchServerStatus(t *testing.T) {
	aa := New(&http.Client{})
	_, err := aa.FetchServerStatus()
	if err != nil {
		t.Error(err)
	}
}
