/* helper functions */
BEGIN TRANSACTION ISOLATION LEVEL SERIALIZABLE;

/*returns the most recently inserted game's id */
CREATE OR REPLACE FUNCTION most_recent_game_id() RETURNS integer AS $$
DECLARE game_id integer;
BEGIN
    SELECT currval(pg_get_serial_sequence('games', 'id'))
    INTO game_id;
    RETURN game_id;
END;
$$ LANGUAGE plpgsql;

/* returns the most recently inserted player_stat's id */
CREATE OR REPLACE FUNCTION most_recent_player_stat_id() RETURNS integer AS $$
DECLARE player_stat_id integer;
BEGIN
    SELECT currval(pg_get_serial_sequence('games', 'id'))
    INTO player_stat_id;
    RETURN player_stat_id;
END;
$$ LANGUAGE plpgsql;

/* creates a new game and returns the id */
CREATE OR REPLACE FUNCTION create_game(stage_name text, player_count int)
RETURNS integer AS $$
DECLARE
	game_id  integer;
    stage_id integer;
BEGIN
    /* get stage id from name */
    SELECT INTO stage_id id FROM stages WHERE name=stage_name;

	INSERT INTO games (time, stage_id, player_count)
	VALUES (current_timestamp, stage_id, $2)
	RETURNING id INTO game_id;
	RETURN game_id;
END;
$$ LANGUAGE plpgsql;

/* creates a new player_stat and returns the id */
CREATE OR REPLACE FUNCTION create_player_stat(
	game_id        int,
	player_email   text,
	number         int,
	rank           int,
	character_name text,
	is_random      boolean,
	kills          int,
	deaths         int,
	self_destructs int
) RETURNS integer AS $$
DECLARE
	player_stat_id int;
	player_id      int;
	character_id   int;
BEGIN
	/* get player id from email */
	SELECT INTO player_id id FROM players WHERE email=player_email;

	/* get character id from name */
	SELECT INTO character_id id FROM characters WHERE name=character_name;

	INSERT INTO player_stats (
		game_id,
		player_id,
		number,
		rank,
		character_id,
		is_random,
		kills,
		deaths,
		self_destructs
	) VALUES ($1, player_id, $3, $4, character_id, $6, $7, $8, $9)
    RETURNING id INTO player_stat_id;

    RETURN player_stat_id;
END;
$$ LANGUAGE plpgsql;

COMMIT;

/* clear tables before adding seed data */
BEGIN TRANSACTION ISOLATION LEVEL SERIALIZABLE;
TRUNCATE games, players, player_stats CASCADE;
COMMIT;

/* seeded players */
BEGIN TRANSACTION ISOLATION LEVEL SERIALIZABLE;
INSERT INTO players (first_name, last_name, email, tag, is_disabled) VALUES
    ('Ben',   'Higginbotham', 'ben.higginbotham@hudl.com', 'ZACH',  false),
    ('Brett', 'Kercher',      'brett.kercher@hudl.com',    'BRETT', false),
    ('Ross',  'Bayer',        'ross.bayer@hudl.com',       'Ross',  false);
COMMIT;

/* create a new test game */
BEGIN TRANSACTION ISOLATION LEVEL SERIALIZABLE;
SELECT create_game('Final Destination', 2);
SELECT create_player_stat(most_recent_game_id(), 'ben.higginbotham@hudl.com', 1, 1, 'Ganondorf', false, 0, 0, 0);
SELECT create_player_stat(most_recent_game_id(), 'brett.kercher@hudl.com',    2, 2, 'Fox',       false, 0, 0, 0);
COMMIT;
