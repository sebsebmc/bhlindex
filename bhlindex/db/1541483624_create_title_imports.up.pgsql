CREATE TABLE title_imports (
	id serial REFERENCES titles,
	data jsonb NOT NULL
);
