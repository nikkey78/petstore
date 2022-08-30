DROP TABLE IF EXISTS pet;

CREATE TABLE IF NOT EXISTS pet(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	age REAL,
	image_url TEXT,
	description TEXT,
	category_id INTEGER,
	breed_id INTEGER ,
	location_id INTEGER, 
	FOREIGN KEY (category_id) REFERENCES category(id),
	FOREIGN KEY (breed_id)    REFERENCES breed(id),
	FOREIGN KEY (location_id) REFERENCES location(id)
);
