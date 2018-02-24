CREATE TABLE IF NOT EXISTS client__domains(
    name VARCHAR(255) PRIMARY KEY,
    client_id INT NOT NULL,
    INDEX client (client_id),
    CONSTRAINT fk_domain_client FOREIGN KEY (client_id) REFERENCES client__clients(id) ON DELETE CASCADE
)
