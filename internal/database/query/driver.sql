-- name: InsertDriver :one
INSERT INTO driver (
    name,
    race_name,
    email,
    phone,
    fk_created_by_user_id,
    created_date
) VALUES (
    @name::VARCHAR,
    @race_name::VARCHAR,
    @email::VARCHAR,
    @phone::VARCHAR,
    @fk_created_by_user_id::BIGINT,
    @created_date::TIMESTAMP
) RETURNING id;

-- name: UpdateDriver :exec
UPDATE driver SET 
    name = COALESCE(@name::VARCHAR, name),
    race_name = COALESCE(@race_name::VARCHAR, race_name),
    email = COALESCE(@email::VARCHAR, email),
    phone = COALESCE(@phone::VARCHAR, phone),
    fk_updated_by_user_id = @fk_updated_by_user_id::BIGINT,
    updated_date = @updated_date::TIMESTAMP
WHERE id = @id::BIGINT;

-- name: DeleteDriver :exec
DELETE FROM 
    driver
WHERE 
    id = @id::BIGINT;

-- name: SelectListDrivers :many
SELECT 
    id::BIGINT,
    name::VARCHAR
FROM
    driver
OFFSET $1::INTEGER
LIMIT $2::INTEGER;

-- name: GetDriver :one
SELECT 
    id::BIGINT,
    name::VARCHAR,
    race_name::VARCHAR,
    email::VARCHAR,
    phone::VARCHAR,
    fk_created_by_user_id::BIGINT,
    fk_updated_by_user_id::BIGINT,
    created_date::TIMESTAMP,
    updated_date::TIMESTAMP
FROM
    driver
WHERE id = $1::BIGINT;
