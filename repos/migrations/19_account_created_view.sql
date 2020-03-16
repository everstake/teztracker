CREATE OR REPLACE VIEW tezos.account_create_time_view
AS
select account_id, min(asof) as asof
from tezos.accounts_history
group by account_id;

create index accounts_history_account_id_index
  on tezos.accounts_history (account_id);