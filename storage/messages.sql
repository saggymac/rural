create table IF NOT EXISTS messages (
	id integer not null primary key,
	msg text,
	sent integer CURRENT_TIMESTAMP
);