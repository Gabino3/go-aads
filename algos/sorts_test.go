package algos

import (
	"reflect"
	"testing"
	"time"
)

func Test_MergeSort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				a: []int{5, 3, 1, 4, 2},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "test2",
			args: args{
				a: []int{5, 1, 1, 4, 2},
			},
			want: []int{1, 1, 2, 4, 5},
		},
		{
			name: "test3",
			args: args{
				a: []int{5, 3, 1, 4, 2},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "test4",
			args: args{
				a: []int{5, 3, 1, 4, 2},
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ConcurrentMergeSort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				a: []int{5, 3, 1, 4, 2},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "test2",
			args: args{
				a: []int{5, 1, 1, 4, 2},
			},
			want: []int{1, 1, 2, 4, 5},
		},
		{
			name: "test3",
			args: args{
				a: []int{5, 3, 1, 4, 2},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "test4",
			args: args{
				a: []int{5, 3, 1, 4, 2},
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConcurrentMergeSort(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timed(t *testing.T) {

	x := make([][]int, 1000)

	for i := range x {
		x[i] = GenerateRandomIntSlice(500)
	}

	start := time.Now()
	for _, a := range x {
		ConcurrentMergeSort(a)
	}
	elapsed := time.Since(start)
	start = time.Now()
	for _, a := range x {
		MergeSort(a)
	}
	otherElapsed := time.Since(start)

	t.Logf("\nConcurrent time: %s\nReg time: %s", elapsed, otherElapsed)
	//turns out concurrent is slower... by a lot
}
