BEGIN;

ALTER TABLE schemas RENAME schema TO auto;
ALTER TABLE schemas ADD COLUMN teleop JSONB NOT NULL DEFAULT '[]';

ALTER TABLE reports ADD COLUMN auto_name TEXT NOT NULL DEFAULT '';

COMMIT;