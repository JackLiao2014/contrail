package models

// ServiceAppliance

import "encoding/json"

// ServiceAppliance
type ServiceAppliance struct {
	IDPerms                         *IdPermsType     `json:"id_perms"`
	DisplayName                     string           `json:"display_name"`
	ServiceApplianceUserCredentials *UserCredentials `json:"service_appliance_user_credentials"`
	ServiceApplianceIPAddress       IpAddressType    `json:"service_appliance_ip_address"`
	UUID                            string           `json:"uuid"`
	ParentType                      string           `json:"parent_type"`
	FQName                          []string         `json:"fq_name"`
	ServiceApplianceProperties      *KeyValuePairs   `json:"service_appliance_properties"`
	Annotations                     *KeyValuePairs   `json:"annotations"`
	Perms2                          *PermType2       `json:"perms2"`
	ParentUUID                      string           `json:"parent_uuid"`

	PhysicalInterfaceRefs []*ServiceAppliancePhysicalInterfaceRef `json:"physical_interface_refs"`
}

// ServiceAppliancePhysicalInterfaceRef references each other
type ServiceAppliancePhysicalInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *ServiceApplianceInterfaceType
}

// String returns json representation of the object
func (model *ServiceAppliance) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceAppliance makes ServiceAppliance
func MakeServiceAppliance() *ServiceAppliance {
	return &ServiceAppliance{
		//TODO(nati): Apply default
		ServiceApplianceUserCredentials: MakeUserCredentials(),
		ServiceApplianceIPAddress:       MakeIpAddressType(),
		UUID:                       "",
		ParentType:                 "",
		FQName:                     []string{},
		IDPerms:                    MakeIdPermsType(),
		DisplayName:                "",
		ServiceApplianceProperties: MakeKeyValuePairs(),
		Annotations:                MakeKeyValuePairs(),
		Perms2:                     MakePermType2(),
		ParentUUID:                 "",
	}
}

// InterfaceToServiceAppliance makes ServiceAppliance from interface
func InterfaceToServiceAppliance(iData interface{}) *ServiceAppliance {
	data := iData.(map[string]interface{})
	return &ServiceAppliance{
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ServiceApplianceProperties: InterfaceToKeyValuePairs(data["service_appliance_properties"]),

		//{"description":"List of Key:Value pairs used by the provider driver of this service appliance.","type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		ServiceApplianceIPAddress: InterfaceToIpAddressType(data["service_appliance_ip_address"]),

		//{"description":"Management Ip address of the service-appliance.","type":"string"}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		ServiceApplianceUserCredentials: InterfaceToUserCredentials(data["service_appliance_user_credentials"]),

		//{"description":"Authentication credentials for driver to access service appliance.","type":"object","properties":{"password":{"type":"string"},"username":{"type":"string"}}}

	}
}

// InterfaceToServiceApplianceSlice makes a slice of ServiceAppliance from interface
func InterfaceToServiceApplianceSlice(data interface{}) []*ServiceAppliance {
	list := data.([]interface{})
	result := MakeServiceApplianceSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceAppliance(item))
	}
	return result
}

// MakeServiceApplianceSlice() makes a slice of ServiceAppliance
func MakeServiceApplianceSlice() []*ServiceAppliance {
	return []*ServiceAppliance{}
}
