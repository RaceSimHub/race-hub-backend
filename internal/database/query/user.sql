-- name: SelectUserByEmail :one
SELECT
    id::BIGINT,
    name::VARCHAR,
    password::VARCHAR,
    status::VARCHAR,
    role::VARCHAR
FROM
    "user"
WHERE
    email = @email::VARCHAR;

-- name: InsertUser :one
INSERT INTO "user" (
    email,
    name,
    password,
    status,
    email_verification_token,
    email_verification_expires_at,
    role,
    created_date
) VALUES (
    @email::VARCHAR,
    @name::VARCHAR,
    @password::VARCHAR,
    @status::VARCHAR,
    @email_verification_token::VARCHAR,
    @email_verification_expires_at::TIMESTAMP,
    @role::VARCHAR,
    now()
) RETURNING id;

-- name: SelectUserByEmailVerificationToken :one
SELECT
    id::BIGINT,
    email_verification_expires_at::TIMESTAMP
FROM
    "user"
WHERE
    email_verification_token = @email_verification_token::VARCHAR
    AND email = @email::VARCHAR
    AND CASE WHEN @status <> '' THEN status = @status::VARCHAR ELSE TRUE END;

-- name: UpdateUserStatus :exec
UPDATE
    "user"
SET
    status = @status::VARCHAR
WHERE
    id = @id::BIGINT;

-- name: UpdateUserPassword :exec
UPDATE
    "user"
SET
    password = @password::VARCHAR,
    status = @status::VARCHAR
WHERE
    id = @id::BIGINT;

-- name: UpdateUserEmailVerificationToken :exec
UPDATE
    "user"
SET
    email_verification_token = @email_verification_token::VARCHAR,
    email_verification_expires_at = @email_verification_expires_at::TIMESTAMP
WHERE
    id = @id::BIGINT;