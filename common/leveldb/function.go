// package
package leveldb

import "encoding/json"

// Models
type Models interface {
	GetID() string
	GetDB() string
}

// GetOne
//
//	@param id
//	@return *T
//	@return error
func GetOne[T Models](id string) (*T, error) {
	var res T
	ldb, err := NewLDB(res.GetDB())
	if err != nil {
		return nil, err
	}
	data, err := ldb.Do(id)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &res)
	return &res, err
}

// GetAll
//
//	@return []T
//	@return error
func GetAll[T Models]() ([]T, error) {
	var t T
	res := make([]T, 0)
	ldb, err := NewLDB(t.GetDB())
	if err != nil {
		return nil, err
	}
	iter := ldb.Db().NewIterator(nil, nil)
	for iter.Next() {
		v, err := GetOne[T](string(iter.Key()))
		if err != nil {
			return nil, err
		}
		res = append(res, *v)
	}
	iter.Release()
	return res, nil
}

// UpdataOne
//
//	@param t
//	@return error
func UpdataOne[T Models](t T) error {
	ldb, err := NewLDB(t.GetDB())
	if err != nil {
		return err
	}
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	_, err = ldb.Do(t.GetID(), data)
	return err
}

// DeleteOne
//
//	@param id
//	@return error
func DeleteOne[T Models](id string) error {
	var t T
	ldb, err := NewLDB(t.GetDB())
	if err != nil {
		return err
	}
	_, err = ldb.Do(id, nil)
	return err
}
