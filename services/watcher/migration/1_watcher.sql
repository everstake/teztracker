CREATE OR REPLACE FUNCTION tezos.notify_event() RETURNS TRIGGER AS $$
    DECLARE
        data json;
        notification json;
    BEGIN

        data = row_to_json(NEW);
        -- Contruct the notification as a JSON string.
        notification = json_build_object(
                          'table',TG_TABLE_NAME,
                          'action', TG_OP,
                          'data', data);

        -- Execute pg_notify(channel, notification)
        PERFORM pg_notify('events',notification::text);

        -- Result is ignored since this is an AFTER trigger
        RETURN NULL;
    END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER blocks_notify_event
AFTER INSERT ON tezos.blocks
    FOR EACH ROW EXECUTE PROCEDURE notify_event();