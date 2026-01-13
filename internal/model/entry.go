package model

import  "time"

type Entry struct {
	ID string `json:"id"`
	Message string `json:"message"`
	Date time.Time `json:"date"`
}