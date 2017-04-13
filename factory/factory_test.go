package factory

import (
	"reflect"
	"testing"
)

func TestRegisterService(t *testing.T) {
	type args struct {
		path    string
		service Factory
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterService(tt.args.path, tt.args.service)
		})
	}
}

func TestGetServices(t *testing.T) {
	tests := []struct {
		name string
		want []Factory
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetServices(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetServices() = %v, want %v", got, tt.want)
			}
		})
	}
}
