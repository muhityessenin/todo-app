CREATE TABLE tasks_table
(
    id        UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title     VARCHAR(255) NOT NULL,
    active_at date,
    status    BOOLEAN DEFAULT FALSE
);
