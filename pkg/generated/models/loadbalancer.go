package models

// Loadbalancer

import "encoding/json"

// Loadbalancer
type Loadbalancer struct {
	LoadbalancerProvider   string            `json:"loadbalancer_provider"`
	UUID                   string            `json:"uuid"`
	ParentUUID             string            `json:"parent_uuid"`
	IDPerms                *IdPermsType      `json:"id_perms"`
	LoadbalancerProperties *LoadbalancerType `json:"loadbalancer_properties"`
	Annotations            *KeyValuePairs    `json:"annotations"`
	Perms2                 *PermType2        `json:"perms2"`
	ParentType             string            `json:"parent_type"`
	FQName                 []string          `json:"fq_name"`
	DisplayName            string            `json:"display_name"`

	ServiceInstanceRefs         []*LoadbalancerServiceInstanceRef         `json:"service_instance_refs"`
	ServiceApplianceSetRefs     []*LoadbalancerServiceApplianceSetRef     `json:"service_appliance_set_refs"`
	VirtualMachineInterfaceRefs []*LoadbalancerVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs"`
}

// LoadbalancerServiceInstanceRef references each other
type LoadbalancerServiceInstanceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LoadbalancerServiceApplianceSetRef references each other
type LoadbalancerServiceApplianceSetRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LoadbalancerVirtualMachineInterfaceRef references each other
type LoadbalancerVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *Loadbalancer) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLoadbalancer makes Loadbalancer
func MakeLoadbalancer() *Loadbalancer {
	return &Loadbalancer{
		//TODO(nati): Apply default
		LoadbalancerProperties: MakeLoadbalancerType(),
		Annotations:            MakeKeyValuePairs(),
		Perms2:                 MakePermType2(),
		ParentType:             "",
		FQName:                 []string{},
		DisplayName:            "",
		LoadbalancerProvider:   "",
		UUID:                   "",
		ParentUUID:             "",
		IDPerms:                MakeIdPermsType(),
	}
}

// InterfaceToLoadbalancer makes Loadbalancer from interface
func InterfaceToLoadbalancer(iData interface{}) *Loadbalancer {
	data := iData.(map[string]interface{})
	return &Loadbalancer{
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		LoadbalancerProperties: InterfaceToLoadbalancerType(data["loadbalancer_properties"]),

		//{"description":"Loadbalancer configuration like  admin state, VIP, VIP subnet etc.","type":"object","properties":{"admin_state":{"type":"boolean"},"operating_status":{"type":"string"},"provisioning_status":{"type":"string"},"status":{"type":"string"},"vip_address":{"type":"string"},"vip_subnet_id":{"type":"string"}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		LoadbalancerProvider: data["loadbalancer_provider"].(string),

		//{"description":"Provider field selects backend provider of the LBaaS, Cloudadmin could offer different levels of service like gold, silver, bronze. Provided by  HA-proxy or various HW or SW appliances in the backend.","type":"string"}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}

	}
}

// InterfaceToLoadbalancerSlice makes a slice of Loadbalancer from interface
func InterfaceToLoadbalancerSlice(data interface{}) []*Loadbalancer {
	list := data.([]interface{})
	result := MakeLoadbalancerSlice()
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancer(item))
	}
	return result
}

// MakeLoadbalancerSlice() makes a slice of Loadbalancer
func MakeLoadbalancerSlice() []*Loadbalancer {
	return []*Loadbalancer{}
}
