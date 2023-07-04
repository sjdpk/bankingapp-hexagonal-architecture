CREATE TABLE IF NOT EXISTS customers (
    customer_id bigserial PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    date_of_birth date NOT NULL,
    city VARCHAR(100) NOT NULL,
    zipcode VARCHAR(10) NOT NULL,
    status smallint NOT NULL DEFAULT 1
);
