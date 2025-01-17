package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				x: 1,
				y: 2,
			},
			want: 3,
		},
		{
			name: "success",
			args: args{
				x: 10,
				y: 90,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				x: 2,
				y: 3,
			},
			want: 6,
		},
		{
			name: "success",
			args: args{
				x: 4,
				y: 4,
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiply(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrime(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success",
			args: args{
				4,
			},
			want: []int{2, 3, 5, 7},
		},
		{
			name: "success",
			args: args{
				9,
			},
			want: []int{2, 3, 5, 7, 11, 13, 17, 19, 23},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prime(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	type args struct {
		input int
	}
	var tests = []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success",
			args: args{
				4,
			},
			want: []int{0, 1, 1, 2},
		},
		{
			name: "success",
			args: args{
				10,
			},
			want: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
