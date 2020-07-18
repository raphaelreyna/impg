CREATE TABLE IF NOT EXISTS title_akas (
       title_id VARCHAR(16),
       ordering INTEGER,
       title VARCHAR(1024),
       region CHAR(16),
       language VARCHAR(32),
       types VARCHAR(32)[],
       attributes VARCHAR(64)[],
       is_original_title BOOLEAN
);
