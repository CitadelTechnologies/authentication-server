CREATE TABLE IF NOT EXISTS user__users(
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at DATETIME,
    last_connected_at DATETIME
);
