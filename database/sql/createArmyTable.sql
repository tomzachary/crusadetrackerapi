CREATE TABLE army (
    id          SERIAL  PRIMARY KEY,
    title       TEXT        NOT NULL,
    description TEXT,
    createdAt   date        NOT NULL,
    modifiedAt  date,
    userId      int         NOT NULL
)