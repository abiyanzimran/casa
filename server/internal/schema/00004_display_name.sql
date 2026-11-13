-- Red proof (throwaway): breaks expand-contract on purpose.

-- +goose Up
ALTER TABLE account ADD COLUMN display_name text NOT NULL;

-- +goose Down
ALTER TABLE account DROP COLUMN display_name;
