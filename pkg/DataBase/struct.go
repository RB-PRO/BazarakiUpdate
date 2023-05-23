package database

import "time"

// Структура, которая будет добавляться в Базу данных
type Data struct {
	ID         int       `json:"id"`          // ID предложения
	Name       string    `json:"name"`        // Название предложения
	Link       string    `json:"link"`        // Ссылка на преложение
	Area       float64   `json:"area"`        // Площадь
	Price      float64   `json:"price"`       // Цена
	TimeCreate time.Time `json:"time_create"` // Дата создания
}
