package main

import "testing"

func Test_trap(t *testing.T) {

	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "用例1",
			args: args{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
			want: 6,
		},
		{name: "用例2",
			args: args{height: []int{4, 2, 0, 3, 2, 5}},
			want: 9,
		},
		{name: "用例3",
			args: args{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
			want: 6,
		}, {
			name: "用例4",
			args: args{height: []int{4, 2, 0, 3, 2, 5}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trap(tt.args.height)
			t.Logf("trap() = %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})

	}

}
