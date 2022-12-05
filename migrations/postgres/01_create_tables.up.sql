CREATE TYPE "user_type_t" AS ENUM ('SUPERADMIN', 'ADMIN', 'USER');

CREATE TABLE IF NOT EXISTS "user" (
	"id" CHAR(36) PRIMARY KEY,
	"username" VARCHAR(255) UNIQUE NOT NULL,
	"password" VARCHAR(255) NOT NULL ,
	"user_type" user_type_t NOT NULL ,
	"created_at" TIMESTAMP DEFAULT now() NOT NULL,
	"updated_at" TIMESTAMP
);
