-- name: SelectUserIDAndPasswordByEmail :one
SELECT
    id::BIGINT,
    password::VARCHAR,
    status::VARCHAR
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
    AND status = @status::VARCHAR;

-- name: UpdateUserStatus :exec
UPDATE
    "user"
SET
    status = @status::VARCHAR
WHERE
    id = @id::BIGINT;