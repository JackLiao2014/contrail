package models

// JunosServicePorts

import "encoding/json"

// JunosServicePorts
type JunosServicePorts struct {
	ServicePort []string `json:"service_port"`
}

//  parents relation object

// String returns json representation of the object
func (model *JunosServicePorts) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeJunosServicePorts makes JunosServicePorts
func MakeJunosServicePorts() *JunosServicePorts {
	return &JunosServicePorts{
		//TODO(nati): Apply default
		ServicePort: []string{},
	}
}

// InterfaceToJunosServicePorts makes JunosServicePorts from interface
func InterfaceToJunosServicePorts(iData interface{}) *JunosServicePorts {
	data := iData.(map[string]interface{})
	return &JunosServicePorts{
		ServicePort: data["service_port"].([]string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ServicePort","GoType":"string","GoPremitive":true},"GoName":"ServicePort","GoType":"[]string","GoPremitive":true}

	}
}

// InterfaceToJunosServicePortsSlice makes a slice of JunosServicePorts from interface
func InterfaceToJunosServicePortsSlice(data interface{}) []*JunosServicePorts {
	list := data.([]interface{})
	result := MakeJunosServicePortsSlice()
	for _, item := range list {
		result = append(result, InterfaceToJunosServicePorts(item))
	}
	return result
}

// MakeJunosServicePortsSlice() makes a slice of JunosServicePorts
func MakeJunosServicePortsSlice() []*JunosServicePorts {
	return []*JunosServicePorts{}
}