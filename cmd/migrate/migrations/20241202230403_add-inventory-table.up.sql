CREATE TABLE IF NOT EXISTS inventory (
    product_id INT UNSIGNED PRIMARY KEY,
    quantity INT NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);