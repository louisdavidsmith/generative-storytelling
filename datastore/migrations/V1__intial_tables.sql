CREATE EXTENSION IF NOT EXISTS vector;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS adventure (
  adventure_id UUID PRIMARY KEY,
  adventure_name VARCHAR(255) NOT NULL,
  lore_id UUID NOT NULL,
  premise TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS game (
  game_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  adventure_id UUID NOT NULL,
  character_id UUIT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS narrative (
  adventure_id UUID NOT NULL,
  content_id UUID NOT NULL,
  role VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS game_character (
  game_id UUID NOT NULL,
  character_id UUID DEFAULT gen_random_uuid(),
  name VARCHAR(255) NOT NULL,
);

CREATE TABLE IF NOT EXISTS character_stat (
  game_id UUID NOT NULL,
  character_id UUID NOT NULL,
  characteristic VARCHAR(255) NOT NULL,
  value int NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS lore (
  lore_id UUID NOT NULL,
  description text NOT NULL,
  name VARCHAR(255) NOT NULL,
)

CREATE TABLE IF NOT EXISTS lore_content (
  lore_id UUID NOT NULL,
  lore_content_id UUID NOT NULL DEFAULT gen_random_uuid(),
  embedding VECTOR(768) NOT NULL,
  content text NOT NULL
);

CREATE TABLE IF NOT EXISTS item (
  game_id UUID NOT NULL,
  item_id UUID NOT NULL,
  description text NOT NULL,
)
