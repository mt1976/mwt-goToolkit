package main

import (
	"testing"

	_ "github.com/denisenkom/go-mssqldb"
)

func Test_isAudit(t *testing.T) {
	type args struct {
		fn string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"Test 1", args{"in"}, false},
		{"Test 2", args{"out"}, false},
		{"Test 3", args{""}, false},
		{"Test 4", args{"_audit"}, true},
		{"Test 5", args{"audit_sql"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAudit(tt.args.fn); got != tt.want {
				t.Errorf("isAudit() = %v, want %v", got, tt.want)
			}
		})
	}
}
