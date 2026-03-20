-- name: CreatePost :one
insert into post
  (word, definition, part_of_speech, pronunciation,
  synonyms, antonyms, example, visibility, user_id)
values
  (sqlc.arg(word), sqlc.arg(definition), sqlc.arg(part_of_speech),
  sqlc.arg(pronunciation), sqlc.arg(synonyms), sqlc.arg(antonyms),
  sqlc.arg(example), sqlc.arg(visibility), sqlc.arg(user_id))
returning *;

-- name: DeletePost :exec
delete from post where id = sqlc.arg(post_id);

-- name: GetRecentPosts :many
SELECT
  p.id, p.visibility, p.created_at, p.updated_at, p.user_id,
  p.word, p.definition, p.part_of_speech, p.pronunciation,
  p.synonyms, p.antonyms, p.example,
  u.username, u.displayname, u.image_url, u.color,
  COALESCE(lc.count, 0) AS like_count,
  COALESCE(bc.count, 0) AS bookmark_count,
  EXISTS (
    SELECT 1 FROM post_like pl
    WHERE pl.post_id = p.id AND pl.user_id = sqlc.arg(user_id)
  ) AS is_liked,
  EXISTS (
    SELECT 1 FROM post_bookmark pb
    WHERE pb.post_id = p.id AND pb.user_id = sqlc.arg(user_id)
  ) AS is_bookmarked
FROM post p
JOIN profile u ON p.user_id = u.id
LEFT JOIN (
    SELECT post_id, COUNT(*) as count 
    FROM post_like GROUP BY post_id
) lc ON lc.post_id = p.id
LEFT JOIN (
    SELECT post_id, COUNT(*) as count 
    FROM post_bookmark GROUP BY post_id
) bc ON bc.post_id = p.id
WHERE
  p.visibility = 'public'
  OR p.user_id = sqlc.arg(user_id)
ORDER BY p.created_at DESC
LIMIT sqlc.arg(limit) OFFSET sqlc.arg(offset);

-- name: GetFollowingPosts :many
SELECT
  p.id, p.visibility, p.created_at, p.updated_at, p.user_id,
  p.word, p.definition, p.part_of_speech, p.pronunciation,
  p.synonyms, p.antonyms, p.example,
  u.username, u.displayname, u.image_url, u.color,
  COALESCE(lc.count, 0) AS like_count,
  COALESCE(bc.count, 0) AS bookmark_count,
  EXISTS (
    SELECT 1 FROM post_like pl
    WHERE pl.post_id = p.id AND pl.user_id = sqlc.arg(user_id)
  ) AS is_liked,
  EXISTS (
    SELECT 1 FROM post_bookmark pb
    WHERE pb.post_id = p.id AND pb.user_id = sqlc.arg(user_id)
  ) AS is_bookmarked
FROM post p
JOIN profile u ON p.user_id = u.id
LEFT JOIN (
    SELECT post_id, COUNT(*) as count 
    FROM post_like GROUP BY post_id
) lc ON lc.post_id = p.id
LEFT JOIN (
    SELECT post_id, COUNT(*) as count 
    FROM post_bookmark GROUP BY post_id
) bc ON bc.post_id = p.id
WHERE
  p.visibility = 'public'
  AND p.user_id IN (
    SELECT followee_id FROM user_follows 
    WHERE follower_id = sqlc.arg(user_id)
  )
ORDER BY p.created_at DESC
LIMIT sqlc.arg(limit) OFFSET sqlc.arg(offset);

-- name: GetProfilePosts :many
SELECT
  p.id, p.visibility, p.created_at, p.updated_at, p.user_id,
  p.word, p.definition, p.part_of_speech, p.pronunciation,
  p.synonyms, p.antonyms, p.example,
  u.username, u.displayname, u.image_url, u.color,
  COALESCE(lc.count, 0) AS like_count,
  COALESCE(bc.count, 0) AS bookmark_count,
  EXISTS (
    SELECT 1 FROM post_like pl
    WHERE pl.post_id = p.id AND pl.user_id = sqlc.arg(user_id)
  ) AS is_liked,
  EXISTS (
    SELECT 1 FROM post_bookmark pb
    WHERE pb.post_id = p.id AND pb.user_id = sqlc.arg(user_id)
  ) AS is_bookmarked
FROM post p
JOIN profile u ON p.user_id = u.id
LEFT JOIN (
    SELECT post_id, COUNT(*) as count 
    FROM post_like GROUP BY post_id
) lc ON lc.post_id = p.id
LEFT JOIN (
    SELECT post_id, COUNT(*) as count 
    FROM post_bookmark GROUP BY post_id
) bc ON bc.post_id = p.id
WHERE
  (sqlc.arg(can_see_private) OR p.visibility = 'public')
  AND p.user_id = sqlc.arg(profile_id)
ORDER BY p.created_at DESC
LIMIT sqlc.arg(limit) OFFSET sqlc.arg(offset);

-- name: LikePost :exec
insert into post_like
  (post_id, user_id)
values
  (sqlc.arg(post_id), sqlc.arg(user_id));

-- name: UnlikePost :exec
delete from post_like
where
  post_id = sqlc.arg(post_id)
  and user_id = sqlc.arg(user_id);

-- name: BookmarkPost :exec
insert into post_bookmark
  (post_id, user_id)
values
  (sqlc.arg(post_id), sqlc.arg(user_id));

-- name: UnbookmarkPost :exec
delete from post_bookmark
where
  post_id = sqlc.arg(post_id)
  and user_id = sqlc.arg(user_id);

-- name: GetTotalPostCount :one
select count(*) from post;

-- name: GetTotalFollowingPostCount :one
select
  count(*)
from post
where 
  user_id in (select followee_id from user_follows where follower_id = sqlc.arg(user_id));

-- name: GetTotalProfilePostCount :one
select
  count(*)
from post
where
  user_id = sqlc.arg(profile_id)
  and (sqlc.arg(can_see_private) or visibility = 'public');