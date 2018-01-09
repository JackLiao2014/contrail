package models

// FloatingIPPool

import "encoding/json"

// FloatingIPPool
type FloatingIPPool struct {
	FQName                []string                  `json:"fq_name"`
	IDPerms               *IdPermsType              `json:"id_perms"`
	Annotations           *KeyValuePairs            `json:"annotations"`
	UUID                  string                    `json:"uuid"`
	ParentUUID            string                    `json:"parent_uuid"`
	FloatingIPPoolSubnets *FloatingIpPoolSubnetType `json:"floating_ip_pool_subnets"`
	DisplayName           string                    `json:"display_name"`
	Perms2                *PermType2                `json:"perms2"`
	ParentType            string                    `json:"parent_type"`

	FloatingIPs []*FloatingIP `json:"floating_ips"`
}

// String returns json representation of the object
func (model *FloatingIPPool) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeFloatingIPPool makes FloatingIPPool
func MakeFloatingIPPool() *FloatingIPPool {
	return &FloatingIPPool{
		//TODO(nati): Apply default
		FloatingIPPoolSubnets: MakeFloatingIpPoolSubnetType(),
		FQName:                []string{},
		IDPerms:               MakeIdPermsType(),
		Annotations:           MakeKeyValuePairs(),
		UUID:                  "",
		ParentUUID:            "",
		ParentType:            "",
		DisplayName:           "",
		Perms2:                MakePermType2(),
	}
}

// InterfaceToFloatingIPPool makes FloatingIPPool from interface
func InterfaceToFloatingIPPool(iData interface{}) *FloatingIPPool {
	data := iData.(map[string]interface{})
	return &FloatingIPPool{
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		FloatingIPPoolSubnets: InterfaceToFloatingIpPoolSubnetType(data["floating_ip_pool_subnets"]),

		//{"description":"Subnets that restrict floating ip allocation from the corresponding virtual network.","type":"object","properties":{"subnet_uuid":{"type":"array","item":{"type":"string"}}}}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}

	}
}

// InterfaceToFloatingIPPoolSlice makes a slice of FloatingIPPool from interface
func InterfaceToFloatingIPPoolSlice(data interface{}) []*FloatingIPPool {
	list := data.([]interface{})
	result := MakeFloatingIPPoolSlice()
	for _, item := range list {
		result = append(result, InterfaceToFloatingIPPool(item))
	}
	return result
}

// MakeFloatingIPPoolSlice() makes a slice of FloatingIPPool
func MakeFloatingIPPoolSlice() []*FloatingIPPool {
	return []*FloatingIPPool{}
}
