use real_world;
create table if not exists user
(
    id       varchar(255) primary key comment '用户 id',
    username varchar(255) unique comment '用户名',
    password varchar(255) comment '密码加盐散列后的密文',
    email    varchar(255) unique comment '邮箱地址',
    bio      text comment '介绍',
    image    varchar(255) comment '头像'
) comment = '用户表';
