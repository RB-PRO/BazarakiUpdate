package database

// Получить значение из базы данных по ID
func (db *DB) Select(id int) (Data, error) {
	var DataQuery Data
	ErrorQuery := db.QueryRow("SELECT * FROM bazarakiLis WHERE ID = ?", id).Scan(&DataQuery.ID, &DataQuery.Name, &DataQuery.Link, &DataQuery.Area, &DataQuery.Price, &DataQuery.Rubric, &DataQuery.TimeCreate)
	if ErrorQuery != nil {
		return Data{}, ErrorQuery
	}
	return DataQuery, nil
}

// Получить все значения из базы данных по ID
func (db *DB) Selects() ([]Data, error) {
	rows, err := db.Query("SELECT * FROM bazarakiLis")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Срез альбомов для хранения данных из возвращенных строк.
	var DatasQuery []Data

	// Цикл по строкам,
	// используя Scan для назначения данных столбца полям структуры.
	for rows.Next() {
		var DataQuery Data
		if err := rows.Scan(&DataQuery.ID, &DataQuery.Name, &DataQuery.Link, &DataQuery.Area, &DataQuery.Price, &DataQuery.Rubric, &DataQuery.TimeCreate); err != nil {
			return DatasQuery, err
		}
		DatasQuery = append(DatasQuery, DataQuery)
	}
	if err = rows.Err(); err != nil {
		return DatasQuery, err
	}
	return DatasQuery, nil
}
