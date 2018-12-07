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
	if len(ns) == 0 {
		t.Error("empty notice")
	}
	for _, v := range ns {
		if v.Title == "" {
			t.Error(errors.New("Cannot Parse Notice Title"))
		}
		// fmt.Println(v)
	}
}
