package simple

import (
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		got  []int
		want []string
	}{
		{[]int{6, 8, 9}, []string{"6", "8->9"}},
		{[]int{4, 5, 7}, []string{"4->5", "7"}},
		{[]int{}, []string{}},
		{[]int{4, 5, 6, 7}, []string{"4->7"}},
		{[]int{4}, []string{"4"}},
	}

	for i, test := range tests {
		if !reflect.DeepEqual(SummaryRanges(test.got), test.want) {
			t.Errorf("failed %v", i)
		}
	}
}
