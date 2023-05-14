package entities

import "time"

type User struct {
	id_nasabah int64
	username   string
	email      string
	password   string
	photo      int
	create_at  time.Time
	update_at  time.Time
}
