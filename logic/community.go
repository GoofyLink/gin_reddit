package logic

import (
	"GoofyLink/dao/mysql"
	"GoofyLink/models"
)

func QueryCommunity() (community []*models.Community, err error) {
	return mysql.QueryCommunityList()
}
