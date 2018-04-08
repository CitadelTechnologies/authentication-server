CREATE TABLE IF NOT EXISTS client__clients(
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) UNIQUE NOT NULL,
    token VARCHAR(255) NOT NULL,
    secret VARCHAR(255) NOT NULL,
    redirect_url VARCHAR(255) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME
);
