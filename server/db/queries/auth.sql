-- name: GetUserByID :one
select * from profile where id = sqlc.arg(user_id);

-- name: GetUserByEmail :one
select * from profile where email = sqlc.arg(email);

-- name: GetUserByUsername :one
select * from profile where username = sqlc.arg(username);

-- name: DeleteUnverifiedAccount :exec
delete from profile where email_verified = false
and (email = sqlc.arg(email) or username = sqlc.arg(username));

-- name: CreateUser :one
insert into profile
  (email, username, displayname, color, password)
values
  (sqlc.arg(email), sqlc.arg(username), sqlc.arg(displayname), sqlc.arg(color), sqlc.arg(password))
returning *;

-- name: CreateVerificationRecord :exec
insert into email_verification
  (user_id, code)
values
  (sqlc.arg(user_id), sqlc.arg(code));

-- name: DeleteVerificationRecord :exec
delete from email_verification where user_id = sqlc.arg(user_id);

-- name: VerifyUser :one
update profile
set email_verified = true
where id = sqlc.arg(user_id)
returning *;

-- name: GetVerificationByCode :one
select * from email_verification where code = sqlc.arg(code);

-- name: UpdatePassword :exec
update profile set password = sqlc.arg(hashed_password) where id = sqlc.arg(user_id);

-- name: CreatePasswordResetRecord :one
insert into password_reset
  (user_email)
values
  (sqlc.arg(email))
returning *;

-- name: DeletePasswordResetRecord :exec
delete from password_reset where user_email = sqlc.arg(email);

-- name: GetPasswordResetByCode :one
select
  pr.code,
  p.id
from password_reset pr
join profile p on pr.user_email = p.email
where code = sqlc.arg(code);