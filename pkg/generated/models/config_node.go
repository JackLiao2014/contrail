package models

// ConfigNode

import "encoding/json"

// ConfigNode
type ConfigNode struct {
	ConfigNodeIPAddress IpAddressType  `json:"config_node_ip_address"`
	ParentUUID          string         `json:"parent_uuid"`
	ParentType          string         `json:"parent_type"`
	Annotations         *KeyValuePairs `json:"annotations"`
	UUID                string         `json:"uuid"`
	FQName              []string       `json:"fq_name"`
	IDPerms             *IdPermsType   `json:"id_perms"`
	DisplayName         string         `json:"display_name"`
	Perms2              *PermType2     `json:"perms2"`
}

// String returns json representation of the object
func (model *ConfigNode) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeConfigNode makes ConfigNode
func MakeConfigNode() *ConfigNode {
	return &ConfigNode{
		//TODO(nati): Apply default
		FQName:              []string{},
		IDPerms:             MakeIdPermsType(),
		DisplayName:         "",
		Perms2:              MakePermType2(),
		ConfigNodeIPAddress: MakeIpAddressType(),
		ParentUUID:          "",
		ParentType:          "",
		Annotations:         MakeKeyValuePairs(),
		UUID:                "",
	}
}

// InterfaceToConfigNode makes ConfigNode from interface
func InterfaceToConfigNode(iData interface{}) *ConfigNode {
	data := iData.(map[string]interface{})
	return &ConfigNode{
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ConfigNodeIPAddress: InterfaceToIpAddressType(data["config_node_ip_address"]),

		//{"description":"Ip address of the config node, set while provisioning.","type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}

	}
}

// InterfaceToConfigNodeSlice makes a slice of ConfigNode from interface
func InterfaceToConfigNodeSlice(data interface{}) []*ConfigNode {
	list := data.([]interface{})
	result := MakeConfigNodeSlice()
	for _, item := range list {
		result = append(result, InterfaceToConfigNode(item))
	}
	return result
}

// MakeConfigNodeSlice() makes a slice of ConfigNode
func MakeConfigNodeSlice() []*ConfigNode {
	return []*ConfigNode{}
}
