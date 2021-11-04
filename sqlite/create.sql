create table user ( 
    userid integer(10) primary key,
    username varchar(150)
);

create table account (
    userid integer(10),
    accountid char(11) primary key,
    balance integer,
    foreign key(userid) references customer(userid)
);

create table history (
    accountid char(11),
    historyid char(36) primary key,
    deposit integer,
    withdraw integer,
    historytime timestamp default current_timestamp not null,
    foreign key(accountid) references account(accountid)
);