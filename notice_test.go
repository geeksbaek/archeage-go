package archeage

import (
	"errors"
	"net/http"
	"testing"
)

func TestFetchNotice(t *testing.T) {
	aa := New(&http.Client{})
	ns, err := aa.FetchNotice()
	if err != nil {
		t.Error(err)
	}
	for _, v := range ns {
		if v.Title == "" {
			t.Error(errors.New("Cannot Parse Notice Title"))
		}
	}
}
