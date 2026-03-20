create table if not exists profile (
  id text primary key not null default (lower(hex(randomblob(16)))),
  email text unique not null,
  username text unique not null,
  displayname text not null,
  color text not null,
  image_url text,

  email_verified boolean default false not null,
  password text not null,
  access_token text,
  refresh_token text,

  created_at datetime default current_timestamp not null,
  updated_at datetime default current_timestamp not null
);

create table if not exists email_verification (
  code text primary key not null,
  user_id text not null,
  created_at datetime default current_timestamp not null,

  foreign key (user_id) references profile (id) on delete cascade
);

create table if not exists password_reset (
  code text primary key not null default (lower(hex(randomblob(16)))),
  user_email text not null,
  created_at datetime default current_timestamp not null,

  foreign key (user_email) references profile (email) on delete cascade
);

create table if not exists post (
  id text primary key not null default (lower(hex(randomblob(16)))),
  user_id text not null,
  visibility text not null,
  created_at datetime default current_timestamp not null,
  updated_at datetime default current_timestamp not null,

  word text not null,
  definition text not null,
  part_of_speech text not null,
  pronunciation text not null,
  example text not null,
  synonyms text not null,
  antonyms text not null,

  foreign key (user_id) references profile (id) on delete cascade
);

create table if not exists post_like (
  user_id text not null,
  post_id text not null,
  created_at datetime default current_timestamp not null,

  foreign key (user_id) references profile (id) on delete cascade,
  foreign key (post_id) references post (id) on delete cascade,
  primary key (user_id, post_id)
);

create table if not exists post_bookmark (
  user_id text not null,
  post_id text not null,
  created_at datetime default current_timestamp not null,

  foreign key (user_id) references profile (id) on delete cascade,
  foreign key (post_id) references post (id) on delete cascade,
  primary key (user_id, post_id)
);

create table if not exists user_follows (
  follower_id text not null,
  followee_id text not null,
  created_at datetime default current_timestamp not null,
  
  foreign key (follower_id) references profile (id) on delete cascade,
  foreign key (followee_id) references profile (id) on delete cascade,
  primary key (follower_id, followee_id)
);

create table if not exists definition (
  word text not null primary key,
  content text not null,
  created_at datetime default current_timestamp not null,
  updated_at datetime default current_timestamp not null
);