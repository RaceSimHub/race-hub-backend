ALTER TABLE "user" ADD COLUMN fk_driver_id BIGINT REFERENCES driver(id);

CREATE TABLE driver_link (
    id BIGSERIAL PRIMARY KEY,
    fk_user_id BIGINT REFERENCES "user"(id),
    fk_driver_id BIGINT REFERENCES driver(id),
    status VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);