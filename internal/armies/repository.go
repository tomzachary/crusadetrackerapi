package armies

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAllArmies() ([]Army, error) {
	rows, err := r.db.Query(
		`SELECT id, title, description, userId, createdAt, modifiedAt, isDeleted
		 FROM army
		 WHERE isDeleted = false`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	armies := make([]Army, 0)
	for rows.Next() {
		var a Army
		var userId int
		if err := rows.Scan(&a.Id, &a.Name, &a.Description, &userId, &a.CreateTime, &a.ModifiedTime, &a.IsDeleted); err != nil {
			return nil, err
		}
		armies = append(armies, a)
	}
	return armies, rows.Err()
}

func (r *Repository) CreateArmy(army Army) (Army, error) {
	var userId int
	err := r.db.QueryRow(
		`INSERT INTO army (title, description, userId, createdAt, modifiedAt, isDeleted)
		 VALUES ($1, $2, $3, NOW(), NOW(), false)
		 RETURNING id, title, description, userId, createdAt, modifiedAt, isDeleted`,
		army.Name, army.Description, army.UserId,
	).Scan(&army.Id, &army.Name, &army.Description, &userId, &army.CreateTime, &army.ModifiedTime, &army.IsDeleted)
	return army, err
}

func (r *Repository) UpdateArmy(id int, army Army) (Army, error) {
	var userId int
	err := r.db.QueryRow(
		`UPDATE army
		 SET title = $1, description = $2, modifiedAt = NOW()
		 WHERE id = $3 AND isDeleted = false
		 RETURNING id, title, description, userId, createdAt, modifiedAt, isDeleted`,
		army.Name, army.Description, id,
	).Scan(&army.Id, &army.Name, &army.Description, &userId, &army.CreateTime, &army.ModifiedTime, &army.IsDeleted)
	return army, err
}

func (r *Repository) DeleteArmy(id int) error {
	_, err := r.db.Exec(`UPDATE army SET isDeleted = true WHERE id = $1`, id)
	return err
}
