package service

import (
	"dou-survey/app/repository"
	"reflect"
	"testing"
)

func TestNewEmployeeService(t *testing.T) {
	type args struct {
		employeeRepository repository.EmployeeRepositoryInterface
	}
	tests := []struct {
		name string
		args args
		want EmployeeServiceInterface
	}{
		{
			name: "success",
			args: args{
				employeeRepository: nil,
			},
			want: &EmployeeService{
				employeeRepo: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmployeeService(tt.args.employeeRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EmployeeInit() = %v, want %v", got, tt.want)
			}
		})
	}
}
