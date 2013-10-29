use skeezy;
insert into posts (id, authors, body, created, tags) values (ea4aba7d-9344-4d08-8ca5-873aa1214068, ['al'], 'first post!', 'now', []);
insert into comments (id, postId, parentId, created, author, email, body) values (5ee89e7c-a725-4cfd-9531-7c73b606b51d, ea4aba7d-9344-4d08-8ca5-873aa1214068, 5ee89e7c-a725-4cfd-9531-7c73b606b51d, 'now', 'al', 'al@example.com', 'first post first post!');
insert into comments (id, postId, parentId, created, author, email, body) values (fa97f6ea-da0c-4624-a29e-75bb5112571a, ea4aba7d-9344-4d08-8ca5-873aa1214068, 5ee89e7c-a725-4cfd-9531-7c73b606b51d, 'now', 'al', 'al@example.com', 'first post first post!');

