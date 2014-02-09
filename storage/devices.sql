CREATE TABLE IF NOT EXISTS devices (
	id integer not null primary key,
	deviceId text not null,
	appId text not null,
	appVersion text not null
);


CREATE UNIQUE INDEX IF NOT EXISTS devAppIndex ON devices (deviceId,appId);