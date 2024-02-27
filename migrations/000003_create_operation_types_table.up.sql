CREATE TABLE IF NOT EXISTS sc_pismo.operation_types (
    "id" BIGSERIAL NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NULL,
    "deleted_at" TIMESTAMP NULL,
    CONSTRAINT "PK_OperationTypes" PRIMARY KEY ("id")
);