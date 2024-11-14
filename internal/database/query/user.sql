-- name: SelectUserIDAndPasswordByEmail :one
SELECT
    id::BIGINT,
    password::VARCHAR
FROM
    "user"
WHERE
    email = @email::VARCHAR;

-- name: InsertUser :one
INSERT INTO "user" (
    email,
    name,
    password,
    created_date
) VALUES (
    @email::VARCHAR,
    @name::VARCHAR,
    @password::VARCHAR,
    now()
) RETURNING id;