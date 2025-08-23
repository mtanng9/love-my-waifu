-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY, 
    displayName text, 
    email text, 
    password text, 
    createdDate timestamp, 
    updatedDate timestamp
);

CREATE TABLE userSettings (
    id BIGSERIAL PRIMARY KEY, 
    userId BIGSERIAL REFERENCES users(id),
    createdDate timestamp, 
    updatedDate timestamp
);

CREATE TABLE waifus (
    id BIGSERIAL PRIMARY KEY, 
    name text, 
    hobby text, 
    bio text, 
    image text, 
    createdDate timestamp, 
    updatedDate timestamp
);

CREATE TABLE activeWaifus (
    id BIGSERIAL PRIMARY KEY, 
    userId BIGSERIAL REFERENCES users(id),
    waifuId BIGSERIAL REFERENCES waifus(id),
    createdDate timestamp, 
    updatedDate timestamp
);

CREATE TABLE activities (
    id BIGSERIAL PRIMARY KEY, 
    userId BIGSERIAL REFERENCES users(id),
    activity text, 
    dueDate timestamp, 
    completedDate timestamp, 
    createdDate timestamp, 
    updatedDate timestamp
);

CREATE TABLE favorites (
    id BIGSERIAL PRIMARY KEY, 
    userId BIGSERIAL REFERENCES users(id),
    waifuId BIGSERIAL REFERENCES waifus(id),
    createdDate timestamp, 
    updatedDate timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE userSettings;
DROP TABLE waifus;
DROP TABLE activeWaifus;
DROP TABLE activities;
DROP TABLE favorites;
-- +goose StatementEnd













