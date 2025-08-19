-- 0000_bootstrap.sql
\set ON_ERROR_STOP on

-- Роль можно создать внутри DO (транзакция допустима)
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'app') THEN
    CREATE ROLE app LOGIN PASSWORD 'secret' NOSUPERUSER NOCREATEROLE NOCREATEDB;
  END IF;
END
$$;

-- CREATE DATABASE нельзя внутри транзакции; делаем условно через \gexec
SELECT 'CREATE DATABASE wb_level0_db OWNER app'
WHERE NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'wb_level0_db')
\gexec

-- Настройки роли (идентификатор без кавычек; строк не нужно)
ALTER ROLE app SET search_path = app, public;