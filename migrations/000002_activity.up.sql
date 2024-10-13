CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    course_type VARCHAR(255),
    description VARCHAR(1024)
);

CREATE TABLE lectures (
    id SERIAL PRIMARY KEY,
    eventId INTEGER REFERENCES events(id),
    startAt TIMESTAMP,
    endAt TIMESTAMP
)