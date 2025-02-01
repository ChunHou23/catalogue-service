CREATE TABLE IF NOT EXISTS cart (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    updated_by VARCHAR(255)
);

CREATE INDEX idx_cart_id ON cart(id);
CREATE INDEX idx_cart_active ON cart(active);
