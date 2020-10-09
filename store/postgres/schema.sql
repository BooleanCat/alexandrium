CREATE TABLE author (
   id serial PRIMARY KEY,
   uuid uuid UNIQUE NOT NULL,
   name VARCHAR NOT NULL
);

INSERT INTO
    author(uuid, name)
VALUES
    ('ea1ff7d7-67cd-477c-8cb7-8756619e275d', 'Adrian Tchaikovsky');
