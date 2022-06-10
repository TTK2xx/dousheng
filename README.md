建video表

```
create table t_video
( id int not null primary key auto_increment,author int not null,play_url varchar(1000),cover_url varchar(1000),
 favorite_count bigint(64), comment_count bigint(64), 
 is_favorite bool, title varchar(1000),
 foreign key(author) references t_user(id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

