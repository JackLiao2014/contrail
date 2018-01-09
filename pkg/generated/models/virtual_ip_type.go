package models

// VirtualIpType

import "encoding/json"

// VirtualIpType
type VirtualIpType struct {
	ProtocolPort          int                      `json:"protocol_port"`
	Status                string                   `json:"status"`
	StatusDescription     string                   `json:"status_description"`
	Protocol              LoadbalancerProtocolType `json:"protocol"`
	PersistenceCookieName string                   `json:"persistence_cookie_name"`
	PersistenceType       SessionPersistenceType   `json:"persistence_type"`
	AdminState            bool                     `json:"admin_state"`
	Address               IpAddressType            `json:"address"`
	SubnetID              UuidStringType           `json:"subnet_id"`
	ConnectionLimit       int                      `json:"connection_limit"`
}

// String returns json representation of the object
func (model *VirtualIpType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualIpType makes VirtualIpType
func MakeVirtualIpType() *VirtualIpType {
	return &VirtualIpType{
		//TODO(nati): Apply default
		AdminState:            false,
		Address:               MakeIpAddressType(),
		ProtocolPort:          0,
		Status:                "",
		StatusDescription:     "",
		Protocol:              MakeLoadbalancerProtocolType(),
		PersistenceCookieName: "",
		PersistenceType:       MakeSessionPersistenceType(),
		SubnetID:              MakeUuidStringType(),
		ConnectionLimit:       0,
	}
}

// InterfaceToVirtualIpType makes VirtualIpType from interface
func InterfaceToVirtualIpType(iData interface{}) *VirtualIpType {
	data := iData.(map[string]interface{})
	return &VirtualIpType{
		AdminState: data["admin_state"].(bool),

		//{"description":"Administrative up or down.","type":"boolean"}
		Address: InterfaceToIpAddressType(data["address"]),

		//{"description":"IP address automatically allocated by system.","type":"string"}
		ProtocolPort: data["protocol_port"].(int),

		//{"description":"Layer 4 protocol destination port.","type":"integer"}
		Status: data["status"].(string),

		//{"description":"Operating status for this virtual ip.","type":"string"}
		StatusDescription: data["status_description"].(string),

		//{"description":"Operating status description this virtual ip.","type":"string"}
		Protocol: InterfaceToLoadbalancerProtocolType(data["protocol"]),

		//{"description":"IP protocol string like http, https or tcp.","type":"string","enum":["HTTP","HTTPS","TCP","UDP","TERMINATED_HTTPS"]}
		PersistenceCookieName: data["persistence_cookie_name"].(string),

		//{"description":"Set this string if the relation of client and server(pool member) need to persist.","type":"string"}
		PersistenceType: InterfaceToSessionPersistenceType(data["persistence_type"]),

		//{"description":"Method for persistence. HTTP_COOKIE, SOURCE_IP or APP_COOKIE.","type":"string","enum":["SOURCE_IP","HTTP_COOKIE","APP_COOKIE"]}
		SubnetID: InterfaceToUuidStringType(data["subnet_id"]),

		//{"description":"UUID of subnet in which to allocate the Virtual IP.","type":"string"}
		ConnectionLimit: data["connection_limit"].(int),

		//{"description":"Maximum number of concurrent connections","type":"integer"}

	}
}

// InterfaceToVirtualIpTypeSlice makes a slice of VirtualIpType from interface
func InterfaceToVirtualIpTypeSlice(data interface{}) []*VirtualIpType {
	list := data.([]interface{})
	result := MakeVirtualIpTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualIpType(item))
	}
	return result
}

// MakeVirtualIpTypeSlice() makes a slice of VirtualIpType
func MakeVirtualIpTypeSlice() []*VirtualIpType {
	return []*VirtualIpType{}
}
