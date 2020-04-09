CREATE OR REPLACE FUNCTION tezos.baking_rewards()
 RETURNS trigger LANGUAGE plpgsql
AS $$
BEGIN
insert into tezos.baking_rewards
select delegate_value, NEW.snp_cycle, count(1), sum(balance)
from (select account_id, max(block_level) block_level
      from tezos.accounts_history
             left join tezos.snapshots on NEW.snp_cycle = snp_cycle
      where cycle <= (NEW.snp_cycle - 7)
        and block_level <= snp_block_level
      group by account_id) s
       left join tezos.accounts_history ah on s.account_id = ah.account_id and s.block_level = ah.block_level
where delegate_value is not null
  and balance > 0
group by delegate_value;
RETURN NEW;
END $$;

CREATE TRIGGER baker_rewards_insert
  AFTER INSERT
  ON tezos.snapshots
  FOR EACH ROW
EXECUTE PROCEDURE tezos.baking_rewards();
