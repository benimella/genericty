package entity

import (
	"fmt"
	"testing"
)

// go test -v .\entity\ -run TestEntNonGeneric
func TestEntNonGeneric(t *testing.T) {
	fmt.Printf("user1:%+v user2:%+v \n", NewUser1(100), NewUser2("100"))
}

// go test -v .\entity\ -run TestEntGeneric
func TestEntGeneric(t *testing.T) {
	fmt.Printf("user1:%+v user2:%+v \n", NewUser(100), NewUser("100"))
}

// go test -v .\entity\ -run TestGenericX1
func TestGenericX1(t *testing.T) {
	name := "pod1"
	kind := "pod"
	p1 := GetObject[*Pod](name, kind)
	fmt.Printf("p1 => %+v\n", p1)

	name = "deployment1"
	kind = "deployment"
	d1 := GetObject[*Deployment](name, kind)
	fmt.Printf("d1 => %+v\n", d1)
}
