insert into users (nickname, first_name, last_name, about, email, password, image)
values ('nick', 'First', 'Last', 'about', 'mail', 'pass', 'image'),         -- id = 1
       ('nick2', 'First2', 'Last2', 'about2', 'mail2', 'pass2', 'image2');  -- id = 2

insert into groups (name)
values ('name1'),  -- id = 1
       ('name2');  -- id = 2

insert into posts (user_id, group_id, content, created, image)
values (1, 2, 'content', '2004-10-19 10:23:54+02','image'),    -- id = 1
       (2, 2, 'content2', '2004-10-19 10:24:54+02','image2'),  -- id = 2
       (2, 2, 'content3', '2004-10-19 10:25:54+02','image3'),  -- id = 3
       (1, 1, 'content4', '2004-10-19 10:26:54+02','image4');  -- id = 4