
		CREATE TABLE IF NOT EXISTS artists (
		    id SERIAL PRIMARY KEY,
		    NAME TEXT NOT NULL 
		);

		CREATE TABLE IF NOT EXISTS songs (
    	id SERIAL PRIMARY KEY,
   		file_path TEXT NOT NULL,
    	title TEXT NOT NULL,
    	artist INT NOT NULL, 
    	FOREIGN KEY (artist) REFERENCES artists(id));

