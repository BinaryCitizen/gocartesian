package gocartesian

import (
	"testing"
	"strings"
	"fmt"
)

type testset struct {
	input [][]string
	expect []string
}

var tests = []testset {
	{
		[][]string{[]string{"1"}, []string{"A"}, []string{"X"}},
		[]string{"1,A,X"},
	},
	{
		[][]string{[]string{"P"}},	
		[]string{"P"},
	},
	{
		[][]string{[]string{"P", "A"}, []string{"2"}},
		[]string{"P,2", "A,2"},
	},
	{
		[][]string{[]string{"1"}, []string{"A"}, []string{"X"}, []string{"Z"}, []string{"1"}, []string{"A"}, []string{"X"}, []string{"Z"}},
		[]string{"1,A,X,Z,1,A,X,Z"},
	},
	{
		[][]string{[]string{"1", "2", "3"}, []string{"A", "B"}, []string{"X", "Y", "Z"}},
		[]string{"1,A,X", "1,A,Y", "1,A,Z", "1,B,X", "1,B,Y", "1,B,Z", "2,A,X", "2,A,Y", "2,A,Z", "2,B,X", "2,B,Y", "2,B,Z", "3,A,X", "3,A,Y", "3,A,Z", "3,B,X", "3,B,Y", "3,B,Z"},
	},
	{
		[][]string{[]string{"L"}, []string{"O", "1", "Q", "8"}, []string{"X", "Y", "Z"}, []string{"T"}, []string{"3", "9", "G", "0"}},
		[]string{"L,O,X,T,3", "L,O,X,T,9", "L,O,X,T,G", "L,O,X,T,0", "L,O,Y,T,3", "L,O,Y,T,9", "L,O,Y,T,G", "L,O,Y,T,0", "L,O,Z,T,3", "L,O,Z,T,9", "L,O,Z,T,G", "L,O,Z,T,0", "L,1,X,T,3", "L,1,X,T,9", "L,1,X,T,G", "L,1,X,T,0", "L,1,Y,T,3", "L,1,Y,T,9", "L,1,Y,T,G", "L,1,Y,T,0", "L,1,Z,T,3", "L,1,Z,T,9", "L,1,Z,T,G", "L,1,Z,T,0", "L,Q,X,T,3", "L,Q,X,T,9", "L,Q,X,T,G", "L,Q,X,T,0", "L,Q,Y,T,3", "L,Q,Y,T,9", "L,Q,Y,T,G", "L,Q,Y,T,0", "L,Q,Z,T,3", "L,Q,Z,T,9", "L,Q,Z,T,G", "L,Q,Z,T,0", "L,8,X,T,3", "L,8,X,T,9", "L,8,X,T,G", "L,8,X,T,0", "L,8,Y,T,3", "L,8,Y,T,9", "L,8,Y,T,G", "L,8,Y,T,0", "L,8,Z,T,3", "L,8,Z,T,9", "L,8,Z,T,G", "L,8,Z,T,0"},
	},
	{
		[][]string{[]string{"P", "A"}, []string{}, []string{"2"}, []string{}},
		[]string{"P,2", "A,2"},
	},
	{
		[][]string{[]string{"1", "2", "3"}, []string{}, []string{}, []string{"A", "B"}, []string{"X", "Y", "Z"}, []string{}},
		[]string{"1,A,X", "1,A,Y", "1,A,Z", "1,B,X", "1,B,Y", "1,B,Z", "2,A,X", "2,A,Y", "2,A,Z", "2,B,X", "2,B,Y", "2,B,Z", "3,A,X", "3,A,Y", "3,A,Z", "3,B,X", "3,B,Y", "3,B,Z"},
	},
	{
		[][]string{[]string{}, []string{}, []string{"A"}},
		[]string{"A"},
	},
	{
		[][]string{[]string{}, []string{"B"}, []string{}, []string{"A"}},
		[]string{"B,A"},
	},
	{
		[][]string{},
		[]string{""},
	},
	{
		[][]string{[]string{}},
		[]string{""},
	},
}

func TestGetCartesianProduct(t *testing.T){
	for i, test := range tests {
		res := GetCartesianProduct(test.input)
		for x, r := range res {
			if test.expect[x] != strings.Join(r, ",") {
				errStr := fmt.Sprintf("Expected %s but got %s in test case %d (permutation %d)", test.expect[x], strings.Join(r, ","), i, (x+1))
				t.Error(errStr)
			}
		}
	}
}