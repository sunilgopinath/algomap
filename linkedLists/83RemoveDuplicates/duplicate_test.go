package removeduplicates_test

import (
	"reflect"
	"testing"

	removeduplicates "github.com/sunilgopinath/algomap/linkedLists/83RemoveDuplicates"
)

func TestDeleteDuplicates(t *testing.T) {

	type test struct {
		input *removeduplicates.ListNode
		want  *removeduplicates.ListNode
	}

	tests := []test{
		{
			input: &removeduplicates.ListNode{Val: 1, Next: &removeduplicates.ListNode{Val: 1, Next: &removeduplicates.ListNode{Val: 2, Next: &removeduplicates.ListNode{Val: 3, Next: &removeduplicates.ListNode{Val: 3, Next: nil}}}}},
			want:  &removeduplicates.ListNode{Val: 1, Next: &removeduplicates.ListNode{Val: 2, Next: &removeduplicates.ListNode{Val: 3, Next: nil}}},
		},
		{
			input: &removeduplicates.ListNode{Val: 1, Next: &removeduplicates.ListNode{Val: 1, Next: &removeduplicates.ListNode{Val: 1, Next: nil}}},
			want:  &removeduplicates.ListNode{Val: 1, Next: nil},
		},
		{
			input: &removeduplicates.ListNode{Val: 1, Next: &removeduplicates.ListNode{Val: 2, Next: &removeduplicates.ListNode{Val: 3, Next: nil}}},
			want:  &removeduplicates.ListNode{Val: 1, Next: &removeduplicates.ListNode{Val: 2, Next: &removeduplicates.ListNode{Val: 3, Next: nil}}},
		},
	}
	
	for _, tc := range tests {
		got := removeduplicates.DeleteDuplicates(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("deleteDuplicates(%v) = %v; want %v", tc.input, got, tc.want)
		}
	}
}