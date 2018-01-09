package models

// Location

import "encoding/json"

// Location
type Location struct {
	ProvisioningProgress             int            `json:"provisioning_progress"`
	PrivateDNSServers                string         `json:"private_dns_servers"`
	PrivateOspdPackageURL            string         `json:"private_ospd_package_url"`
	PrivateOspdUserName              string         `json:"private_ospd_user_name"`
	PrivateOspdVMRAMMB               string         `json:"private_ospd_vm_ram_mb"`
	Type                             string         `json:"type"`
	UUID                             string         `json:"uuid"`
	FQName                           []string       `json:"fq_name"`
	PrivateRedhatSubscriptionPasword string         `json:"private_redhat_subscription_pasword"`
	GCPRegion                        string         `json:"gcp_region"`
	AwsRegion                        string         `json:"aws_region"`
	ParentType                       string         `json:"parent_type"`
	AwsSubnet                        string         `json:"aws_subnet"`
	Annotations                      *KeyValuePairs `json:"annotations"`
	DisplayName                      string         `json:"display_name"`
	PrivateOspdVMName                string         `json:"private_ospd_vm_name"`
	PrivateOspdVMVcpus               string         `json:"private_ospd_vm_vcpus"`
	PrivateRedhatSubscriptionKey     string         `json:"private_redhat_subscription_key"`
	GCPSubnet                        string         `json:"gcp_subnet"`
	PrivateOspdUserPassword          string         `json:"private_ospd_user_password"`
	PrivateRedhatPoolID              string         `json:"private_redhat_pool_id"`
	PrivateRedhatSubscriptionUser    string         `json:"private_redhat_subscription_user"`
	ProvisioningState                string         `json:"provisioning_state"`
	ProvisioningProgressStage        string         `json:"provisioning_progress_stage"`
	PrivateNTPHosts                  string         `json:"private_ntp_hosts"`
	GCPAccountInfo                   string         `json:"gcp_account_info"`
	AwsSecretKey                     string         `json:"aws_secret_key"`
	IDPerms                          *IdPermsType   `json:"id_perms"`
	PrivateOspdVMDiskGB              string         `json:"private_ospd_vm_disk_gb"`
	GCPAsn                           int            `json:"gcp_asn"`
	AwsAccessKey                     string         `json:"aws_access_key"`
	ParentUUID                       string         `json:"parent_uuid"`
	Perms2                           *PermType2     `json:"perms2"`
	ProvisioningStartTime            string         `json:"provisioning_start_time"`
	ProvisioningLog                  string         `json:"provisioning_log"`

	PhysicalRouters []*PhysicalRouter `json:"physical_routers"`
}

// String returns json representation of the object
func (model *Location) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLocation makes Location
func MakeLocation() *Location {
	return &Location{
		//TODO(nati): Apply default
		PrivateOspdVMRAMMB:               "",
		ProvisioningProgress:             0,
		PrivateDNSServers:                "",
		PrivateOspdPackageURL:            "",
		PrivateOspdUserName:              "",
		Type:                             "",
		UUID:                             "",
		FQName:                           []string{},
		ParentType:                       "",
		PrivateRedhatSubscriptionPasword: "",
		GCPRegion:                        "",
		AwsRegion:                        "",
		GCPSubnet:                        "",
		AwsSubnet:                        "",
		Annotations:                      MakeKeyValuePairs(),
		DisplayName:                      "",
		PrivateOspdVMName:                "",
		PrivateOspdVMVcpus:               "",
		PrivateRedhatSubscriptionKey:     "",
		ProvisioningState:                "",
		PrivateOspdUserPassword:          "",
		PrivateRedhatPoolID:              "",
		PrivateRedhatSubscriptionUser:    "",
		IDPerms:                   MakeIdPermsType(),
		ProvisioningProgressStage: "",
		PrivateNTPHosts:           "",
		GCPAccountInfo:            "",
		AwsSecretKey:              "",
		ParentUUID:                "",
		PrivateOspdVMDiskGB:       "",
		GCPAsn:                    0,
		AwsAccessKey:              "",
		Perms2:                    MakePermType2(),
		ProvisioningStartTime:     "",
		ProvisioningLog:           "",
	}
}

