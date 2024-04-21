package snippet

import "testing"

func Test_lcs(t *testing.T) {
	type args struct {
		w1 string
		w2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				w1: "abcde",
				w2: "ace",
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				w1: "abcde",
				w2: "abce",
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				w1: "abc",
				w2: "abc",
			},
			want: 3,
		},
		{
			name: "빈값",
			args: args{
				w1: "",
				w2: "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lcs(tt.args.w1, tt.args.w2); got != tt.want {
				t.Errorf("lcs() = %v, want %v", got, tt.want)
			}
		})
	}
}
