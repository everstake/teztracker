CREATE MATERIALIZED VIEW tezos.account_materialized_view
AS
select acc.*, alias as account_name
from (select account_id, min(asof) as created_at, max(asof) as last_active
      from tezos.accounts_history
      group by account_id) as acc
       left join tezos.known_addresses on acc.account_id = address;

CREATE INDEX account_created_time
  ON tezos.account_materialized_view (created_at);

//After sync

CREATE UNIQUE INDEX account_materialized_view_unique_index
  on tezos.account_materialized_view (account_id);

CREATE TRIGGER update_materialized_view
  AFTER INSERT
  ON tezos.blocks
  FOR EACH ROW
EXECUTE PROCEDURE refresh_account_materialized_view();

CREATE OR REPLACE FUNCTION refresh_account_materialized_view()
  RETURNS TRIGGER LANGUAGE plpgsql
  AS $$
  BEGIN
  REFRESH MATERIALIZED VIEW CONCURRENTLY tezos.account_materialized_view;
  RETURN NULL;
  END $$;

