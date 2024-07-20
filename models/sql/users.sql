/* DDL for users table */
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL
);


/* DML for users table */
INSERT INTO users (email, password_hash)
VALUES
    ('a@a.com', 'a'),
    ('b@b.com', 'b'),
    ('c@c.com', 'c'),
    ('d@d.com', 'd'),
    ('e@e.com', 'e');