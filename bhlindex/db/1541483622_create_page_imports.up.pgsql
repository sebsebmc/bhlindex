CREATE TABLE page_imports (
	page_id varchar(255) REFERENCES pages,
	importData jsonb NOT NULL
);
