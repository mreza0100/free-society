package test

import (
	"testing"
)

func CheckFail(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func FailIf(t *testing.T, fail bool, why error) {
	if fail {
		CheckFail(t, why)
	}
}
