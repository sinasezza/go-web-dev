/* DDL for sessions table */
CREATE TABLE IF NOT EXISTS sessions (
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
    token_hash TEXT UNIQUE NOT NULL
);


/* DML for sessions table */
INSERT INTO sessions (user_id, token_hash)
VALUES
    (1, 'hashed_token'),
    (2, 'hashed_token');