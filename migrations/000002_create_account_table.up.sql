CREATE TABLE IF NOT EXISTS accounts (
    account_id bigserial PRIMARY KEY,
    customer_id bigserial NOT NULL,
    opening_date date NOT NULL,
    account_type VARCHAR(10) NOT NULL,
    status smallint NOT NULL DEFAULT 1
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("customer_id");
