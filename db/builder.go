package db

import sq "github.com/Masterminds/squirrel"

// SelectBuilder ...
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
