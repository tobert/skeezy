CREATE KEYSPACE skeezy WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

use skeezy;

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
