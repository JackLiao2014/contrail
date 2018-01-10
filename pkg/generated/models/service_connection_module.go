package models

// ServiceConnectionModule

import "encoding/json"

// ServiceConnectionModule
type ServiceConnectionModule struct {
	IDPerms     *IdPermsType          `json:"id_perms"`
	Perms2      *PermType2            `json:"perms2"`
	ParentUUID  string                `json:"parent_uuid"`
	ParentType  string                `json:"parent_type"`
	FQName      []string              `json:"fq_name"`
	ServiceType ServiceConnectionType `json:"service_type"`
	DisplayName string                `json:"display_name"`
	Annotations *KeyValuePairs        `json:"annotations"`
	UUID        string                `json:"uuid"`
	E2Service   E2servicetype         `json:"e2_service"`

	ServiceObjectRefs []*ServiceConnectionModuleServiceObjectRef `json:"service_object_refs"`
}

// ServiceConnectionModuleServiceObjectRef references each other
type ServiceConnectionModuleServiceObjectRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *ServiceConnectionModule) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceConnectionModule makes ServiceConnectionModule
func MakeServiceConnectionModule() *ServiceConnectionModule {
	return &ServiceConnectionModule{
		//TODO(nati): Apply default
		Annotations: MakeKeyValuePairs(),
		UUID:        "",
		E2Service:   MakeE2servicetype(),
		DisplayName: "",
		Perms2:      MakePermType2(),
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		ServiceType: MakeServiceConnectionType(),
		IDPerms:     MakeIdPermsType(),
	}
}

// InterfaceToServiceConnectionModule makes ServiceConnectionModule from interface
func InterfaceToServiceConnectionModule(iData interface{}) *ServiceConnectionModule {
	data := iData.(map[string]interface{})
	return &ServiceConnectionModule{
		ServiceType: InterfaceToServiceConnectionType(data["service_type"]),

		//{"description":"Type of service assigned for this object","type":"string","enum":["vpws-l2ckt","vpws-l2vpn","vpws-evpn","fabric-interface"]}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		E2Service: InterfaceToE2servicetype(data["e2_service"]),

		//{"description":"E2 service type.","type":"string","enum":["point-to-point","point-to-list","multi-point-to-multi-point"]}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}

	}
}

// InterfaceToServiceConnectionModuleSlice makes a slice of ServiceConnectionModule from interface
func InterfaceToServiceConnectionModuleSlice(data interface{}) []*ServiceConnectionModule {
	list := data.([]interface{})
	result := MakeServiceConnectionModuleSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceConnectionModule(item))
	}
	return result
}

// MakeServiceConnectionModuleSlice() makes a slice of ServiceConnectionModule
func MakeServiceConnectionModuleSlice() []*ServiceConnectionModule {
	return []*ServiceConnectionModule{}
}
