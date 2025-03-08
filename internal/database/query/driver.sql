-- name: InsertDriver :one
INSERT INTO driver (
    name,
    email,
    secondary_email,
    phone,
    secondary_phone,
    license,
    number,
    secondary_number,
    neighborhood,
    state,
    city,
    cep,
    address,
    address_number,
    country,
    team,
    id_iracing,
    id_steam,
    instagram,
    facebook,
    twitch,
    photo_url,
    fk_created_by_user_id,
    created_date,
    fk_updated_by_user_id,
    updated_date
) VALUES (
    @name::VARCHAR,
    @email::VARCHAR,
    @secondary_email,
    @phone,
    @secondary_phone,
    @license,
    @number,
    @secondary_number,
    @neighborhood,
    @state,
    @city,
    @cep,
    @address,
    @address_number,
    @country,
    @team,
    @id_iracing,
    @id_steam,
    @instagram,
    @facebook,
    @twitch,
    @photo_url,
    @fk_created_by_user_id::BIGINT,
    @created_date::TIMESTAMP,
    @fk_created_by_user_id::BIGINT,
    @created_date::TIMESTAMP
) RETURNING id::BIGINT;

-- name: UpdateDriver :exec
UPDATE driver SET 
    name = @name::VARCHAR,
    email = @email::VARCHAR,
    secondary_email = @secondary_email,
    phone = @phone,
    secondary_phone = @secondary_phone,
    license = @license,
    number = @number,
    secondary_number = @secondary_number,
    neighborhood = @neighborhood,
    state = @state,
    city = @city,
    cep = @cep,
    address = @address,
    address_number = @address_number,
    country = @country,
    team = @team,
    id_iracing = @id_iracing,
    id_steam = @id_steam,
    instagram = @instagram,
    facebook = @facebook,
    twitch = @twitch,
    photo_url = @photo_url,
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
    name::VARCHAR,
    email::VARCHAR,
    phone,
    team
FROM
    driver
WHERE 
    CASE WHEN @search::VARCHAR != '' THEN 
        name ILIKE '%' || @search || '%' OR
        email ILIKE '%' || @search || '%' OR
        phone ILIKE '%' || @search || '%' OR
        team ILIKE '%' || @search || '%'
    ELSE
        TRUE
    END
ORDER BY
    id
OFFSET $1
LIMIT $2;

-- name: SelectCountListDrivers :one
SELECT 
    COUNT(1) AS count
FROM
    driver
WHERE 
    CASE WHEN @search::VARCHAR != '' THEN 
        name ILIKE '%' || @search || '%' OR
        email ILIKE '%' || @search || '%' OR
        phone ILIKE '%' || @search || '%' OR
        team ILIKE '%' || @search || '%'
    ELSE
        TRUE
    END;

-- name: GetDriver :one
SELECT 
    id::BIGINT,
    name::VARCHAR,
    email::VARCHAR,
    secondary_email,
    phone,
    secondary_phone,
    license,
    number,
    secondary_number,
    neighborhood,
    state,
    city,
    cep,
    address,
    address_number,
    country,
    team,
    id_iracing,
    id_steam,
    instagram,
    facebook,
    twitch,
    photo_url,
    irating_formula_car,
    irating_oval,
    irating_sports_car,
    fk_created_by_user_id::BIGINT,
    fk_updated_by_user_id::BIGINT,
    created_date::TIMESTAMP,
    updated_date::TIMESTAMP
FROM
    driver
WHERE id = $1::BIGINT;

-- name: SelectIDIracingByID :one
SELECT 
    id_iracing
FROM
    driver
WHERE id = $1::BIGINT;

-- name: UpdateIratingsByID :exec
UPDATE driver SET 
    irating_sports_car = CASE WHEN @irating_sports_car::INT > 0 THEN @irating_sports_car::INT ELSE irating_sports_car END,
    irating_oval = CASE WHEN @irating_oval::INT > 0 THEN @irating_oval::INT ELSE irating_oval END,
    irating_formula_car = CASE WHEN @irating_formula_car::INT > 0 THEN @irating_formula_car::INT ELSE irating_formula_car END
WHERE id = @id::BIGINT;