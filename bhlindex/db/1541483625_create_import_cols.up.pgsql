CREATE TABLE import_cols (
    id SERIAL NOT NULL,
    name varchar(255) NOT NULL
);

-- Start at 1 because I think I used 0 as an error value?
-- ALTER SEQUENCE import_cols_id_seq RESTART WITH 1;
