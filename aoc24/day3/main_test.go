package main

import (
	"strings"
	"testing"
)

func TestFind(t *testing.T) {
	testString := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	inputReader := strings.NewReader(testString)

	foundOperations := find(inputReader)

	if len(foundOperations) != 4 {
		t.Error("did not get any hits when we expected some")
	}

}
