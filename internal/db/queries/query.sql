-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1; 

-- name: CreateUser :one
INSERT INTO users (
    displayName, email, password
) VALUES (
    $1, $2, $3
) 
RETURNING *;

-- name: UpdateUser :one
UPDATE users
    set displayName = $2, 
    email = $3, 
    password = $4 
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec 
DELETE FROM users
WHERE id = $1; 

-- name: GetActiveWaifu :one
SELECT * FROM activeWaifus
WHERE userId = $1 LIMIT 1; 

-- name: CreateActiveWaifu :one
INSERT INTO activeWaifus (
    userId, waifuId
) VALUES (
    $1, $2
) 
RETURNING *;

-- name: UpdateActiveWaifus :one
UPDATE activeWaifus
    set userId = $2, 
    waifuId = $3
WHERE id = $1
RETURNING *;

-- name: DeleteActiveWaifus :exec 
DELETE FROM activeWaifus
WHERE id = $1; 

-- name: GetWaifu :one
SELECT * FROM waifus
WHERE id = $1 LIMIT 1; 

-- name: CreateWaifu :one
INSERT INTO waifus (
    name, hobby, bio, image
) VALUES (
    $1, $2, $3, $4
) 
RETURNING *;

-- name: UpdateWaifus :one
UPDATE waifus
    set name = $2, 
    hobby = $3, 
    bio = $4, 
    image = $5
WHERE id = $1
RETURNING *;

-- name: DeleteWaifus :exec 
DELETE FROM waifus
WHERE id = $1; 

-- name: GetActivities :one
SELECT * FROM activities
WHERE id = $1 LIMIT 1; 

-- name: CreateActivities :one
INSERT INTO activities (
    userId, activity, dueDate, completedDate
) VALUES (
    $1, $2, $3, $4
) 
RETURNING *;

-- name: UpdateActivities :one
UPDATE activities
    set userId = $2, 
    activity = $3, 
    dueDate = $4, 
    completedDate = $5
WHERE id = $1
RETURNING *;

-- name: DeleteActivities :exec 
DELETE FROM activities
WHERE id = $1; 


-- name: GetFavorites :one
SELECT * FROM favorites
WHERE id = $1 LIMIT 1; 

-- name: CreateFavorites :one
INSERT INTO favorites (
    userId, waifuId
) VALUES (
    $1, $2
) 
RETURNING *;

-- name: UpdateFavorites :one
UPDATE favorites
    set userId = $2, 
    waifuId = $3
WHERE id = $1
RETURNING *;

-- name: DeleteFavorites :exec 
DELETE FROM favorites
WHERE id = $1; 