package archeage

import (
	"net/http"
	"testing"
)

func TestFetchNotice(t *testing.T) {
	aa := ArcheAge(&http.Client{})
	_, err := aa.FetchNotice()
	if err != nil {
		t.Error(err)
	}
}
