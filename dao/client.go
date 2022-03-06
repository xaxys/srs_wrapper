package dao

import (
	"srs_wrapper/database"
	. "srs_wrapper/model"

	"github.com/patrickmn/go-cache"
)

var guestGroupID = GetGroupByName("guest").ID

func GetGroupByClient(clientID string) *Group {
	groupID, ok := database.Cache.Get(clientID)
	if !ok {
		return nil
	}
	return GetGroupByID(groupID.(uint))
}

func CreateClientWithUserID(clientID string, userID uint) {
	user := GetUserByID(userID)
	database.Cache.Set(clientID, user.GroupID, cache.NoExpiration)
}

func CreateGuestClient(clientID string) {
	database.Cache.Set(clientID, guestGroupID, cache.NoExpiration)
}

func DeleteClient(clientID string) {
	database.Cache.Delete(clientID)
}

func GetAllClientsCount() int {
	return database.Cache.ItemCount()
}
