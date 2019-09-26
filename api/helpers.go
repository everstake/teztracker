package api

type Limits struct {
	limit  uint
	offset uint
}

func (l *Limits) Limit() uint {
	return l.limit
}
func (l *Limits) Offset() uint {
	return l.offset
}

func NewLimiter(limit, offset *int64) *Limits {
	var l, o uint
	if limit != nil {
		l = uint(*limit)
	}
	if offset != nil {
		o = uint(*offset)
	}
	return &Limits{limit: l, offset: o}
}
