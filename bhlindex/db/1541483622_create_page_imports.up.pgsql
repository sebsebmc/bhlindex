CREATE TABLE page_imports (
	page_id varchar(255) UNIQUE REFERENCES pages
);
