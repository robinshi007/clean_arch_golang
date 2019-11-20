BEGIN;
  ALTER TABLE user_profiles DROP CONSTRAINT user_profiles_uid_fkey;
  ALTER TABLE user_accounts DROP CONSTRAINT user_accounts_uid_fkey;
  DROP TABLE IF EXISTS user_profiles;
  DROP TABLE IF EXISTS user_accounts;
COMMIT;
