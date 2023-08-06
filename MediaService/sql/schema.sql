CREATE TABLE songs (
    id SERIAL PRIMARY KEY,
    file_path TEXT NOT NULL,
    title TEXT NOT NULL,
    artist INT NOT NULL
);
