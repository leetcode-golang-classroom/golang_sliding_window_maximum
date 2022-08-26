package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	for idx := 0; idx < b.N; idx++ {
		maxSlidingWindow(nums, k)
	}
}
func Test_maxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums = [1,3,-1,-3,5,3,6,7], k = 3",
			args: args{nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3},
			want: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name: "nums = [1], k = 1",
			args: args{nums: []int{1}, k: 1},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
