BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE theaters (
  token UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  number INT NOT NULL,
  description TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE movies (
  token UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name TEXT NOT NULL,
  director TEXT NOT NULL,
  duration INT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE sessions (
  token UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  movie_token UUID NOT NULL,
  thread_token UUID NOT NULL,
  session_datetime TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (movie_token) REFERENCES movies(token),
  FOREIGN KEY (thread_token) REFERENCES theaters(token)
);

INSERT INTO theaters (number, description) VALUES
(1, 'Sala 1 - VIP'),
(2, 'Sala 2 - 3D'),
(3, 'Sala 3 - IMAX'),
(4, 'Sala 4 - LUX'),
(5, 'Sala 5 - Standard');

INSERT INTO movies (name, director, duration) VALUES
('Echoes of the Past', 'Steven Spielberg', 95),
('Streets of Fury', 'Martin Scorsese', 100), 
('Duel at Dusk', 'Quentin Tarantino', 110),
('Time Rift', 'Christopher Nolan', 120), 
('Beyond the Horizon', 'Ridley Scott', 90), 
('Abyssal Quest', 'James Cameron', 135), 
('Whimsical Shadows', 'Tim Burton', 105),      
('Mind Game', 'David Fincher', 115),         
('The Quirky Hotel', 'Wes Anderson', 95),     
('Mystery of the Manor', 'Alfred Hitchcock', 110), 
('Golden Era', 'Francis Ford Coppola', 125),   
('Forgotten Realms', 'Peter Jackson', 140),    
('The Labyrinth of Solitude', 'Guillermo del Toro', 130), 
('Manhattan Memoirs', 'Woody Allen', 100), 
('Rhythm of the Street', 'Spike Lee', 120); 

INSERT INTO sessions (movie_token, thread_token, session_datetime) VALUES
((SELECT token FROM movies WHERE name = 'Echoes of the Past'), (SELECT token FROM theaters WHERE number = 1), '2024-06-21 10:00:00'),
((SELECT token FROM movies WHERE name = 'Streets of Fury'), (SELECT token FROM theaters WHERE number = 1), '2024-06-21 14:00:00'),
((SELECT token FROM movies WHERE name = 'Duel at Dusk'), (SELECT token FROM theaters WHERE number = 1), '2024-06-21 18:00:00'),
((SELECT token FROM movies WHERE name = 'Time Rift'), (SELECT token FROM theaters WHERE number = 2), '2024-06-21 10:30:00'),
((SELECT token FROM movies WHERE name = 'Beyond the Horizon'), (SELECT token FROM theaters WHERE number = 2), '2024-06-21 14:30:00'),
((SELECT token FROM movies WHERE name = 'Abyssal Quest'), (SELECT token FROM theaters WHERE number = 2), '2024-06-21 18:30:00'),
((SELECT token FROM movies WHERE name = 'Whimsical Shadows'), (SELECT token FROM theaters WHERE number = 3), '2024-06-21 11:00:00'),
((SELECT token FROM movies WHERE name = 'Mind Game'), (SELECT token FROM theaters WHERE number = 3), '2024-06-21 15:00:00'),
((SELECT token FROM movies WHERE name = 'The Quirky Hotel'), (SELECT token FROM theaters WHERE number = 3), '2024-06-21 19:00:00'),
((SELECT token FROM movies WHERE name = 'Mystery of the Manor'), (SELECT token FROM theaters WHERE number = 4), '2024-06-21 11:30:00'),
((SELECT token FROM movies WHERE name = 'Golden Era'), (SELECT token FROM theaters WHERE number = 4), '2024-06-21 15:30:00'),
((SELECT token FROM movies WHERE name = 'Forgotten Realms'), (SELECT token FROM theaters WHERE number = 4), '2024-06-21 19:30:00'),
((SELECT token FROM movies WHERE name = 'The Labyrinth of Solitude'), (SELECT token FROM theaters WHERE number = 5), '2024-06-21 12:00:00'),
((SELECT token FROM movies WHERE name = 'Manhattan Memoirs'), (SELECT token FROM theaters WHERE number = 5), '2024-06-21 16:00:00'),
((SELECT token FROM movies WHERE name = 'Rhythm of the Street'), (SELECT token FROM theaters WHERE number = 5), '2024-06-21 20:00:00');


COMMIT;
