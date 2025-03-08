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
    license VARCHAR(50),
    number INT,
    secondary_number INT,
    neighborhood VARCHAR(100),
    state VARCHAR(10),
    city VARCHAR(100),
    cep VARCHAR(20),
    address VARCHAR(255),
    address_number VARCHAR(20),
    phone VARCHAR(50),
    secondary_phone VARCHAR(50),
    country VARCHAR(50),
    email VARCHAR(255) NOT NULL,
    secondary_email VARCHAR(255),
    team VARCHAR(255),
    id_iracing VARCHAR(50),
    id_steam VARCHAR(50),
    instagram VARCHAR(255),
    facebook VARCHAR(255),
    twitch VARCHAR(255),
    photo_url VARCHAR(255),
    irating_sports_car INT,
    irating_oval INT,
    irating_formula_car INT,
    created_date TIMESTAMP NOT NULL DEFAULT now(),
    updated_date TIMESTAMP NOT NULL DEFAULT now()
);