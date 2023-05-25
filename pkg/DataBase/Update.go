package database

// Обновить цену в определённой записи
func (db *DB) UpdatePrice(id int, Price string) error {
	_, err := db.Exec("UPDATE bazarakiLis SET Price = %d  where ID = ?", Price, id)
	return err
}
