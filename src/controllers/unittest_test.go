package controllers

import (
	"reflect"
	"testing"

	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
)

func TestAdd(t *testing.T) {
	want := 10

	if result := add(3, 7); !reflect.DeepEqual(result, want) {
		t.Errorf("TestAdd() = %+v, want %+v", result, want)
	}
	t.Log("TestAdd(): success")
}

func TestWithContext(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	t.Logf("appID: %v", appengine.AppID(ctx))
}
