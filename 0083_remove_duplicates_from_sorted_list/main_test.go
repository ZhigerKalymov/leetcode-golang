package _083_remove_duplicates_from_sorted_list

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}

	var tests = []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test_1.1.2",
			args{&ListNode{1, &ListNode{1, &ListNode{2, nil}}}},
			&ListNode{1, &ListNode{2, nil}},
		},
		{
			"test_1.1.1",
			args{&ListNode{1, &ListNode{1, &ListNode{1, nil}}}},
			&ListNode{1, nil},
		},
		{
			"test_1",
			args{&ListNode{1, nil}},
			&ListNode{1, nil},
		},
		{
			"test_1.1.2.3.3",
			args{&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}},
			&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
