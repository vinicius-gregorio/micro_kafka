-- init_data.sql


CREATE DATABASE IF NOT EXISTS wallet;

-- Balance microservice
CREATE DATABASE IF NOT EXISTS wallet_balance;

-- Switch to the wallet database
USE wallet;


DROP TABLE IF EXISTS clients;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS transactions;






-- Create the clients table
CREATE TABLE IF NOT EXISTS clients (
    id varchar(255),
    name varchar(255),
    email varchar(255),
    created_at date
);

-- Create the accounts table
CREATE TABLE IF NOT EXISTS accounts (
    id varchar(255),
    client_id varchar(255),
    balance float,
    created_at date
);

-- Create the transactions table
CREATE TABLE IF NOT EXISTS transactions (
    id varchar(255),
    account_id_from varchar(255),
    account_id_to varchar(255),
    amount float,
    created_at date
);


-- Insert initial clients
INSERT INTO clients (id, name, email, created_at) VALUES ('1', 'John Doe', 'john@example.com', CURRENT_TIMESTAMP);
INSERT INTO clients (id, name, email, created_at) VALUES ('2', 'Jane Smith', 'jane@example.com', CURRENT_TIMESTAMP);
INSERT INTO clients (id, name, email, created_at) VALUES ('3', 'Rocambole', 'rocambole@example.com', CURRENT_TIMESTAMP);



-- Insert initial accounts
INSERT INTO accounts (id, client_id, balance, created_at) VALUES ('1', '1', 100.00, CURRENT_TIMESTAMP);
INSERT INTO accounts (id, client_id, balance, created_at) VALUES ('2', '1', 200.00, CURRENT_TIMESTAMP);
INSERT INTO accounts (id, client_id, balance, created_at) VALUES ('3', '2', 300.00, CURRENT_TIMESTAMP);
INSERT INTO accounts (id, client_id, balance, created_at) VALUES ('4', '2', 100.00, CURRENT_TIMESTAMP);
INSERT INTO accounts (id, client_id, balance, created_at) VALUES ('5', '3', 200.00, CURRENT_TIMESTAMP);
INSERT INTO accounts (id, client_id, balance, created_at) VALUES ('6', '3', 300.00, CURRENT_TIMESTAMP);






-- Balance microservice

USE wallet_balance;

CREATE TABLE IF NOT EXISTS balances (
    account_id varchar(255),
    amount float
);

INSERT INTO balances (account_id, amount) VALUES ('1', 100.00);
INSERT INTO balances (account_id, amount) VALUES ('2', 200.00);
INSERT INTO balances (account_id, amount) VALUES ('3', 300.00);
INSERT INTO balances (account_id, amount) VALUES ('4', 100.00);
INSERT INTO balances (account_id, amount) VALUES ('5', 200.00);
INSERT INTO balances (account_id, amount) VALUES ('6', 300.00);