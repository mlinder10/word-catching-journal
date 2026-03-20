-- name: GetProfile :one
select
  u.id,
  u.email,
  u.username,
  u.displayname,
  u.image_url,
  u.color,
  (select count(*) from post where user_id = u.id) as post_count,
  (select count(*) from user_follows f1 where f1.followee_id = u.id) as following_count,
  (select count(*) from user_follows f2 where f2.follower_id = u.id) as follower_count,
  coalesce((select 1 from user_follows f3 where f3.follower_id = sqlc.arg(user_id) and f3.followee_id = u.id), 0) as is_following
from profile u
where u.id = sqlc.arg(profile_id);

-- name: FollowUser :exec
insert into user_follows
  (follower_id, followee_id)
values
  (sqlc.arg(follower_id), sqlc.arg(followee_id));

-- name: UnfollowUser :exec
delete from user_follows
where
  follower_id = sqlc.arg(follower_id)
  and followee_id = sqlc.arg(followee_id);

-- name: GetFollowing :many
select
  id,
  email,
  username,
  displayname,
  image_url,
  color
from profile
where id in (select followee_id from user_follows where follower_id = sqlc.arg(user_id))
limit sqlc.arg(limit)
offset sqlc.arg(offset);

-- name: GetFollowingCount :one
select count(*) from user_follows where follower_id = sqlc.arg(user_id);

-- name: GetFollowers :many
select
  id,
  email,
  username,
  displayname,
  image_url,
  color
from profile
where id in (select follower_id from user_follows where followee_id = sqlc.arg(user_id))
limit sqlc.arg(limit)
offset sqlc.arg(offset);

-- name: GetFollowersCount :one
select count(*) from user_follows where followee_id = sqlc.arg(user_id);