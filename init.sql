-- create a table
CREATE TABLE DefaultTable(
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  message VARCHAR(100)
);

-- add test data
INSERT INTO DefaultTable (message)
  VALUES ('message 1'),
  ('message 2');