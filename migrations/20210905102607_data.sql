-- +goose Up
INSERT INTO "user" (id, name)
VALUES (1, 'Test Ivan');

-- +goose Down
DELETE
FROM "user"
WHERE id = 1;
