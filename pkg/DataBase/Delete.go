package database

// Удалить значение из базы данных по ID
func (db *DB) Delete(id int) error {
	_, err := db.Exec("delete from bazarakiLis where ID = ?", id)
	return err
}
