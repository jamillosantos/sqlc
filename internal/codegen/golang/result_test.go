package golang

import (
	"reflect"
	"testing"

	"github.com/sqlc-dev/sqlc/internal/metadata"
	"github.com/sqlc-dev/sqlc/internal/plugin"
)

func TestPutOutColumns_ForZeroColumns(t *testing.T) {
	tests := []struct {
		cmd  string
		want bool
	}{
		{
			cmd:  metadata.CmdExec,
			want: false,
		},
		{
			cmd:  metadata.CmdExecResult,
			want: false,
		},
		{
			cmd:  metadata.CmdExecRows,
			want: false,
		},
		{
			cmd:  metadata.CmdExecLastId,
			want: false,
		},
		{
			cmd:  metadata.CmdMany,
			want: true,
		},
		{
			cmd:  metadata.CmdOne,
			want: true,
		},
		{
			cmd:  metadata.CmdCopyFrom,
			want: false,
		},
		{
			cmd:  metadata.CmdBatchExec,
			want: false,
		},
		{
			cmd:  metadata.CmdBatchMany,
			want: true,
		},
		{
			cmd:  metadata.CmdBatchOne,
			want: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.cmd, func(t *testing.T) {
			query := &plugin.Query{
				Cmd:     tc.cmd,
				Columns: []*plugin.Column{},
			}
			got := putOutColumns(query)
			if got != tc.want {
				t.Errorf("putOutColumns failed. want %v, got %v", tc.want, got)
			}
		})
	}
}

func TestPutOutColumns_AlwaysTrueWhenQueryHasColumns(t *testing.T) {
	query := &plugin.Query{
		Cmd:     metadata.CmdMany,
		Columns: []*plugin.Column{{}},
	}
	if putOutColumns(query) != true {
		t.Error("should be true when we have columns")
	}
}

func Test_filterStructs(t *testing.T) {
	s1 := Struct{
		Name: "User",
	}
	s2 := Struct{
		Name: "User2",
	}
	s3 := Struct{
		Name: "Order",
	}

	tests := []struct {
		name    string
		filter  []string
		structs []Struct
		want    []Struct
	}{
		{
			"empty filter should return all structs",
			[]string{},
			[]Struct{s1, s2, s3},
			[]Struct{s1, s2, s3},
		},
		{
			"remove a specific struct",
			[]string{"-" + s1.Name},
			[]Struct{s1, s2, s3},
			[]Struct{s2, s3},
		},
		{
			"remove all and add a specific struct",
			[]string{"-*", s1.Name},
			[]Struct{s1, s2, s3},
			[]Struct{s1},
		},
		{
			"remove all and add a specific struct with +",
			[]string{"-*", "+" + s1.Name},
			[]Struct{s1, s2, s3},
			[]Struct{s1},
		},
		{
			"remove all and add a specific struct",
			[]string{"+*", "-" + s1.Name},
			[]Struct{s1, s2, s3},
			[]Struct{s2, s3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterStructs(tt.filter, tt.structs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterStructs() = %v, want %v", got, tt.want)
			}
		})
	}
}
