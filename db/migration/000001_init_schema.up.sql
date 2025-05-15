CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE "frequency_enum" AS ENUM (
  'hourly',
  'daily'
);

-- SELECT  n.nspname AS enum_schema,
--         t.typname AS enum_name,
--         e.enumlabel AS enum_value
-- FROM    pg_type t JOIN
--         pg_enum e ON t.oid = e.enumtypid JOIN
--         pg_catalog.pg_namespace n ON n.oid = t.typnamespace
-- WHERE   t.typname = 'frequency_enum';


CREATE TABLE "subscriptions" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "city" varchar NOT NULL,
  "frequency" frequency_enum NOT NULL,
  "token" uuid NOT NULL,
  "confirmed" boolean NOT NULL DEFAULT false
);

-- docker exec -ti db bash
-- psql -U postgres -d weather_db
-- \dt
-- \d+ subscriptions