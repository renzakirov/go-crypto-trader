CREATE TABLE IF NOT EXISTS operation_type(
  id serial not null PRIMARY key,
  operation_type_name VARCHAR (50) UNIQUE NOT NULL,
  description VARCHAR (300)
);

CREATE TABLE IF NOT EXISTS trade (
  trade_id serial not null PRIMARY key,
  id integer,
  operation_type_id integer not null REFERENCES operation_type(id),
  price numeric,
  quantity numeric,
  amount numeric,
  date timestamp without time zone
);

create index if not exists IdxTradeId on trade(id);
create index if not exists IdxTradeOperationType on trade(operation_type_id);
create index if not exists IdxTradeDate on trade(date);
