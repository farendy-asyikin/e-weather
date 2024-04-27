INSERT INTO users (name, email, password, role, is_active, created_at, updated_at)
VALUES
    ('admin', 'admin@gmail.com', '$2a$10$Ibi3s2t1XrKfFZ81Nz9Aw.eGHGj.lwS59.Xap62i.ZPvqfswKMNae', 'superuser', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)

-- plain text password = password