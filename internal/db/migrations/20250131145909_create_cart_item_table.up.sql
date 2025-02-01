CREATE TABLE IF NOT EXISTS cart_item (
    id UUID PRIMARY KEY,
    quantity INT NOT NULL,
    product_id UUID NOT NULL,
    cart_id UUID NOT NULL,
    deleted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    updated_by VARCHAR(255)
);

CREATE INDEX idx_cart_item_id ON cart_item(id);
CREATE INDEX idx_cart_item_cart_id ON cart_item(cart_id);
CREATE INDEX idx_cart_item_product_id ON cart_item(product_id);

