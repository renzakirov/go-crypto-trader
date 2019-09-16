CREATE TABLE IF NOT EXISTS operation_type(
  id serial not null PRIMARY key,
  operation_type_name VARCHAR (50) UNIQUE NOT NULL,
  description VARCHAR (300)
);

INSERT INTO operation_type(operation_type_name, description) VALUES ('sell', 'sell curency'), ('buy', 'buy curency');

CREATE TABLE IF NOT EXISTS trade (
  id serial not null PRIMARY key,
  trade_id integer UNIQUE not null,
  operation_type_id integer not null REFERENCES operation_type(id),
  price numeric,
  quantity numeric,
  amount numeric,
  date timestamp without time zone
);

create index if not exists IdxTradeId on trade(id);
create index if not exists IdxTradeOperationType on trade(operation_type_id);
create index if not exists IdxTradeDate on trade(date);
