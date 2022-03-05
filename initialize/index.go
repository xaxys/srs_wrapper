package initialize

import (
	_ "srs_wrapper/database"
	. "srs_wrapper/initialize/group"
	. "srs_wrapper/initialize/permission"
	. "srs_wrapper/initialize/user"
)

func InitDefaultData() {
	CreateDefaultPermissions()
	CreateDefaultGroups()
	CreateDefaultUsers()
}
