package main

import "testing"

var tests = map[string]string{
	"0":                     "0",
	"+1":                    "+1",
	"-1":                    "-1",
	"+1.23":                 "+1.23",
	"-1.233":                "-1.233",
	"1.233456":              "1.233,456",
	"-1111.233456":          "-1,111.233,456",
	"-5431111.233456":       "-5,431,111.233,456",
}

var performanceTest = "5723984759232794823984793284"

func TestComma(t *testing.T) {
	for input, expected := range tests {
		got := comma(input)
		if got != expected {
			t.Errorf("expected: %q, got: %q\n", expected, got)
		}
	}
}

func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		comma(performanceTest)
	}
}
