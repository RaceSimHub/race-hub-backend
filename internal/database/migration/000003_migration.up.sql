ALTER TABLE "user" ADD COLUMN status VARCHAR(255);
ALTER TABLE "user" ADD COLUMN email_verification_token VARCHAR(255);
ALTER TABLE "user" ADD COLUMN email_verification_expires_at TIMESTAMP;

UPDATE "user" SET status = 'ACTIVE';

ALTER TABLE "user" ALTER COLUMN status SET NOT NULL;