-- 0000_bootstrap.sql

-- роль можно через DO (это в транзакции допустимо)
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'app') THEN
    CREATE ROLE app LOGIN PASSWORD 'secret' NOSUPERUSER NOCREATEROLE NOCREATEDB;
  END IF;
END$$;

-- ВАЖНО: CREATE DATABASE — отдельным выражением, без DO
CREATE DATABASE wb_level0_db OWNER app;

-- настройки роли
ALTER ROLE app SET search_path = app, public;