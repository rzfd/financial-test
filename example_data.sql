-- Create User Table
CREATE TABLE IF NOT EXISTS users (
    id CHAR(36) PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20) UNIQUE NOT NULL,
    address TEXT,
    pin VARCHAR(10),
    balance BIGINT DEFAULT 0,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_date DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Create Transaction Table
CREATE TABLE IF NOT EXISTS transactions (
    id CHAR(36) PRIMARY KEY,
    status VARCHAR(20) NOT NULL,
    user_id CHAR(36) NOT NULL,
    transaction_type VARCHAR(10) NOT NULL,
    amount BIGINT NOT NULL,
    remarks TEXT,
    balance_before BIGINT NOT NULL,
    balance_after BIGINT NOT NULL,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Insert Sample Users
INSERT INTO users (id, first_name, last_name, phone_number, address, pin, balance) VALUES
('bc1c823e-b0fb-4b20-88c0-dff25e283252', 'John', 'Doe', '1234567890', '123 Elm St, Springfield', '1234', 500000),
('f3e8d65c-9e58-4f5b-a8e0-d1ebc4b85c25', 'Jane', 'Smith', '0987654321', '456 Oak St, Springfield', '5678', 1000000);

-- Insert Sample Transactions
INSERT INTO transactions (id, status, user_id, transaction_type, amount, remarks, balance_before, balance_after, created_date) VALUES
('a7d39cf6-44b6-41fc-b3e9-7b16df5321c5', 'SUCCESS', 'bc1c823e-b0fb-4b20-88c0-dff25e283252', 'DEBIT', 30000, 'Hadiah Ultah', 400000, 370000, '2021-04-01 22:23:20'),
('13bcb11c-111e-4a65-9afd-90a86a01cd21', 'SUCCESS', 'bc1c823e-b0fb-4b20-88c0-dff25e283252', 'DEBIT', 10000, 'Pulsa Telkomsel 100k', 500000, 490000, '2021-04-01 22:22:00'),
('201ddde1-f797-484b-b1a0-07d1190e790a', 'SUCCESS', 'f3e8d65c-9e58-4f5b-a8e0-d1ebc4b85c25', 'CREDIT', 500000, '', 0, 500000, '2021-04-01 22:21:21');
