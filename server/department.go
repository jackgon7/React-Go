package department

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// emp "server"

type DepartmentList struct {
	Departments map[string] *Department
}


type Employee struct {
	Name string `json:"Name"`
	Id	json.Number `json:"ID"`
	Department string `json:"Department"`
}

type Department struct {
	Employee []Employee
}



func (dl *DepartmentList) AddEmployee(w http.ResponseWriter, r *http.Request){

	var data Employee
	// var dl DepartmentList


	
	
	v := json.NewDecoder(r.Body).Decode(&data)
	// fmt.Println("v: ", v, "data", data)
	if v != nil {
		return
	}
	
	res := dl.addToDeptList(data)
	
	s := json.NewEncoder(w).Encode(res)
	// dl := &DepartmentList{}
	
	// fmt.Println(res)

	if s != nil {
		return
	}
	// fmt.Println(data.Name, "\t", data.Id, "\t", data.Department)
}

func (dl *DepartmentList)addToDeptList(emp Employee) *DepartmentList{
	
	if dl.Departments == nil {
		dl.Departments = make(map[string]*Department)
		
	}

	if dl.Departments[emp.Department] == nil {
		dl.Departments[emp.Department] = &Department{}
	}
	
	dl.Departments[emp.Department].Employee = append(dl.Departments[emp.Department].Employee, emp)

	fmt.Println(dl.Departments[emp.Department])
	// fmt.Println(len(dl.Departments), "\t", emp.Department, "\t", d.Employee)

	return dl
	
}