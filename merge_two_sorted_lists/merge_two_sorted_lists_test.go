package mergeTwoSortedLists

import (
	"reflect"
	"testing"

	"../utils"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *utils.ListNode
		l2 *utils.ListNode
	}
	tests := []struct {
		name string
		args args
		want *utils.ListNode
	}{
		// TODO: Add test cases.
		{
			name: "case 0",
			args: args{
				l1: utils.ConvertIntArrayToListNode([]int{1, 2, 4}),
				l2: utils.ConvertIntArrayToListNode([]int{1, 3, 4}),
			},
			want: utils.ConvertIntArrayToListNode([]int{1, 1, 2, 3, 4, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(utils.ConvertListNodeToIntArray(got), utils.ConvertListNodeToIntArray(tt.want)) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
