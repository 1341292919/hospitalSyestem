package db

import (
	"time"
)

type Doctor struct {
	DoctorId     int64
	Name         string
	Specialty    string
	Title        string
	Password     string
	Department   string
	ContactPhone string
	CreateTime   time.Time
	UpdateTime   time.Time
}

type Nurse struct {
	NurseId      int64
	Name         string
	ContactPhone string
	Position     string
	Password     string
	Department   string
	CreateTime   time.Time
	UpdateTime   time.Time
}
type Admin struct {
	AdminId      int64
	Username     string
	Password     string
	Status       int64
	ContactPhone string
	CreateTime   time.Time
}
type User struct {
	Id           int64
	Name         string
	Identity     string
	ContactPhone string
	Department   string
	Title        string
	Specialty    string
	Position     string
	CreateTime   time.Time
}

type UpdateDoctorRequest struct {
	Name         string
	Specialty    string
	Title        string
	Department   string
	ContactPhone string
	Id           int64
}

type UpdateNurseRequest struct {
	Name         string
	Position     string
	Department   string
	ContactPhone string
	Id           int64
}
type UpdateAdminRequest struct {
	Name         string
	ContactPhone string
	Id           int64
}
