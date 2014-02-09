CREATE TABLE IF NOT EXISTS devices (
	id integer not null primary key,
	deviceId text not null,
	appId text not null,
	appVersion text not null,
	lastSeen integer DEFAULT CURRENT_TIMESTAMP,
	UNIQUE( deviceId, appId)
);

