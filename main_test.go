package main

import (
	"testing"
)

func Test_swap(t *testing.T) {
	type args struct {
		s string
		i int
		j int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "ABC => BAC",
			args: args{
				s: "ABC",
				i: 0,
				j: 1,
			},
			want: "BAC",
		},
		// TODO: Add test cases.
		{
			name: "ABC => CBA",
			args: args{
				s: "ABC",
				i: 0,
				j: 2,
			},
			want: "CBA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swap(tt.args.s, tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("swap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_permutations(t *testing.T) {
	var strOut string
	type args struct {
		str string
		i   int
		n   int
		out *string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "ABC",
			args: args{
				str: "ABC",
				i:   0,
				n:   3,
				out: &strOut,
			},
			want: "ABCACBBACBCACBACAB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			permutations(tt.args.str, tt.args.i, tt.args.n, tt.args.out)
			if *tt.args.out != tt.want {
				t.Errorf("permutations() = %s, want %s", *tt.args.out, tt.want)
			}
		})
	}
}

func Test_createDNAPattern(t *testing.T) {
	var p string
	type args struct {
		dna        []string
		dnaPattern *string
	}
	tests := []struct {
		name       string
		args       args
		wantLength int
	}{
		// TODO: Add test cases.
		{
			name: "ABCD",
			args: args{
				dna:        []string{`A`, `B`, `C`, `D`},
				dnaPattern: &p,
			},
			wantLength: 384,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createDNAPattern(tt.args.dna, tt.args.dnaPattern)
			if len(*tt.args.dnaPattern) != tt.wantLength {
				t.Errorf("permutations() = %d, want length %d", len(*tt.args.dnaPattern), tt.wantLength)
			}
		})
	}
}

func Test_createFindKey(t *testing.T) {
	outArr := []string{}
	outPattern := ""

	type args struct {
		outPattern *string
		dnaPattern string
		arr        *[]string
		dnaArray   []string
	}
	tests := []struct {
		name       string
		args       args
		depenFunc  func(dna []string, dnaPattern *string)
		wantLength int
	}{
		// TODO: Add test cases.
		{
			name: "ABCD",
			args: args{
				dnaPattern: "ABCD",
				arr:        &outArr,
				dnaArray:   []string{`A`, `B`, `C`, `D`},
				outPattern: &outPattern,
			},
			depenFunc:  createDNAPattern,
			wantLength: 2992,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.depenFunc(tt.args.dnaArray, tt.args.outPattern)
			got := createFindKey(*tt.args.outPattern)
			if tt.wantLength != len(*got) {
				t.Errorf("createFindKey() = %d, want length %d", len(*got), tt.wantLength)
			}
		})
	}
}

func Test_distinct(t *testing.T) {
	outArr := []string{}
	outPattern := ""

	type args struct {
		outPattern *string
		dnaPattern string
		arr        *[]string
		dnaArray   []string
	}
	tests := []struct {
		name       string
		args       args
		depenFunc  func(dna []string, dnaPattern *string)
		wantLength int
	}{
		// TODO: Add test cases.
		{
			name: "ABCD",
			args: args{
				dnaPattern: "ABCD",
				arr:        &outArr,
				dnaArray:   []string{`A`, `B`, `C`, `D`},
				outPattern: &outPattern,
			},
			depenFunc:  createDNAPattern,
			wantLength: 1403,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.depenFunc(tt.args.dnaArray, tt.args.outPattern)
			got := createFindKey(*tt.args.outPattern)
			distinct(got)
			if tt.wantLength != len(*got) {
				t.Errorf("createFindKey() = %d, want length %d", len(*got), tt.wantLength)
			}
		})
	}
}

func Test_counter(t *testing.T) {
	outArr := []string{}
	outPattern := ""

	type args struct {
		outPattern *string
		dnaPattern string
		arr        *[]string
		dnaArray   []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "ABCD",
			args: args{
				dnaPattern: "ABCD",
				arr:        &outArr,
				dnaArray:   []string{`A`, `B`, `C`, `D`},
				outPattern: &outPattern,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			createDNAPattern(tt.args.dnaArray, tt.args.outPattern)
			got := createFindKey(*tt.args.outPattern)
			distinct(got)

			gotMap := counter(*tt.args.outPattern, got)
			if len(*got) != len(gotMap) {
				t.Errorf("counter() = %d, want length %d", len(*got), len(gotMap))
			}

		})
	}
}
