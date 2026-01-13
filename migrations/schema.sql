CREATE TABLE IF NOT EXISTS links (
    id TEXT PRIMARY KEY,
    original_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    visits INT DEFAULT 0
);

CREATE INDEX IF NOT EXISTS idx_original_url ON links(original_url);