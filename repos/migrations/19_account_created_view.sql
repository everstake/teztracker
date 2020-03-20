CREATE MATERIALIZED VIEW tezos.account_materialized_view
AS
select acc.*, baker_name as account_name
from (select account_id, min(asof) as created_at, max(asof) as last_active
      from tezos.accounts_history
      group by account_id) as acc
       left join tezos.public_bakers on acc.account_id = delegate;

CREATE TRIGGER update_materialized_view
  AFTER INSERT
  ON tezos.blocks
  FOR EACH ROW
EXECUTE PROCEDURE refresh_account_materialized_view();

CREATE OR REPLACE FUNCTION refresh_account_materialized_view()
  RETURNS TRIGGER LANGUAGE plpgsql
  AS $$
  BEGIN
  REFRESH MATERIALIZED VIEW tezos.account_materialized_view;
  RETURN NULL;
  END $$;

create index accounts_history_account_id_index
  on tezos.accounts_history (account_id);

create index accounts_account_id_acc_index
	on tezos.accounts (account_id)
where account_id like 'tz%';

create index accounts_account_id_kt_index
	on tezos.accounts (account_id)
where account_id like 'KT1%';