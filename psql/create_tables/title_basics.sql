CREATE TABLE IF NOT EXISTS title_basics (
       tconst VARCHAR(16),
       title_type VARCHAR(32),
       primary_title VARCHAR(1024),
       original_title VARCHAR(1024),
       is_adult BOOLEAN,
       start_year SMALLINT,
       end_year SMALLINT,
       runtime_minutes SMALLINT,
       genres VARCHAR(64)[]
);
