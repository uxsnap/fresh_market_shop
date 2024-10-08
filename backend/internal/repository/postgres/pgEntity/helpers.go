package pgEntity

import (
	"time"

	"github.com/jackc/pgtype"
	uuid "github.com/satori/go.uuid"
)

func pgStatusFromTime(t time.Time) pgtype.Status {
	if t.Unix() == 0 {
		return pgtype.Null
	}
	return pgtype.Present
}

func pgUidFromUUID(uid uuid.UUID) pgtype.UUID {
	return pgtype.UUID{
		Bytes:  uid,
		Status: pgtype.Present,
	}
}
