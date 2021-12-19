package employeeService

import (
	"goa-golang/app/repository/employeeRepository"
	"reflect"
	"testing"
)

func TestNewEmployeeService(t *testing.T) {
	type args struct {
		employeeRepository employeeRepository.EmployeeRepositoryInterface
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