// InterfaceToLocation makes Location from interface
func InterfaceToLocation(iData interface{}) *Location {
	data := iData.(map[string]interface{})
	return &Location{
		ProvisioningProgressStage: data["provisioning_progress_stage"].(string),

		//{"title":"Provisioning Progress Stage","default":"","type":"string","permission":["create","update"]}
		PrivateNTPHosts: data["private_ntp_hosts"].(string),

		//{"title":"NTP Hosts","description":"List of NTP sources","default":"","type":"string","permission":["create","update"]}
		GCPAccountInfo: data["gcp_account_info"].(string),

		//{"title":"Account info","description":"copy and paste contents of account.json","default":"","type":"string","permission":["create","update"]}
		AwsSecretKey: data["aws_secret_key"].(string),

		//{"title":"Secret Key","default":"","type":"string","permission":["create","update"]}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		PrivateOspdVMDiskGB: data["private_ospd_vm_disk_gb"].(string),

		//{"title":"OSPD Disk Size in gigabytes","description":"disk spae to assign to RH OSPD vm","default":"100","type":"string","permission":["create","update"]}
		GCPAsn: data["gcp_asn"].(int),

		//{"title":"ASN","default":65001,"type":"integer","permission":["create","update"]}
		AwsAccessKey: data["aws_access_key"].(string),

		//{"title":"Access Key","default":"","type":"string","permission":["create","update"]}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ProvisioningStartTime: data["provisioning_start_time"].(string),

		//{"title":"Time provisioning started","default":"","type":"string","permission":["create","update"]}
		ProvisioningLog: data["provisioning_log"].(string),

		//{"title":"Provisioning Log","default":"","type":"string","permission":["create","update"]}
		ProvisioningProgress: data["provisioning_progress"].(int),

		//{"title":"Provisioning Progress","default":0,"type":"integer","permission":["create","update"]}
		PrivateDNSServers: data["private_dns_servers"].(string),

		//{"title":"DNS Servers","description":"List of DNS servers","default":"8.8.8.8","type":"string","permission":["create","update"]}
		PrivateOspdPackageURL: data["private_ospd_package_url"].(string),

		//{"title":"Location of OSPD Contrail Networking Packages","description":"Location of Contrail Networking Packages","default":"","type":"string","permission":["create","update"]}
		PrivateOspdUserName: data["private_ospd_user_name"].(string),

		//{"title":"OSPD User Name","description":"OSPD Non-Root User Account","default":"stack","type":"string","permission":["create"]}
		PrivateOspdVMRAMMB: data["private_ospd_vm_ram_mb"].(string),

		//{"title":"OSPD RAM in megabyts","description":"ram to assign to RH OSPD vm","default":"24576","type":"string","permission":["create","update"]}
		Type: data["type"].(string),

		//{"title":"Location Type","description":"Type of location","default":"private","type":"string","permission":["create","update"],"enum":["private","aws","gcp","openstack"]}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		PrivateRedhatSubscriptionPasword: data["private_redhat_subscription_pasword"].(string),

		//{"title":"Redhat Subscription Password","description":"Password for subscription account","default":"","type":"string","permission":["create","update"]}
		GCPRegion: data["gcp_region"].(string),

		//{"title":"Region","default":"us-west1","type":"string","permission":["create","update"]}
		AwsRegion: data["aws_region"].(string),

		//{"title":"Region","default":"us-west-1","type":"string","permission":["create"]}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		AwsSubnet: data["aws_subnet"].(string),

		//{"title":"Subnet","default":"10.0.0.0/16","type":"string","permission":["create","update"]}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		PrivateOspdVMName: data["private_ospd_vm_name"].(string),

		//{"title":"OSPD Virtual Machine Name","description":"Name of RH OSPD VM","default":"undercloud","type":"string","permission":["create"]}
		PrivateOspdVMVcpus: data["private_ospd_vm_vcpus"].(string),

		//{"title":"OSPD vCPUs","description":"vcpus to assign to RH OSPD vm","default":"8","type":"string","permission":["create","update"]}
		PrivateRedhatSubscriptionKey: data["private_redhat_subscription_key"].(string),

		//{"title":"Redhat Subscription Acctivation Key","description":"Subscription Activation Key","default":"","type":"string","permission":["create","update"]}
		GCPSubnet: data["gcp_subnet"].(string),

		//{"title":"Subnet","default":"10.1.0.0/16","type":"string","permission":["create","update"]}
		PrivateOspdUserPassword: data["private_ospd_user_password"].(string),

		//{"title":"OSPD User Passowrd","description":"OSPD Passowrd for account","default":"","type":"string","permission":["create","update"]}
		PrivateRedhatPoolID: data["private_redhat_pool_id"].(string),

		//{"title":"Redhat Pool ID","description":"Repo Pool ID","default":"","type":"string","permission":["create","update"]}
		PrivateRedhatSubscriptionUser: data["private_redhat_subscription_user"].(string),

		//{"title":"Redhat Subscription User","description":"User name for RedHat subscription account","default":"","type":"string","permission":["create","update"]}
		ProvisioningState: data["provisioning_state"].(string),

		//{"title":"Provisioning Status","default":"CREATED","type":"string","permission":["create","update"],"enum":["CREATED","IN_CREATE_PROGRESS","UPDATED","IN_UPDATE_PROGRESS","DELETED","IN_DELETE_PROGRESS","ERROR"]}

	}
}

// InterfaceToLocationSlice makes a slice of Location from interface
func InterfaceToLocationSlice(data interface{}) []*Location {
	list := data.([]interface{})
	result := MakeLocationSlice()
	for _, item := range list {
		result = append(result, InterfaceToLocation(item))
	}
	return result
}

// MakeLocationSlice() makes a slice of Location
func MakeLocationSlice() []*Location {
	return []*Location{}
}
