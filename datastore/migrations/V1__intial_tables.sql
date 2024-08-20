CREATE EXTENSION IF NOT EXISTS vector;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS adventure (
  game_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  adventure_id UUID NOT NULL,
  lore_id UUID NOT NULL,
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
  character_id UUID NOT NULL,
  name VARCHAR(255)
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
  lore_content_id UUID NOT NULL,
  content TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS lore_embedding (
  lore_id UUID NOT NULL,
  lore_content_id UUID NOT NULL,
  embedding VECTOR(768)
);
