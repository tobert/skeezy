CREATE KEYSPACE skeezy WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

use skeezy;

CREATE TABLE users (
	id      uuid,
	username varchar,
	email    varchar,
	created timestamp,
	PRIMARY KEY(id)
);

CREATE INDEX userUsernameIndex ON users (username);
CREATE INDEX userEmailIndex ON users (email);

CREATE TABLE posts (
	id      uuid,
	body    varchar,
	created timestamp,
	authors list<varchar>,
	tags    list<varchar>,
	PRIMARY KEY(id)
);

CREATE TABLE comments (
	id       uuid,
	postId   uuid,
	parentId uuid,
	created  timestamp,
	author   varchar,
	email    varchar,
	body     varchar,
	PRIMARY KEY(id, postId)
);

CREATE INDEX commentsPostIndex ON comments (postid);
