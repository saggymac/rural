CREATE TABLE IF NOT EXISTS messages (
	id integer not null primary key,
	msg text,
	appId text,
	sent integer CURRENT_TIMESTAMP
);