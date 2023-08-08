package exercise

import "testing"

func Test_LongestPalindrome(t *testing.T) {

	type test struct {
		data   string
		result []string
	}
	tests := []test{
		{"babad", []string{"bab", "aba"}},
		{"adam", []string{"ada"}},
		{"a", []string{"a"}},
		{"cbbd", []string{"bb"}},
		{"bb", []string{"bb"}},
		{"abb", []string{"bb"}},
		{"ccc", []string{"ccc"}},
		{"aacabdkacaa", []string{"aca"}},
		{"abbcccbbbcaaccbababcbcabca", []string{"bbcccbb"}},
	}

	for _, v := range tests {
		x := LongestPalindrome(v.data)
		isSucceed := false
		for _, v := range v.result {
			if v == x {
				isSucceed = true
			}
		}
		if !isSucceed {
			t.Error("Expected", v.result, "Got", x)
		}
	}
}
