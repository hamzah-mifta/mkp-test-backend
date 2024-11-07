CREATE TABLE IF NOT EXISTS schedules (
    id SERIAL PRIMARY KEY NOT NULL,
    movie_id INTEGER NOT NULL,
    theater_id INTEGER NOT NULL,
    show_date DATE NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (movie_id) REFERENCES movies(id),
    FOREIGN KEY (theater_id) REFERENCES theaters(id)
);