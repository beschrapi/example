CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    total DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id)
);
-- Add some sample data
INSERT INTO products (name, price)
VALUES ('Apple', 0.50);
INSERT INTO products (name, price)
VALUES ('Banana', 0.30);
INSERT INTO products (name, price)
VALUES ('Cherry', 0.75);