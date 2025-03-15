-- name: InsertTrack :one
INSERT INTO track (
    name,
    country,
    fk_created_by_user_id
) VALUES (
    @name::VARCHAR,
    @country::VARCHAR,
    @fk_created_by_user_id::BIGINT
) RETURNING id;

-- name: UpdateTrack :exec
UPDATE track SET
    name = COALESCE(@name::VARCHAR, name),
    country = COALESCE(@country::VARCHAR, country),
    fk_updated_by_user_id = @fk_updated_by_user_id::BIGINT
WHERE 
    id = @id::BIGINT;

-- name: DeleteTrack :exec
DELETE FROM
    track
WHERE
    id = @id::BIGINT;

-- name: SelectListTracks :many
SELECT
    id::BIGINT,
    name::VARCHAR,
    country::VARCHAR
FROM
    track
WHERE 
    CASE WHEN @search::VARCHAR != '' THEN 
        name ILIKE '%' || @search || '%' OR
        country ILIKE '%' || @search || '%'
    ELSE
        TRUE
    END
ORDER BY
    id DESC
OFFSET $1::INTEGER
LIMIT $2::INTEGER;

-- name: SelectListTracksCount :one
SELECT
    COUNT(*) AS count
FROM
    track
WHERE 
    CASE WHEN @search::VARCHAR != '' THEN 
        name ILIKE '%' || @search || '%' OR
        country ILIKE '%' || @search || '%'
    ELSE
        TRUE
    END;

-- name: SelectTrackById :one
SELECT
    id::BIGINT,
    name::VARCHAR,
    country::VARCHAR,
    created_date::TIMESTAMP,
    updated_date
FROM
    track
WHERE
    id = $1::BIGINT;