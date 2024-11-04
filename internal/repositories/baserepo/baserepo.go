package baserepo

import (
	"fmt"
	"reflect"

	"github.com/gocraft/dbr/v2"
)

type BaseRepository[T any] struct {
	db *dbr.Session
}

func New[T any](db *dbr.Session) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

func getColumns(record interface{}) []string {
	val := reflect.ValueOf(record).Elem()
	var columns []string

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		dbTag := field.Tag.Get("db")
		if dbTag != "" {
			columns = append(columns, dbTag)
		}
	}
	return columns
}

func (r *BaseRepository[T]) Insert(table string, record *T) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.RollbackUnlessCommitted()

	_, err = tx.InsertInto(table).Columns(getColumns(record)...).Record(record).Exec()
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *BaseRepository[T]) List(table string) ([]T, error) {
	var results []T
	_, err := r.db.Select("*").From(table).Load(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r *BaseRepository[T]) Get(table string, idField string, idValue interface{}) (*T, error) {
	var result T
	_, err := r.db.Select("*").From(table).Where(fmt.Sprintf("%s = ?", idField), idValue).Load(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *BaseRepository[T]) Update(table string, idField string, idValue interface{}, updates map[string]interface{}) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.RollbackUnlessCommitted()

	_, err = tx.Update(table).SetMap(updates).Where(dbr.Eq(idField, idValue)).Exec()
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *BaseRepository[T]) DB() *dbr.Session {
	return r.db
}
