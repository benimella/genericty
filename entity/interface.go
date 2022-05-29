package entity

import (
	"encoding/json"
	"log"
)

type Object interface {
	GetName() string
}

func GetObject[T Object](name string, kind string) T {
	var ret T
	list := []map[string]interface{}{
		{"name": "pod1", "kind": "pod", "pod_uuid": "101"},
		{"name": "daemonset1", "kind": "daemonset", "daemonset_uuid": "1"},
		{"name": "deployment1", "kind": "deployment", "deployment_uuid": "1"},
		{"mame": "pod2", "kind": "pod", "pod_uuid": "102"},
		{"name": "pod1", "kind": "deployment", "deployment_uuid": "2"},
	}
	item := (map[string]interface{})(nil)
	for _, v := range list {
		if v["name"] == name && v["kind"] == kind {
			item = v
			break
		}
	}
	if item == nil {
		log.Println("没有找到对应资源")
		return ret
	}
	b, err := json.Marshal(item)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, &ret)
	if err != nil {
		log.Fatal(err)
	}
	return ret
}
