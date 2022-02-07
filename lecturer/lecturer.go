package lecturer

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"encore.dev/storage/sqldb"
)

type LECTURER struct {
	ID   string // ID
	NAME string // lecturer's name
}

// generateID generates a random short ID.
func generateID() (string, error) {
	var data [6]byte // 6 bytes of entropy
	if _, err := rand.Read(data[:]); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(data[:]), nil
}

// Insert inserts a LECTURER into the database.
func insert(ctx context.Context, id string, name string) error {
	_, err := sqldb.Exec(ctx, `
        INSERT INTO lecturer (id, name)
        VALUES ($1, $2)
    `, id, name)
	return err
}

// Insert inserts a LECTURER into the database.
func update(ctx context.Context, id string, name string) error {
	_, err := sqldb.Exec(ctx, `
        UPDATE lecturer
        SET name = $2, updated_at = NOW()
		WHERE id = $1
    `, id, name)
	return err
}

// Insert inserts a LECTURER into the database.
func delete(ctx context.Context, id string) error {
	_, err := sqldb.Exec(ctx, `
        DELETE lecturer
		WHERE id = $1
    `, id)
	return err
}

// Get retrieves the original LECTURER for the id.
//encore:api public method=GET path=/lecturer/:id
func GetOne(ctx context.Context, id string) (*LECTURER, error) {
	lecturer := &LECTURER{ID: id}
	err := sqldb.QueryRow(ctx, `
        SELECT name FROM lecturer
        WHERE id = $1
    `, id).Scan(&lecturer.NAME)
	return lecturer, err
}

// Post inserts a LECTURER into the database.
//encore:api public method=POST path=/lecturer
func Post(ctx context.Context, lecturer *LECTURER) (*LECTURER, error) {
	id, err := generateID()
	if err != nil {
		return nil, err
	} else if err := insert(ctx, id, lecturer.NAME); err != nil {
		return nil, err
	}
	return &LECTURER{ID: id, NAME: lecturer.NAME}, nil
}

// Update the LECTURER name with that id.
//encore:api public method=UPDATE path=/update/:id
func UpdateOne(ctx context.Context, id string, lecturer *LECTURER) error {
	err := update(ctx, id, lecturer.NAME)
	return err
}

// Delete the LECTURER  with that id.
//encore:api public method=DELETE path=/delete/:id
func DeleteOne(ctx context.Context, id string) error {
	err := delete(ctx, id)
	return err
}
