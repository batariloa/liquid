
		CREATE TABLE IF NOT EXISTS artists (
		    id SERIAL PRIMARY KEY,
		    name TEXT NOT NULL 
		);

		CREATE TABLE IF NOT EXISTS songs (
    	id SERIAL PRIMARY KEY,
   		file_path TEXT NOT NULL,
    	title TEXT NOT NULL,
    	artist INT NOT NULL,
		uploadedBy INT NOT NULL,
    	FOREIGN KEY (artist) REFERENCES artists(id),
		FOREIGN KEY (uploadedBy) REFERENCES users(id)
	);

		INSERT INTO artists (id, name) VALUES (5, 'Artist Example');

