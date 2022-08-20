package mysql

import (
	"database/sql"

	"github.com/bavlayan/GoToDo/pkg/models"
)

type TodoItemModel struct {
	DB *sql.DB
}

func (t *TodoItemModel) Save(description string) (int, error) {
	sql_query := `INSERT INTO tbl_todoitems (id, description) VALUES(uuid(), ?)`

	result, err := t.DB.Exec(sql_query, description)
	if err != nil {
		return 0, err
	}

	affected_rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(affected_rows), nil
}

func (t *TodoItemModel) Get(id string) (*models.TodoItems, error) {
	sql_query := `SELECT * FROM tbl_todoitems WHERE id = ?`

	row := t.DB.QueryRow(sql_query, id)
	td := &models.TodoItems{}

	err := row.Scan(
		&td.ID,
		&td.Completed,
		&td.CreatedDate,
		&td.Description,
		&td.Deleted,
	)

	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return td, nil
}

func (t *TodoItemModel) GetDaily() ([]*models.TodoItems, error) {
	sql_query := `SELECT * FROM tbl_todoitems WHERE created_date <= now() AND created_date >= DATE_SUB(NOW(), INTERVAL 1 DAY)`

	rows, err := t.DB.Query(sql_query)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	defer rows.Close()

	td_items := []*models.TodoItems{}

	for rows.Next() {
		td := &models.TodoItems{}
		err = rows.Scan(
			&td.ID,
			&td.Completed,
			&td.CreatedDate,
			&td.Description,
			&td.Deleted,
		)

		if err != nil {
			return nil, err
		}
		td_items = append(td_items, td)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return td_items, nil

}
