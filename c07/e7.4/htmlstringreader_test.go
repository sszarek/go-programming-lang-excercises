package htmlstringreader

import (
	"testing"

	"golang.org/x/net/html"
)

func Test_Read(t *testing.T) {
	_, err := html.Parse(&htmlstringreader{"<html></html>"})

	if err != nil {
		t.Error(err)
	}
}
