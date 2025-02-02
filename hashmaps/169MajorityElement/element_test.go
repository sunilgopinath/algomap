package majorityelement_test

import (
	"testing"

	majorityelement "github.com/sunilgopinath/algomap/hashmaps/169MajorityElement"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}
	for _, tt := range tests {
		if got := majorityelement.MajorityElement(tt.nums); got != tt.want {
			t.Errorf("MajorityElement(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}

func TestMajorityElementBoyersMoore(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}
	for _, tt := range tests {
		if got := majorityelement.MajorityElementBoyersMoore(tt.nums); got != tt.want {
			t.Errorf("MajorityElementBoyersMoore(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}