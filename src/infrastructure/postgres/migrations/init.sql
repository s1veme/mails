CREATE TABLE users
(
    id            UUID PRIMARY KEY,
    name          TEXT NOT NULL,
    email         TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    is_active     BOOL      DEFAULT true,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE mails
(
    id         UUID PRIMARY KEY,
    subject    TEXT,
    body       TEXT,
    sender_id  UUID REFERENCES users (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE mail_recipients
(
    id              UUID PRIMARY KEY,
    mail_id         UUID REFERENCES mails (id),
    recipient_email TEXT NOT NULL,
    type            VARCHAR(3) CHECK (type IN ('TO', 'CC', 'BCC')),
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE folders
(
    id         UUID PRIMARY KEY,
    user_id    UUID REFERENCES users (id),
    name       TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

);

CREATE TABLE mail_folders
(
    id        UUID PRIMARY KEY,
    mail_id   UUID REFERENCES mails (id),
    folder_id UUID REFERENCES folders (id)
);

CREATE TABLE contacts
(
    id         UUID PRIMARY KEY,
    user_id    UUID REFERENCES users (id),
    email      TEXT NOT NULL,
    name       TEXT,
    frequency  INTEGER   DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
