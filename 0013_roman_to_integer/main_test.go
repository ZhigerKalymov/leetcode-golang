package _013_roman_to_integer

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"III",
			args{"III"},
			3,
		},
		{
			"IV",
			args{"IV"},
			4,
		},
		{
			"IX",
			args{"IX"},
			9,
		},
		{
			"LVIII",
			args{"LVIII"},
			58,
		},
		{
			"MCMXCIV",
			args{"MCMXCIV"},
			1994,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}


func benchmarkRomanToInt(input string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		romanToInt(input)
	}
}

func BenchmarkRomanToInt3(b *testing.B) { benchmarkRomanToInt("III", b) }
func BenchmarkRomanToInt9(b *testing.B) { benchmarkRomanToInt("IX", b) }
func BenchmarkRomanToInt58(b *testing.B) { benchmarkRomanToInt("LVIII", b) }
func BenchmarkRomanToInt1994(b *testing.B) { benchmarkRomanToInt("MCMXCIV", b) }