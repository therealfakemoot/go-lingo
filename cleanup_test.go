package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func Test_StripHeaders(t *testing.T) {
	raw, err := os.Open("corpus/test/headers.input")
	if err != nil {
		t.Logf("error opening header input file: %s", err)
		t.Fail()
	}
	expected, err := os.Open("corpus/test/headers.expected")
	if err != nil {
		t.Logf("error opening header expected file: %s", err)
		t.Fail()
	}
	expectedBytes, err := ioutil.ReadAll(expected)
	if err != nil {
		t.Logf("error loading `expected` byte slice: %s", err)
		t.Fail()
	}

	actual, err := StripHeaders(raw)
	if err != nil {
		t.Logf("error stripping headers: %s", err)
		t.Fail()
	}
	actualBytes, err := ioutil.ReadAll(actual)
	if err != nil {
		t.Logf("error loading `actual` byte slice: %s", err)
		t.Fail()
	}

	if bytes.Compare(expectedBytes, actualBytes) != 0 {
		ioutil.WriteFile("latest.actual", actualBytes, 0644)
		t.Log("StripHeaders mistmatch")
		t.Fail()
	}

}
