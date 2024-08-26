// for this to work, you have to do
// go mod init unit_test_cube_func
// then : go test
// output will look like:
//
// PASS
// ok      unit_test_cube_func     0.555s

package main

import "testing"

func TestCube(t *testing.T) {
	result := cube(10)
	if result != 1000 {
		t.Error("Test failed , was expecting 1000")
	}
}
