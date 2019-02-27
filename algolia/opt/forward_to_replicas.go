// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type ForwardToReplicasOption struct {
	value bool
}

func ForwardToReplicas(v bool) ForwardToReplicasOption {
	return ForwardToReplicasOption{v}
}

func (o ForwardToReplicasOption) Get() bool {
	return o.value
}

func (o ForwardToReplicasOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *ForwardToReplicasOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = false
		return nil
	}
	return json.Unmarshal(data, &o.value)
}
