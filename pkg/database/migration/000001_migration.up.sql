CREATE TABLE notification (
    id BIGSERIAL PRIMARY KEY,
    message TEXT,
    first_driver VARCHAR(255),
    second_driver VARCHAR(255),
    third_driver VARCHAR(255),
    license_points INTEGER,
    created_date TIMESTAMP
);