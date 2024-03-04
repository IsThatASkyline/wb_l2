package main

import "testing"

func TestUnpuck(t *testing.T) {
	testMap := map[string]string{
		"a4bc2d5e": "aaaabccddddde",
		"abcd":     "abcd",
		"45":       "",
		"":         "",
		`qwe\4\5`:  "qwe45",
		`qwe\45`:   "qwe44444",
		`qwe\\5`:   `qwe\\\\\`,
	}
	for k, v := range testMap {
		actual, _ := unpuck(k)
		if actual != v {
			t.Error("Excpected:", v, "Actual:", actual)
		}
	}
}
