-- name: InsertDriverLink :one
INSERT INTO driver_link (
    fk_driver_id,
    fk_user_id,
    status,
    created_at
) VALUES (
    @fk_driver_id::BIGINT,
    @fk_user_id::BIGINT,
    @status::VARCHAR,
    @created_at::TIMESTAMP
) RETURNING id::BIGINT;

-- name: SelectDriverLinks :many
SELECT
    driver_link.id,
    driver.name AS driver_name,
    "user".name AS user_name,
    driver_link.status,
    created_at
FROM 
    driver_link
INNER JOIN "user" ON driver_link.fk_user_id = "user".id
INNER JOIN driver ON driver_link.fk_driver_id = driver.id
WHERE
    CASE WHEN @search::VARCHAR != '' THEN 
        driver.name ILIKE '%' || @search || '%' OR
        "user".name ILIKE '%' || @search || '%'
    ELSE
        TRUE
    END
ORDER BY 
    driver_link.id DESC
OFFSET $1
LIMIT $2;

-- name: SelectCountDriverLinks :one
SELECT 
    COUNT(1) AS count
FROM
    driver_link
INNER JOIN "user" ON driver_link.fk_user_id = "user".id
INNER JOIN driver ON driver_link.fk_driver_id = driver.id
WHERE 
    CASE WHEN @search::VARCHAR != '' THEN 
        driver.name ILIKE '%' || @search || '%' OR
        "user".name ILIKE '%' || @search || '%'
    ELSE
        TRUE
    END;

-- name: GetDriverLink :one
SELECT
    driver_link.id,
    driver.name AS driver_name,
    "user".name AS user_name,
    driver_link.status,
    created_at
FROM
    driver_link
INNER JOIN "user" ON driver_link.fk_user_id = "user".id
INNER JOIN driver ON driver_link.fk_driver_id = driver.id
WHERE
    id = @id::BIGINT;

-- name: SelectDriverLinkStatusByUserID :one
SELECT
    driver_link.status::VARCHAR
FROM
    driver_link
WHERE
    fk_user_id = @fk_user_id::BIGINT;