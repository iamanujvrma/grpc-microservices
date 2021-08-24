package main

import (
	"context"
	"fmt"
	"net"

	"github.com/Fatema-Moaiyadi/rebelITT-common/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rebelITT-service/employee_service/config"
	"google.golang.org/grpc"
)

type employeeService struct {
}

type Employee struct {
	ID   int64
	Name string
}

type EmployeeList struct {
	EmployeeInfo []Employee
}

func (s *employeeService) GetEmployees(ctx context.Context, emp *empty.Empty) (*proto.EmployeesListResponse, error) {
	employeesInfo := EmployeeList{
		EmployeeInfo: []Employee{
			{
				ID:   1,
				Name: "test employee1",
			},
			{
				ID:   2,
				Name: "test employee2",
			},
		},
	}

	var employees []*proto.Employees
	var employee *proto.Employees
	for _, e := range employeesInfo.EmployeeInfo {
		employee = &proto.Employees{
			ID:   e.ID,
			Name: e.Name,
		}
		employees = append(employees, employee)
	}

	response := &proto.EmployeesListResponse{
		Employees: employees,
	}
	return response, nil
}

func main() {
	err := config.Init()
	if err != nil {
		fmt.Printf("error in init. error: %s", err.Error())
	}

	addr := config.GetGRPCPort()
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("error in listening at grpc port. error: %s", err.Error())
	}

	srv := grpc.NewServer()

	proto.RegisterEmployeeServiceServer(srv, &employeeService{})
	srv.Serve(lis)
}
