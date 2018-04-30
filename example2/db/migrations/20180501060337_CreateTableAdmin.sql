
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table Admins (
  id    serial PRIMARY KEY,
  name  varchar(120),
  age   integer,
  isman boolean
);

CREATE TABLE Careers (
  id    serial PRIMARY KEY,
  fromdate  date,
  todate    date,
  description text,
  admin_id integer REFERENCES admins (id) on update cascade on delete cascade
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE Admins cascade;
DROP TABLE Careers;
