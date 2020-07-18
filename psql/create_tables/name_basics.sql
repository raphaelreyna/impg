CREATE TABLE IF NOT EXISTS name_basics (
       nconst VARCHAR(16),
       primary_name VARCHAR(256),
       birth_year SMALLINT,
       death_year SMALLINT,
       primary_profession VARCHAR(128)[],
       known_for_titles VARCHAR(16)[]
);
