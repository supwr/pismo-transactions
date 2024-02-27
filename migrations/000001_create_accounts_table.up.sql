CREATE TABLE IF NOT EXISTS sc_pismo.accounts (
    "id" BIGSERIAL NOT NULL,
    "document" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NULL,
    "deleted_at" TIMESTAMP NULL,
    CONSTRAINT "PK_Accounts" PRIMARY KEY ("id")
);