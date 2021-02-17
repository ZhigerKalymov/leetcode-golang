package _700_search_in_a_binary_search_tree

import (
	"reflect"
	"testing"
)

func Test_searchBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"[4,2,7,1,3],2 = [2,1,3]",
			args{
				&TreeNode{
					4,
					&TreeNode{
						2,
						&TreeNode{
							1,
							nil,
							nil},
						&TreeNode{
							3,
							nil,
							nil},
					},
					&TreeNode{
						7,
						nil,
						nil},
				},
				2,
			},
			&TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}},
		},
		{
			"[4,2,7,1,3],5 = []",
			args{
				&TreeNode{
					4,
					&TreeNode{
						2,
						&TreeNode{
							1,
							nil,
							nil},
						&TreeNode{
							3,
							nil,
							nil},
					},
					&TreeNode{
						7,
						nil,
						nil},
				},
				5,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
