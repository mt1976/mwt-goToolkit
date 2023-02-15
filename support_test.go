package main

import (
	"testing"

	_ "github.com/denisenkom/go-mssqldb"
)

func Test_wrapVariable(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Test 1", args{"in"}, "{{.in}}"},
		{"Test 2", args{"out"}, "{{.out}}"},
		{"Test 3", args{""}, "{{.}}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wrapVariable(tt.args.in); got != tt.want {
				t.Errorf("wrapVariable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_enquote(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Test 1", args{"in"}, "\"in\""},
		{"Test 2", args{"out"}, "\"out\""},
		{"Test 3", args{""}, "\"\""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := enquote(tt.args.in); got != tt.want {
				t.Errorf("enquote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wrapTemplate(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Test 1", args{"in"}, "{{template \"in\" .}}"},
		{"Test 2", args{"out"}, "{{template \"out\" .}}"},
		{"Test 3", args{""}, "{{template \"\" .}}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wrapTemplate(tt.args.in); got != tt.want {
				t.Errorf("wrapTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tf(t *testing.T) {
	type args struct {
		in bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Test 1", args{true}, "Y"},
		{"Test 2", args{false}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tf(tt.args.in); got != tt.want {
				t.Errorf("tf() = %v, want %v", got, tt.want)
			}
		})
	}
}
