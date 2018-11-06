CREATE TABLE page_imports (
	page_id varchar(255) REFERENCES pages,
	data jsonb NOT NULL
);

