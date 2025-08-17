-- 0000_bootstrap.sql

-- роль можно через DO (это в транзакции допустимо)
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = '${APP_DB_USER}') THEN
    CREATE ROLE '${APP_DB_USER}' LOGIN PASSWORD '${APP_DB_PASSWORD}' NOSUPERUSER NOCREATEROLE NOCREATEDB;
  END IF;
END$$;

-- ВАЖНО: CREATE DATABASE — отдельным выражением, без DO
CREATE DATABASE '${APP_DB_NAME}' OWNER '${APP_DB_USER}';

-- настройки роли
ALTER ROLE '${APP_DB_USER}' SET search_path = app, public;