CREATE TABLE IF NOT EXISTS employees (
    employee_id INT PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    date_of_birth VARCHAR(50),
    email VARCHAR(50) NOT NULL UNIQUE,
    phone_number VARCHAR(15),
    department VARCHAR(50) NOT NULL,
    job_title VARCHAR(100) NOT NULL,
    salary DECIMAL(10, 2) CHECK (salary >= 0.00),
    hire_date VARCHAR(50) NOT NULL,
    INDEX (email),
    INDEX (phone_number)
);