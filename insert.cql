use skeezy;

insert into users (id, username, email, created, updated) values (b56c89bc-6787-11e3-a635-fefd4312bbfd, 'AlTobey', 'tobert@gmail.com', 'now', a9137df6-6787-11e3-8349-fefd4312bbfd);
insert into users (id, username, email, created, updated) values (c0aaa2e6-6787-11e3-965b-fefd4312bbfd, 'Jepsen Networks', '1337@aphyr.com', 'now', c0fe5bca-6787-11e3-8652-fefd4312bbfd);

insert into posts (id, authors, body, created, tags) values (ea4aba7d-9344-4d08-8ca5-873aa1214068, ['al'], 'first post!', 'now', []);
insert into comments (id, postId, parentId, created, author, email, body) values (5ee89e7c-a725-4cfd-9531-7c73b606b51d, ea4aba7d-9344-4d08-8ca5-873aa1214068, 5ee89e7c-a725-4cfd-9531-7c73b606b51d, 'now', 'al', 'al@example.com', 'first post first post!');
insert into comments (id, postId, parentId, created, author, email, body) values (fa97f6ea-da0c-4624-a29e-75bb5112571a, ea4aba7d-9344-4d08-8ca5-873aa1214068, 5ee89e7c-a725-4cfd-9531-7c73b606b51d, 'now', 'al', 'al@example.com', 'first post first post!');

