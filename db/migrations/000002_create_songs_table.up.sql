CREATE TABLE songs (
    id VARCHAR(50) PRIMARY KEY,
    title TEXT NOT NULL,
    year INTEGER NOT NULL,
    genre TEXT NOT NULL,
    performer TEXT NOT NULL,
    duration INTEGER,
    album_id VARCHAR(50),
    FOREIGN KEY (album_id) REFERENCES albums(id)
);