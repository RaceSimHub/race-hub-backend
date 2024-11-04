-- name: InsertNotification :one
INSERT INTO notification (
    message, 
    first_driver, 
    second_driver, 
    third_driver, 
    license_points, 
    created_date
) VALUES (
    @message::VARCHAR, 
    @first_driver::VARCHAR,
    @second_driver::VARCHAR,
    @third_driver::VARCHAR,
    @license_points::INTEGER,
    @created_date::TIMESTAMP
) RETURNING id;

-- name: UpdateNotification :exec
UPDATE notification SET 
    message = COALESCE(@message::VARCHAR, message), 
    first_driver = COALESCE(@first_driver::VARCHAR, first_driver),
    second_driver = COALESCE(@second_driver::VARCHAR, second_driver),
    third_driver = COALESCE(@third_driver::VARCHAR, third_driver),
    license_points = COALESCE(@license_points::INTEGER, license_points)
WHERE id = @id::BIGINT;

-- name: DeleteNotification :exec
DELETE FROM 
    notification
WHERE 
    id = @id::BIGINT;

-- name: SelectListNotifications :many
SELECT 
    id::BIGINT, 
    message::VARCHAR, 
    first_driver::VARCHAR, 
    second_driver::VARCHAR, 
    third_driver::VARCHAR, 
    license_points::INTEGER, 
    created_date::TIMESTAMP
FROM
    notification
OFFSET $1::INTEGER
LIMIT $2::INTEGER;

-- name: GetLastNotificationMessage :one
SELECT 
    message::VARCHAR
FROM
    notification    
ORDER BY id DESC
LIMIT 1;