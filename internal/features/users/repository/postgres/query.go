package postgres

var (
	insertQuery = `
	INSERT INTO todo.users (id, version, email, email_verified, password_hash, full_name, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id, version, email, email_verified, password_hash, full_name, created_at, updated_at;
	`
	findByIDQuery = `
	SELECT id, version, email, email_verified, password_hash, full_name, created_at, updated_at
	FROM todo.users
	WHERE id=$1;
	`
	listQuery = `
	SELECT id, version, email, email_verified, password_hash, full_name, created_at, updated_at
	FROM todo.users
	ORDER BY id ASC
	LIMIT $1
	OFFSET $2;
	`
	deleteQuery = `
	DELETE FROM todo.users
	WHERE id=$1;
	`
)
