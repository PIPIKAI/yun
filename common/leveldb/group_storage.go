package leveldb

import (
	"encoding/json"

	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/models"
)

func (ldb *LevelDb) UpdateGroup(g models.Group) error {

	ldbData, err := json.Marshal(g)
	if err != nil {
		return err
	}
	ldb.Do(g.Name, ldbData)
	return nil

}
func (ldb *LevelDb) UpdateStorage(g models.Storage) error {
	group, _ := ldb.Do(g.Group)

	if group == nil {
		newGroup := models.Group{
			Name:     g.Group,
			Cap:      g.Cap,
			Status:   "work",
			Storages: make(map[string]models.Storage),
		}
		newGroup.Storages[g.ServerAddr] = g
		ldbData, err := json.Marshal(newGroup)
		if err != nil {
			return err
		}
		ldb.Do(g.Group, ldbData)
		return nil
	}
	var nowGroup models.Group
	err := json.Unmarshal(group, &nowGroup)
	if err != nil {
		return err
	}
	nowCap := g.Cap

	for _, v := range nowGroup.Storages {
		if v.Status == "work" {
			nowGroup.Status = "work"
		}
		if nowCap > v.Cap {
			nowCap = v.Cap
		}
	}
	nowGroup.Cap = nowCap
	nowGroup.Storages[g.ServerAddr] = g

	ldbData, err := json.Marshal(nowGroup)
	if err != nil {
		return err
	}
	logger.Logger.Info(string(ldbData))
	ldb.Do(g.Group, ldbData)
	return nil
}
func (ldb *LevelDb) GetGroup(groupName string) (*models.Group, error) {
	var group models.Group
	v, err := ldb.Do(groupName)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(v, &group); err != nil {
		return nil, err
	}
	return &group, nil
}

func (ldb *LevelDb) GetAllGroups() (groups []models.Group, err error) {
	iter := ldb.Db().NewIterator(nil, nil)
	for iter.Next() {
		var g models.Group
		err := json.Unmarshal(iter.Value(), &g)
		if err != nil {
			continue
		}
		groups = append(groups, g)
	}
	iter.Release()
	if len(groups) <= 0 {
		return nil, nil
	}
	return groups, nil
}
