CREATE TABLE "user" (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_date TIMESTAMP
);

INSERT INTO "user" (email, name, password, created_date) VALUES ('admin@racesimhub.com','Admin','$2a$10$DIvRLcErlPMW2pPYj7qUr.2nNrHXSBd/BL/ky20ludw1CINuWSXU2', now());