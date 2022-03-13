package dao

import (
	"fmt"
	"srs_wrapper/database"

	"github.com/patrickmn/go-cache"
)

var guestGroupID uint

func GetGuestGroupID() uint {
	if guestGroupID == 0 {
		g, _ := GetGroupByName("guest")
		guestGroupID = g.ID
	}
	return guestGroupID
}

func GetGroupIDByClient(clientID string) (uint, error) {
	groupID, ok := database.Cache.Get(clientID)
	if !ok {
		return 0, fmt.Errorf("Relevant group not found")
	}
	return groupID.(uint), nil
}

func CreateClient(clientID string, userID uint) error {
	user, err := GetUserByID(userID)
	if err != nil {
		return err
	}
	database.Cache.Set(clientID, user.GroupID, cache.NoExpiration)
	return nil
}

func CreateGuestClient(clientID string) {
	database.Cache.Set(clientID, GetGuestGroupID(), cache.NoExpiration)
}

func DeleteClient(clientID string) {
	database.Cache.Delete(clientID)
}

func GetAllClientsCount() int {
	return database.Cache.ItemCount()
}
