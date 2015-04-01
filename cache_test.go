package main

import (
	"testing"
)

func TestGenCacheId(t *testing.T) {
	actual := GenCacheId("test", 1.2)
	expected := "64e94d6c566e13cdb4fcce40c5a5f01f16df664f80f02f5dc78beedd8d339f77"
	if actual.getId() != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
