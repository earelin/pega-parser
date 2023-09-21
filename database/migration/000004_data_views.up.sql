ALTER TABLE candidatura
    ADD COLUMN cabeceira_estatal INT UNSIGNED NOT NULL,
    ADD COLUMN cabeceira_autonomica INT UNSIGNED NOT NULL,
    ADD COLUMN cabeceira_provincial INT UNSIGNED NOT NULL;
