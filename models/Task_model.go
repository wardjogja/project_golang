package Model

import (
	"gorm.io/gorm"
	
)

type Task_model struct {
	gorm.Model
	ID int `gorm:"primaryKey;autoIncrement:true"`
	Nama string
	Pegawai string
	Status string
	Tgl string
}