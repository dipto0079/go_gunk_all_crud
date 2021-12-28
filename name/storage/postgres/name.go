package postgres

import (
	"free_go/name/storage"
	"context"
)

const insertName = `
	INSERT INTO names(
		title
	)VALUES(
		:title
	)RETURNING id;
`

func (s *Storage) Create(ctx context.Context, t storage.Name) (int64, error) {
	stmt, err := s.db.PrepareNamed(insertName)
	if err != nil {
		return 0, err
	}
	var id int64
	if err := stmt.Get(&id, t); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Storage) ListName(ctx context.Context) ([]storage.Name, error) {

	var b []storage.Name

	if err := s.db.Select(&b, "SELECT * FROM names"); err != nil {
		return b, err
	}
	return b, nil
}

func (s *Storage) GetBlog(ctx context.Context, id int64) (storage.Name, error) {
	var b storage.Name

	if err := s.db.Get(&b, "SELECT * FROM names WHERE id=$1", id); err != nil {
		return b, err
	}
	return b, nil
}

const UpdateName = `
	UPDATE names 
	SET
		title =:title
	WHERE
	id =:id
	RETURNING *;
`

func (s *Storage) UpdateBlog(ctx context.Context, t storage.Name) error {
	stmt, err := s.db.PrepareNamed(UpdateName)
	if err != nil {
		return err
	}
	var blog storage.Name
	if err := stmt.Get(&blog, t); err != nil {
		return err
	}
	return nil
}

func (s *Storage) BlogDelete(ctx context.Context, id int64) error {
	var b storage.Name
	if err := s.db.Get(&b, "DELETE FROM names WHERE id=$1 RETURNING * ", id); err != nil {
		return err
	}
	return nil
}
