CREATE SCHEMA taskschema

CREATE TABLE taskschema.users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE taskschema.tasks (
    id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(1000),
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    user_id UUID REFERENCES taskschema.users(id) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE taskschema.tags (
    id UUID PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE taskschema.tasks_tags (
    task_id UUID REFERENCES taskschema.tasks(id) NOT NULL,
    tag_id UUID REFERENCES taskschema.tags(id) NOT NULL,
    PRIMARY KEY (task_id, tag_id)
);
