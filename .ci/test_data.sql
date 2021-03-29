DO
$$
    DECLARE
        album_id   INTEGER;
        artist_id  INTEGER;
        dc_user_id INTEGER;
    BEGIN
        INSERT INTO album(name) VALUES ('Never Give Up') RETURNING id INTO album_id;

        INSERT INTO artist(name) VALUES ('Veela') RETURNING id INTO artist_id;

        INSERT INTO dc_user(name) VALUES ('Albert') RETURNING id INTO dc_user_id;

        INSERT INTO song(name, artist_id, album_id, genre, description, dc_user_id)
        VALUES ('Got Me Thinking', artist_id, album_id, 'Art or something', 'well...', dc_user_id);
    END;
$$ LANGUAGE plpgsql;