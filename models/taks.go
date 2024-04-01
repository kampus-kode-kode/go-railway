package models

import "time"

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

var Tasks = []Task{
	{ID: "1", Title: "Memperbaiki Bug", Description: "Memperbaiki bug kritis dalam modul otentikasi", DueDate: time.Now().AddDate(0, 0, 3), Status: "Pending"},
	{ID: "2", Title: "Menambah Fitur", Description: "Menerapkan fungsionalitas pencarian baru", DueDate: time.Now().AddDate(0, 0, 5), Status: "In Progress"},
	{ID: "3", Title: "Refaktor Kode", Description: "Refaktor kode legacy untuk performa yang lebih baik", DueDate: time.Now().AddDate(0, 0, 7), Status: "Pending"},
	{ID: "4", Title: "Menulis Tes", Description: "Menulis tes unit untuk modul manajemen pengguna", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}
