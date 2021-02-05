package render

import "github.com/guregu/null"

func GetUnixFromNullTime(t null.Time) int64 {
	if t.IsZero() {
		return 0
	}

	return t.Time.Unix()
}
