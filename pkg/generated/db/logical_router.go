package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertLogicalRouterQuery = "insert into `logical_router` (`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`vxlan_network_identifier`,`route_target`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLogicalRouterQuery = "update `logical_router` set `enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`vxlan_network_identifier` = ?,`route_target` = ?,`display_name` = ?,`key_value_pair` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteLogicalRouterQuery = "delete from `logical_router` where uuid = ?"

// LogicalRouterFields is db columns for LogicalRouter
var LogicalRouterFields = []string{
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"owner",
	"owner_access",
	"other_access",
	"group",
	"group_access",
	"vxlan_network_identifier",
	"route_target",
	"display_name",
	"key_value_pair",
	"share",
	"perms2_owner",
	"perms2_owner_access",
	"global_access",
	"uuid",
	"fq_name",
}

// LogicalRouterRefFields is db reference fields for LogicalRouter
var LogicalRouterRefFields = map[string][]string{

	"physical_router": {
	// <common.Schema Value>

	},

	"bgpvpn": {
	// <common.Schema Value>

	},

	"route_target": {
	// <common.Schema Value>

	},

	"virtual_machine_interface": {
	// <common.Schema Value>

	},

	"service_instance": {
	// <common.Schema Value>

	},

	"route_table": {
	// <common.Schema Value>

	},

	"virtual_network": {
	// <common.Schema Value>

	},
}

const insertLogicalRouterRouteTargetQuery = "insert into `ref_logical_router_route_target` (`from`, `to` ) values (?, ?);"

const insertLogicalRouterVirtualMachineInterfaceQuery = "insert into `ref_logical_router_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

const insertLogicalRouterServiceInstanceQuery = "insert into `ref_logical_router_service_instance` (`from`, `to` ) values (?, ?);"

const insertLogicalRouterRouteTableQuery = "insert into `ref_logical_router_route_table` (`from`, `to` ) values (?, ?);"

const insertLogicalRouterVirtualNetworkQuery = "insert into `ref_logical_router_virtual_network` (`from`, `to` ) values (?, ?);"

const insertLogicalRouterPhysicalRouterQuery = "insert into `ref_logical_router_physical_router` (`from`, `to` ) values (?, ?);"

const insertLogicalRouterBGPVPNQuery = "insert into `ref_logical_router_bgpvpn` (`from`, `to` ) values (?, ?);"

