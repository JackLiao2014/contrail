package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertControllerNodeRoleQuery = "insert into `controller_node_role` (`display_name`,`provisioning_state`,`capacity_drives`,`storage_management_bond_interface_members`,`key_value_pair`,`provisioning_progress`,`provisioning_progress_stage`,`provisioning_start_time`,`internalapi_bond_interface_members`,`fq_name`,`global_access`,`share`,`owner`,`owner_access`,`provisioning_log`,`performance_drives`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateControllerNodeRoleQuery = "update `controller_node_role` set `display_name` = ?,`provisioning_state` = ?,`capacity_drives` = ?,`storage_management_bond_interface_members` = ?,`key_value_pair` = ?,`provisioning_progress` = ?,`provisioning_progress_stage` = ?,`provisioning_start_time` = ?,`internalapi_bond_interface_members` = ?,`fq_name` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`provisioning_log` = ?,`performance_drives` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`uuid` = ?;"
const deleteControllerNodeRoleQuery = "delete from `controller_node_role` where uuid = ?"

// ControllerNodeRoleFields is db columns for ControllerNodeRole
var ControllerNodeRoleFields = []string{
	"display_name",
	"provisioning_state",
	"capacity_drives",
	"storage_management_bond_interface_members",
	"key_value_pair",
	"provisioning_progress",
	"provisioning_progress_stage",
	"provisioning_start_time",
	"internalapi_bond_interface_members",
	"fq_name",
	"global_access",
	"share",
	"owner",
	"owner_access",
	"provisioning_log",
	"performance_drives",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"permissions_owner_access",
	"other_access",
	"group",
	"group_access",
	"permissions_owner",
	"enable",
	"uuid",
}

// ControllerNodeRoleRefFields is db reference fields for ControllerNodeRole
var ControllerNodeRoleRefFields = map[string][]string{}

// CreateControllerNodeRole inserts ControllerNodeRole to DB
func CreateControllerNodeRole(tx *sql.Tx, model *models.ControllerNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertControllerNodeRoleQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertControllerNodeRoleQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.DisplayName),
		string(model.ProvisioningState),
		string(model.CapacityDrives),
		string(model.StorageManagementBondInterfaceMembers),
		common.MustJSON(model.Annotations.KeyValuePair),
		int(model.ProvisioningProgress),
		string(model.ProvisioningProgressStage),
		string(model.ProvisioningStartTime),
		string(model.InternalapiBondInterfaceMembers),
		common.MustJSON(model.FQName),
		int(model.Perms2.GlobalAccess),
		common.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.ProvisioningLog),
		string(model.PerformanceDrives),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		bool(model.IDPerms.Enable),
		string(model.UUID))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanControllerNodeRole(values map[string]interface{}) (*models.ControllerNodeRole, error) {
	m := models.MakeControllerNodeRole()

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["provisioning_state"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ProvisioningState = castedValue

	}

	if value, ok := values["capacity_drives"]; ok {

		castedValue := common.InterfaceToString(value)

		m.CapacityDrives = castedValue

	}

	if value, ok := values["storage_management_bond_interface_members"]; ok {

		castedValue := common.InterfaceToString(value)

		m.StorageManagementBondInterfaceMembers = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["provisioning_progress"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ProvisioningProgress = castedValue

	}

	if value, ok := values["provisioning_progress_stage"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ProvisioningProgressStage = castedValue

	}

	if value, ok := values["provisioning_start_time"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ProvisioningStartTime = castedValue

	}

	if value, ok := values["internalapi_bond_interface_members"]; ok {

		castedValue := common.InterfaceToString(value)

		m.InternalapiBondInterfaceMembers = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

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

	if value, ok := values["provisioning_log"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ProvisioningLog = castedValue

	}

	if value, ok := values["performance_drives"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PerformanceDrives = castedValue

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

	if value, ok := values["permissions_owner_access"]; ok {

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

	if value, ok := values["permissions_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	return m, nil
}

// ListControllerNodeRole lists ControllerNodeRole with list spec.
func ListControllerNodeRole(tx *sql.Tx, spec *common.ListSpec) ([]*models.ControllerNodeRole, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "controller_node_role"
	spec.Fields = ControllerNodeRoleFields
	spec.RefFields = ControllerNodeRoleRefFields
	result := models.MakeControllerNodeRoleSlice()
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
		m, err := scanControllerNodeRole(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowControllerNodeRole shows ControllerNodeRole resource
func ShowControllerNodeRole(tx *sql.Tx, uuid string) (*models.ControllerNodeRole, error) {
	list, err := ListControllerNodeRole(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateControllerNodeRole updates a resource
func UpdateControllerNodeRole(tx *sql.Tx, uuid string, model *models.ControllerNodeRole) error {
	//TODO(nati) support update
	return nil
}

// DeleteControllerNodeRole deletes a resource
func DeleteControllerNodeRole(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteControllerNodeRoleQuery)
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