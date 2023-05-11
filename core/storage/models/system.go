package models

import "time"

// SystemInfo
var SystemInfo System

// System
type System struct {
	BeginTime time.Time
}
