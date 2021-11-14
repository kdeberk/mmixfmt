package main

import (
	"reflect"
	"testing"
)

func TestReadFragments(t *testing.T) {
	tts := []struct {
		name  string
		line  string
		frags lineFragments
	}{
		{
			name:  "command with single arg",
			line:  "LOC #100",
			frags: lineFragments{cols: []string{"LOC", "#100"}},
		},
		{
			name:  "command with single arg and comment",
			line:  "LOC #100 This is a comment",
			frags: lineFragments{cols: []string{"LOC", "#100"}, comment: "This is a comment"},
		},
		{
			name:  "labeled command with multiple args",
			line:  "Loop LDO xk,x0,kk",
			frags: lineFragments{cols: []string{"Loop", "LDO", "xk,x0,kk"}},
		},
		{
			name:  "labeled command with multiple args and comment",
			line:  "Loop LDO xk,x0,kk xk ← X[k]",
			frags: lineFragments{cols: []string{"Loop", "LDO", "xk,x0,kk"}, comment: "xk ← X[k]"},
		},
		{
			name:  "labeled command with string literal as arg",
			line:  `String  BYTE  ", world",#a,0`,
			frags: lineFragments{cols: []string{"String", "BYTE", `", world",#a,0`}},
		},
		{
			name:  "labeled command with string literal as arg and comment",
			line:  `String  BYTE  ", world",#a,0 This is another comment`,
			frags: lineFragments{cols: []string{"String", "BYTE", `", world",#a,0`}, comment: "This is another comment"},
		},
		{
			name:  "line containing only a comment",
			line:  " %  This is a comment",
			frags: lineFragments{cols: nil, comment: "%  This is a comment"},
		},
	}

	for _, tt := range tts {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := readFragments(tt.line)
			if !reflect.DeepEqual(got, tt.frags) {
				t.Fatalf("wrong output:\n\tgot:  %#v\n\twant: %#v", got, tt.frags)
			}
		})
	}
}