// CreateLogicalRouter inserts LogicalRouter to DB
func CreateLogicalRouter(tx *sql.Tx, model *models.LogicalRouter) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLogicalRouterQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertLogicalRouterQuery,
	}).Debug("create query")
	_, err = stmt.Exec(bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.VxlanNetworkIdentifier),
		common.MustJSON(model.ConfiguredRouteTargetList.RouteTarget),
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair),
		common.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.UUID),
		common.MustJSON(model.FQName))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertLogicalRouterVirtualMachineInterfaceQuery)
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

	stmtServiceInstanceRef, err := tx.Prepare(insertLogicalRouterServiceInstanceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceInstanceRefs create statement failed")
	}
	defer stmtServiceInstanceRef.Close()
	for _, ref := range model.ServiceInstanceRefs {

		_, err = stmtServiceInstanceRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceInstanceRefs create failed")
		}
	}

	stmtRouteTableRef, err := tx.Prepare(insertLogicalRouterRouteTableQuery)
	if err != nil {
		return errors.Wrap(err, "preparing RouteTableRefs create statement failed")
	}
	defer stmtRouteTableRef.Close()
	for _, ref := range model.RouteTableRefs {

		_, err = stmtRouteTableRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "RouteTableRefs create failed")
		}
	}

	stmtVirtualNetworkRef, err := tx.Prepare(insertLogicalRouterVirtualNetworkQuery)
	if err != nil {
		return errors.Wrap(err, "preparing VirtualNetworkRefs create statement failed")
	}
	defer stmtVirtualNetworkRef.Close()
	for _, ref := range model.VirtualNetworkRefs {

		_, err = stmtVirtualNetworkRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualNetworkRefs create failed")
		}
	}

	stmtPhysicalRouterRef, err := tx.Prepare(insertLogicalRouterPhysicalRouterQuery)
	if err != nil {
		return errors.Wrap(err, "preparing PhysicalRouterRefs create statement failed")
	}
	defer stmtPhysicalRouterRef.Close()
	for _, ref := range model.PhysicalRouterRefs {

		_, err = stmtPhysicalRouterRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "PhysicalRouterRefs create failed")
		}
	}

	stmtBGPVPNRef, err := tx.Prepare(insertLogicalRouterBGPVPNQuery)
	if err != nil {
		return errors.Wrap(err, "preparing BGPVPNRefs create statement failed")
	}
	defer stmtBGPVPNRef.Close()
	for _, ref := range model.BGPVPNRefs {

		_, err = stmtBGPVPNRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "BGPVPNRefs create failed")
		}
	}

	stmtRouteTargetRef, err := tx.Prepare(insertLogicalRouterRouteTargetQuery)
	if err != nil {
		return errors.Wrap(err, "preparing RouteTargetRefs create statement failed")
	}
	defer stmtRouteTargetRef.Close()
	for _, ref := range model.RouteTargetRefs {

		_, err = stmtRouteTargetRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "RouteTargetRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanLogicalRouter(values map[string]interface{}) (*models.LogicalRouter, error) {
	m := models.MakeLogicalRouter()

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

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

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

	if value, ok := values["vxlan_network_identifier"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VxlanNetworkIdentifier = castedValue

	}

	if value, ok := values["route_target"]; ok {

		json.Unmarshal(value.([]byte), &m.ConfiguredRouteTargetList.RouteTarget)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["perms2_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["perms2_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

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
			referenceModel := &models.LogicalRouterVirtualMachineInterfaceRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_service_instance"]; ok {
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
			referenceModel := &models.LogicalRouterServiceInstanceRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.ServiceInstanceRefs = append(m.ServiceInstanceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_route_table"]; ok {
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
			referenceModel := &models.LogicalRouterRouteTableRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.RouteTableRefs = append(m.RouteTableRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_network"]; ok {
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
			referenceModel := &models.LogicalRouterVirtualNetworkRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.VirtualNetworkRefs = append(m.VirtualNetworkRefs, referenceModel)

		}
	}

	if value, ok := values["ref_physical_router"]; ok {
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
			referenceModel := &models.LogicalRouterPhysicalRouterRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.PhysicalRouterRefs = append(m.PhysicalRouterRefs, referenceModel)

		}
	}

	if value, ok := values["ref_bgpvpn"]; ok {
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
			referenceModel := &models.LogicalRouterBGPVPNRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.BGPVPNRefs = append(m.BGPVPNRefs, referenceModel)

		}
	}

	if value, ok := values["ref_route_target"]; ok {
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
			referenceModel := &models.LogicalRouterRouteTargetRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.RouteTargetRefs = append(m.RouteTargetRefs, referenceModel)

		}
	}

	return m, nil
}

// ListLogicalRouter lists LogicalRouter with list spec.
func ListLogicalRouter(tx *sql.Tx, spec *common.ListSpec) ([]*models.LogicalRouter, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "logical_router"
	spec.Fields = LogicalRouterFields
	spec.RefFields = LogicalRouterRefFields
	result := models.MakeLogicalRouterSlice()
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
		m, err := scanLogicalRouter(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowLogicalRouter shows LogicalRouter resource
func ShowLogicalRouter(tx *sql.Tx, uuid string) (*models.LogicalRouter, error) {
	list, err := ListLogicalRouter(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateLogicalRouter updates a resource
func UpdateLogicalRouter(tx *sql.Tx, uuid string, model *models.LogicalRouter) error {
	//TODO(nati) support update
	return nil
}

// DeleteLogicalRouter deletes a resource
func DeleteLogicalRouter(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteLogicalRouterQuery)
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