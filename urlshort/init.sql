-- Create the table
CREATE TABLE url_data (
    id SERIAL PRIMARY KEY,
    url TEXT NOT NULL,
    path TEXT NOT NULL
);

-- Insert sample data
INSERT INTO url_data (url, path) VALUES
    ('https://youtube.com', '/yt'),
    ('https://google.com', '/google');
