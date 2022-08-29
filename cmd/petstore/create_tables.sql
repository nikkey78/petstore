DROP TABLE IF EXISTS category;
DROP TABLE IF EXISTS breed;
DROP TABLE IF EXISTS location;
DROP TABLE IF EXISTS pet;

CREATE TABLE IF NOT EXISTS category(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	category_name TEXT
);

CREATE TABLE IF NOT EXISTS breed(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	breed_name TEXT,
	category_id INTEGER,
	FOREIGN KEY (category_id) REFERENCES category(id)
);

CREATE TABLE IF NOT EXISTS location(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	location_name TEXT
);

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
