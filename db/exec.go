package db

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jinzhu/gorm"
)

// SelectBuilder builds SelectBuilder
func SelectBuilder(s sq.SelectBuilder, eq map[string][]interface{}, gt, lt, gte, lte map[string]interface{}) sq.SelectBuilder {
	for k, v := range eq {
		switch {
		case len(v) == 1:
			eqs := sq.Eq{k: v[0]}
			s = s.Where(eqs)
		case len(v) > 1:
			eqs := sq.Eq{k: v}
			s = s.Where(eqs)
		}
	}
	if len(gt) > 0 {
		m := sq.Gt(gt)
		s = s.Where(m)
	}
	if len(lt) > 0 {
		m := sq.Lt(lt)
		s = s.Where(m)
	}
	if len(gte) > 0 {
		m := sq.GtOrEq(gte)
		s = s.Where(m)
	}
	if len(lte) > 0 {
		m := sq.LtOrEq(lte)
		s = s.Where(m)
	}
	return s
}

// UpdateBuilder builds UpdateBuilder
func UpdateBuilder(s sq.UpdateBuilder, eq map[string][]interface{}, gt, lt, gte, lte map[string]interface{}) sq.UpdateBuilder {
	for k, v := range eq {
		switch {
		case len(v) == 1:
			eqs := sq.Eq{k: v[0]}
			s = s.Where(eqs)
		case len(v) > 1:
			eqs := sq.Eq{k: v}
			s = s.Where(eqs)
		}
	}
	if len(gt) > 0 {
		m := sq.Gt(gt)
		s = s.Where(m)
	}
	if len(lt) > 0 {
		m := sq.Lt(lt)
		s = s.Where(m)
	}
	if len(gte) > 0 {
		m := sq.GtOrEq(gte)
		s = s.Where(m)
	}
	if len(lte) > 0 {
		m := sq.LtOrEq(lte)
		s = s.Where(m)
	}
	return s
}

// ExecUpdate execute raw sql
func ExecUpdate(db *gorm.DB, active sq.UpdateBuilder) (err error) {
	sql, args, err := active.ToSql()
	if err != nil {
		return err
	}
	return db.Exec(sql, args...).Error
}
