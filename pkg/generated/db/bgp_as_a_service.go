package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertBGPAsAServiceQuery = "insert into `bgp_as_a_service` (`share`,`owner`,`owner_access`,`global_access`,`bgpaas_session_attributes`,`bgpaas_ipv4_mapped_ipv6_nexthop`,`bgpaas_ip_address`,`fq_name`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`display_name`,`bgpaas_shared`,`bgpaas_suppress_route_advertisement`,`autonomous_system`,`uuid`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateBGPAsAServiceQuery = "update `bgp_as_a_service` set `share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`bgpaas_session_attributes` = ?,`bgpaas_ipv4_mapped_ipv6_nexthop` = ?,`bgpaas_ip_address` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`bgpaas_shared` = ?,`bgpaas_suppress_route_advertisement` = ?,`autonomous_system` = ?,`uuid` = ?,`key_value_pair` = ?;"
const deleteBGPAsAServiceQuery = "delete from `bgp_as_a_service` where uuid = ?"

// BGPAsAServiceFields is db columns for BGPAsAService
var BGPAsAServiceFields = []string{
	"share",
	"owner",
	"owner_access",
	"global_access",
	"bgpaas_session_attributes",
	"bgpaas_ipv4_mapped_ipv6_nexthop",
	"bgpaas_ip_address",
	"fq_name",
	"creator",
	"user_visible",
	"last_modified",
	"other_access",
	"group",
	"group_access",
	"permissions_owner",
	"permissions_owner_access",
	"enable",
	"description",
	"created",
	"display_name",
	"bgpaas_shared",
	"bgpaas_suppress_route_advertisement",
	"autonomous_system",
	"uuid",
	"key_value_pair",
}

// BGPAsAServiceRefFields is db reference fields for BGPAsAService
var BGPAsAServiceRefFields = map[string][]string{

	"virtual_machine_interface": {
	// <common.Schema Value>

	},

	"service_health_check": {
	// <common.Schema Value>

	},
}

const insertBGPAsAServiceVirtualMachineInterfaceQuery = "insert into `ref_bgp_as_a_service_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

const insertBGPAsAServiceServiceHealthCheckQuery = "insert into `ref_bgp_as_a_service_service_health_check` (`from`, `to` ) values (?, ?);"

// CreateBGPAsAService inserts BGPAsAService to DB
func CreateBGPAsAService(tx *sql.Tx, model *models.BGPAsAService) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBGPAsAServiceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertBGPAsAServiceQuery,
	}).Debug("create query")
	_, err = stmt.Exec(common.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.BgpaasSessionAttributes),
		bool(model.BgpaasIpv4MappedIpv6Nexthop),
		string(model.BgpaasIPAddress),
		common.MustJSON(model.FQName),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.DisplayName),
		bool(model.BgpaasShared),
		bool(model.BgpaasSuppressRouteAdvertisement),
		int(model.AutonomousSystem),
		string(model.UUID),
		common.MustJSON(model.Annotations.KeyValuePair))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertBGPAsAServiceVirtualMachineInterfaceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing VirtualMachineInterfaceRefs create statement failed")
	}
	defer stmtVirtualMachineInterfaceRef.Close()
	for _, ref := range model.VirtualMachineInterfaceRefs {

		_, err = stmtVirtualMachineInterfaceRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualMachineInterfaceRefs create failed")
		}
	}

	stmtServiceHealthCheckRef, err := tx.Prepare(insertBGPAsAServiceServiceHealthCheckQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceHealthCheckRefs create statement failed")
	}
	defer stmtServiceHealthCheckRef.Close()
	for _, ref := range model.ServiceHealthCheckRefs {

		_, err = stmtServiceHealthCheckRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceHealthCheckRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanBGPAsAService(values map[string]interface{}) (*models.BGPAsAService, error) {
	m := models.MakeBGPAsAService()

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["bgpaas_session_attributes"]; ok {

		castedValue := common.InterfaceToString(value)

		m.BgpaasSessionAttributes = castedValue

	}

	if value, ok := values["bgpaas_ipv4_mapped_ipv6_nexthop"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.BgpaasIpv4MappedIpv6Nexthop = castedValue

	}

	if value, ok := values["bgpaas_ip_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.BgpaasIPAddress = models.IpAddressType(castedValue)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["creator"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["other_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["bgpaas_shared"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.BgpaasShared = castedValue

	}

	if value, ok := values["bgpaas_suppress_route_advertisement"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.BgpaasSuppressRouteAdvertisement = castedValue

	}

	if value, ok := values["autonomous_system"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.AutonomousSystem = models.AutonomousSystemType(castedValue)

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_virtual_machine_interface"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.BGPAsAServiceVirtualMachineInterfaceRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_service_health_check"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.BGPAsAServiceServiceHealthCheckRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.ServiceHealthCheckRefs = append(m.ServiceHealthCheckRefs, referenceModel)

		}
	}

	return m, nil
}

// ListBGPAsAService lists BGPAsAService with list spec.
func ListBGPAsAService(tx *sql.Tx, spec *common.ListSpec) ([]*models.BGPAsAService, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "bgp_as_a_service"
	spec.Fields = BGPAsAServiceFields
	spec.RefFields = BGPAsAServiceRefFields
	result := models.MakeBGPAsAServiceSlice()
	query, columns, values := common.BuildListQuery(spec)
	log.WithFields(log.Fields{
		"listSpec": spec,
		"query":    query,
	}).Debug("select query")
	rows, err = tx.Query(query, values...)
	if err != nil {
		return nil, errors.Wrap(err, "select query failed")
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "row error")
	}
	for rows.Next() {
		valuesMap := map[string]interface{}{}
		values := make([]interface{}, len(columns))
		valuesPointers := make([]interface{}, len(columns))
		for _, index := range columns {
			valuesPointers[index] = &values[index]
		}
		if err := rows.Scan(valuesPointers...); err != nil {
			return nil, errors.Wrap(err, "scan failed")
		}
		for column, index := range columns {
			val := valuesPointers[index].(*interface{})
			valuesMap[column] = *val
		}
		log.WithFields(log.Fields{
			"valuesMap": valuesMap,
		}).Debug("valueMap")
		m, err := scanBGPAsAService(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowBGPAsAService shows BGPAsAService resource
func ShowBGPAsAService(tx *sql.Tx, uuid string) (*models.BGPAsAService, error) {
	list, err := ListBGPAsAService(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateBGPAsAService updates a resource
func UpdateBGPAsAService(tx *sql.Tx, uuid string, model *models.BGPAsAService) error {
	//TODO(nati) support update
	return nil
}

// DeleteBGPAsAService deletes a resource
func DeleteBGPAsAService(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteBGPAsAServiceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing delete query failed")
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	if err != nil {
		return errors.Wrap(err, "delete failed")
	}
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return nil
}