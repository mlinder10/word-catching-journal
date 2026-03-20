-- name: GetDefinitions :many
select * from definition where word = sqlc.arg(word);

-- name: CreateDefinition :exec
insert into definition
  (word, content)
values
  (sqlc.arg(word), sqlc.arg(content));