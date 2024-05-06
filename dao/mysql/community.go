package mysql

import (
	"GoofyLink/models"
	"database/sql"

	"go.uber.org/zap"
)

func QueryCommunityList() (communityList []*models.Community, err error) {
	// 查询
	sqlStr := "select community_id, community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("查询社区列表为空")
			err = nil
		}
	}
	return
}
