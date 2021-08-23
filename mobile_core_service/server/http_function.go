package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rebelITT/mobile_core_service/config"
	"github.com/rebelITT/mobile_core_service/proto"
	"google.golang.org/grpc"
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type UserList struct {
	UserInfo []User `json:"users"`
}

type Employee struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type EmployeeList struct {
	EmployeeInfo []Employee `json:"employees"`
}

func GetUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Context-Type", "application/json")

		fmt.Println("starting grpc user client application")
		conn, err := grpc.Dial(config.GetUserServiceGRPCPort(), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		client := proto.NewUserServiceClient(conn)

		var emp empty.Empty
		response, err := client.GetUsers(r.Context(), &emp)
		if err != nil {
			fmt.Printf("error in sending grpc request. error: %s", err.Error())
			return
		}

		var userDet User
		var userListResponse UserList
		for _, userInfo := range response.Users {
			userDet = User{
				ID:   userInfo.ID,
				Name: userInfo.Name,
			}
			userListResponse.UserInfo = append(userListResponse.UserInfo, userDet)
		}
		respBytes, err := json.Marshal(userListResponse)
		if err != nil {
			fmt.Printf("error in marshal.error: %s", err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(respBytes)
	}
}

func GetEmployees() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Context-Type", "application/json")

		fmt.Println("starting grpc employee client application")
		conn, err := grpc.Dial(config.GetEmployeeServiceGRPCPort(), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		client := proto.NewEmployeeServiceClient(conn)

		var emp empty.Empty
		response, err := client.GetEmployees(r.Context(), &emp)
		if err != nil {
			fmt.Printf("error in sending grpc request. error: %s", err.Error())
			return
		}

		var employeeDet Employee
		var employeeListResponse EmployeeList
		for _, employeeInfo := range response.Employees {
			employeeDet = Employee{
				ID:   employeeInfo.ID,
				Name: employeeInfo.Name,
			}
			employeeListResponse.EmployeeInfo = append(employeeListResponse.EmployeeInfo, employeeDet)
		}
		respBytes, err := json.Marshal(employeeListResponse)
		if err != nil {
			fmt.Printf("error in marshal.error: %s", err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(respBytes)
	}
}
