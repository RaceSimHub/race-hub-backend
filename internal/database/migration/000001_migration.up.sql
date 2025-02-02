CREATE TABLE "user" (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_date TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE track (
    id BIGSERIAL PRIMARY KEY,
    fk_created_by_user_id BIGINT REFERENCES "user"(id),
    fk_updated_by_user_id BIGINT REFERENCES "user"(id),
    name VARCHAR(255) NOT NULL,
    country VARCHAR(255) NOT NULL,
    created_date TIMESTAMP NOT NULL DEFAULT now(),
    updated_date TIMESTAMP
);

CREATE TABLE notification (
    id BIGSERIAL PRIMARY KEY,
    fk_created_by_user_id BIGINT REFERENCES "user"(id),
    message TEXT,
    first_driver VARCHAR(255),
    second_driver VARCHAR(255),
    third_driver VARCHAR(255),
    license_points INTEGER,
    created_date TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE driver (
    id BIGSERIAL PRIMARY KEY,
    fk_created_by_user_id BIGINT REFERENCES "user"(id),
    fk_updated_by_user_id BIGINT REFERENCES "user"(id),
    name VARCHAR(255) NOT NULL,
    race_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    created_date TIMESTAMP NOT NULL DEFAULT now(),
    updated_date TIMESTAMP
)