CREATE KEYSPACE skeezy WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

use skeezy;

CREATE TABLE posts (
	id      uuid,
	body    varchar,
	dttm    timestamp,
	authors list<varchar>,
	tags    list<varchar>,
	PRIMARY KEY(id)
);

CREATE TABLE comments (
	id      uuid,
	post_id uuid,
	dttm    timestamp,
	author  varchar,
	email   varchar,
	body    varchar,
	PRIMARY KEY(id, post_id)
);
