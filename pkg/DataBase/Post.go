package database

// Добавить значение в базу данных
func (db *DB) Incert(data Data) error {

	_, err := db.Exec(`INSERT INTO bazarakiLis (ID, Name, Link, Area, Price, TimeCreate) VALUES (?, ?, ?, ?, ?, ?)`, data.ID, data.Name, data.Link, data.Area, data.Price, data.TimeCreate)

	return err
}
