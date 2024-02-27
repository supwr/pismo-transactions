CREATE TABLE IF NOT EXISTS sc_pismo.transactions (
    "id" BIGSERIAL NOT NULL,
    "account_id" BIGINT not null,
    "operation_type_id" BIGINT NOT NULL,
    "amount" DECIMAL(10,2) NOT NULL,
    "operation_date" TIMESTAMP NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NULL,
    "deleted_at" TIMESTAMP NULL,
    CONSTRAINT "PK_Transactions" PRIMARY KEY ("id")
);