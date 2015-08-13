BEGIN;

/* remove any existing tables */
DROP TABLE IF EXISTS characters   CASCADE;
DROP TABLE IF EXISTS stages       CASCADE;
DROP TABLE IF EXISTS games        CASCADE;
DROP TABLE IF EXISTS players      CASCADE;
DROP TABLE IF EXISTS player_stats CASCADE;

/* recreate tables */
CREATE TABLE IF NOT EXISTS characters (
    id      serial PRIMARY KEY,
    name    text   NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS stages (
    id   serial PRIMARY KEY,
    name text   NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS games (
    id           serial    PRIMARY KEY,
    time         timestamp NOT NULL,
    stage_id     integer   NOT NULL,
    player_count integer   NOT NULL,
    CONSTRAINT stage_id_foreign_key
        FOREIGN KEY (stage_id) REFERENCES stages (id) MATCH SIMPLE,
    CONSTRAINT valid_player_count
        CHECK (player_count > 1 AND player_count <= 4)
);

CREATE TABLE IF NOT EXISTS players (
    id          serial                PRIMARY KEY,
    first_name  text                  NOT NULL,
    last_name   text                  NOT NULL,
    email       text                  NOT NULL UNIQUE,
    tag         text                  NOT NULL,
    is_disabled boolean DEFAULT FALSE NOT NULL
);

CREATE TABLE IF NOT EXISTS player_stats (
    id             serial                PRIMARY KEY,
    game_id        integer               NOT NULL,
    player_id      integer               NOT NULL,
    number         integer DEFAULT 1     NOT NULL,
    rank           integer               NOT NULL,
    character_id   integer               NOT NULL,
    is_random      boolean DEFAULT FALSE NOT NULL,
    kills          integer DEFAULT 0     NOT NULL,
    deaths         integer DEFAULT 0     NOT NULL,
    self_destructs integer DEFAULT 0     NOT NULL,
    CONSTRAINT game_id_foreign_key
        FOREIGN KEY (game_id) REFERENCES games (id) MATCH SIMPLE,
    CONSTRAINT player_id_foreign_key
        FOREIGN KEY (player_id) REFERENCES players (id) MATCH SIMPLE,
    CONSTRAINT character_id_foreign_key
        FOREIGN KEY (character_id) REFERENCES characters (id) MATCH SIMPLE,
    CONSTRAINT valid_number         CHECK (number >= 1 AND number <= 4),
    CONSTRAINT valid_rank           CHECK (rank >= 1 AND rank <= 4),
    CONSTRAINT valid_kills          CHECK (kills >= 0),
    CONSTRAINT valid_deaths         CHECK (deaths >= 0),
    CONSTRAINT valid_self_destructs CHECK (self_destructs >= 0)
);

/* seed data */
INSERT INTO characters (name) VALUES
    ('Bowser'),
    ('Captain Falcon'),
    ('Charizard'),
    ('Diddy Kong'),
    ('Donkey Kong'),
    ('Falco'),
    ('Fox'),
    ('Ganondorf'),
    ('Ice Climbers'),
    ('Ike'),
    ('Ivysaur'),
    ('Jigglypuff'),
    ('King Dedede'),
    ('Kirby'),
    ('Link'),
    ('Lucario'),
    ('Lucas'),
    ('Luigi'),
    ('Mario'),
    ('Marth'),
    ('Meta Knight'),
    ('Mewtwo'),
    ('Mr. Game & Watch'),
    ('Ness'),
    ('Olimar'),
    ('Peach'),
    ('Pikachu'),
    ('Pit'),
    ('R.O.B.'),
    ('Roy'),
    ('Samus'),
    ('Sheik'),
    ('Snake'),
    ('Sonic'),
    ('Squirtle'),
    ('Toon Link'),
    ('Wario'),
    ('Wolf'),
    ('Yoshi'),
    ('Zelda'),
    ('Zero Suit Samus');

INSERT INTO stages (name) VALUES
    ('Battlefield'),
    ('Big Blue'),
    ('Bowser''s Castle'),
    ('Brinstar'),
    ('Castle Siege'),
    ('Corneria'),
    ('Delfino''s Secret'),
    ('Distant Planet'),
    ('Dreamland'),
    ('Final Destination'),
    ('Flat Zone 2'),
    ('Fountain of Dreams'),
    ('Fourside'),
    ('Frigate Orpheon'),
    ('Green Hill Zone'),
    ('Halberd'),
    ('Hanenbow'),
    ('Hyrule Castle'),
    ('Infinite Glacier'),
    ('Jungle Japes'),
    ('Kongo Jungle'),
    ('Luigi''s Mansion'),
    ('Lylat Cruise'),
    ('Metal Cavern'),
    ('Norfair'),
    ('Onett'),
    ('Peach''s Castle'),
    ('Pictochat'),
    ('Pirate Ship'),
    ('Pokemon Stadium 2'),
    ('Port Town Aero Dive'),
    ('Rumble Falls'),
    ('Saffron City'),
    ('Shadow Moses Island'),
    ('Skyworld'),
    ('Smashville'),
    ('Spear Pillar'),
    ('Temple'),
    ('Training Room'),
    ('WarioWare, Inc.'),
    ('Yoshi''s Island'),
    ('Yoshi''s Story');

COMMIT;
