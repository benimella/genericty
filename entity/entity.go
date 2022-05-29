package entity

type User1 struct {
	Id   int
	Name string
}

func NewUser1(id int) *User1 {
	return &User1{Id: id}
}

type User2 struct {
	Id   string
	Name string
	Age  int
}

func NewUser2(id string) *User2 {
	return &User2{Id: id}
}

type UserIdKind interface {
	int | string
}

type User[T UserIdKind] struct {
	Id   T
	Name string
}

/*
type User[T any] struct {
	Id   T
	Name string
}
*/

func NewUser[T UserIdKind](id T) *User[T] {
	return &User[T]{Id: id}
}

type Container struct {
	Name string
	Cmd  string
	Args []string
}

type Pod struct {
	PodUUID        string       `json:"pod_uuid"`
	Name           string       `json:"name"`
	Kind           string       `json:"kind"`
	OwnerReference []string     `json:"owner_reference"`
	InitContainers []*Container `json:"init_containers"`
	Containers     []*Container `json:"containers"`
}

func (o *Pod) GetName() string {
	return o.Name
}

type DaemoneSet struct {
	DaemoneSetUUID string       `json:"daemoneset_uuid"`
	Name           string       `json:"name"`
	Kind           string       `json:"kind"`
	OwnerReference []string     `json:"owner_reference"`
	InitContainers []*Container `json:"init_containers"`
	Containers     []*Container `json:"containers"`
}

func (o *DaemoneSet) GetName() string {
	return o.Name
}

type Deployment struct {
	DeploymentUUID string       `json:"deployment_uuid"`
	Name           string       `json:"name"`
	Kind           string       `json:"kind"`
	OwnerReference []string     `json:"owner_reference"`
	InitContainers []*Container `json:"init_containers"`
	Containers     []*Container `json:"containers"`
}

func (o *Deployment) GetName() string {
	return o.Name
}
