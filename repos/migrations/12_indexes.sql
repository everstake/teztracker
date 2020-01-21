CREATE INDEX ix_opeations_transactions_source_operation_id
    ON tezos.operations USING btree
    (source, operation_id)
      WHERE kind='transaction';

CREATE INDEX ix_opeations_transactions_destination_operation_id
    ON tezos.operations USING btree
    (destination, operation_id)
      WHERE kind='transaction';

CREATE INDEX ix_opeations_endorsements_delegate_block_level
    ON tezos.operations USING btree
    (delegate, block_level)
      WHERE kind='endorsement';

CREATE INDEX ix_opeations_endorsements_operation_id
    ON tezos.operations USING btree
    (operation_id)
      WHERE kind='endorsement';

CREATE INDEX ix_opeations_delegations_operation_id
    ON tezos.operations USING btree
    (operation_id)
      WHERE kind='delegation';

CREATE INDEX ix_opeations_endorsements_operation_id
    ON tezos.operations USING btree
    (operation_id)
      WHERE kind='endorsement';

CREATE INDEX ix_accounts_balance
    ON tezos.accounts USING btree
    (balance);



Create table tezos.operation_counters(
    cnt_id SERIAL PRIMARY KEY,
    cnt_last_op_id int not null,
    cnt_operation_type varchar(100) NOT NULL,
    cnt_count bigint not null,
    cnt_created_at timestamp with time zone NULL DEFAULT NULL,
    CONSTRAINT operation_counters_last_op_foreign FOREIGN KEY (cnt_last_op_id) REFERENCES tezos.operations (operation_id)
);


CREATE TABLE tezos.future_baking_rights (
    level integer NOT NULL,
    delegate character varying NOT NULL,
    priority integer NOT NULL,
    estimated_time timestamp without time zone NOT NULL,
    PRIMARY KEY (level, priority)
);

CREATE INDEX future_baking_rights_delegate_idx 
    ON tezos.future_baking_rights USING btree (delegate);

alter table tezos.blocks ADD UNIQUE (level);

CREATE TABLE tezos.snapshots (
    snp_cycle integer PRIMARY KEY ,
    snp_block_level integer NOT NULL,
    snp_rolls integer NOT NULL,
    CONSTRAINT snapshots_block_foreign FOREIGN KEY (snp_block_level) REFERENCES tezos.blocks (level)
);

CREATE OR REPLACE VIEW tezos.operations_for_counters AS
select block_level, amount,fee, case when operations.kind='endorsement' then 1 else 0 end as isendorsement, case when operations.kind='proposals' then 1 else 0 end as isproposals, case when operations.kind='seed_nonce_revelation' then 1 else 0 end as isseed_nonce_revelation, case when operations.kind='delegation' then 1 else 0 end as isdelegation, case when operations.kind='transaction' then 1 else 0 end as istransaction, case when operations.kind='activate_account' then 1 else 0 end as isactivate_account, case when operations.kind='ballot' then 1 else 0 end as isballot, case when operations.kind='origination' then 1 else 0 end as isorigination, case when operations.kind='reveal' then 1 else 0 end as isreveal, case when operations.kind='double_baking_evidence' then 1 else 0 end as isdouble_baking_evidence, case when operations.kind='double_endorsement_evidence' then 1 else 0 end as isdouble_endorsement_evidence
  from tezos.operations;

CREATE OR REPLACE VIEW tezos.block_aggregation_view
    AS
    SELECT operations.block_level AS level,
    COALESCE(sum(operations.amount), 0::numeric) AS volume,
    COALESCE(sum(operations.fee), 0::numeric) AS fees,
    sum(operations.isendorsement) AS endorsements,
    sum(operations.isproposals) AS proposals,
    sum(operations.isseed_nonce_revelation) AS seed_nonce_revelations,
    sum(operations.isdelegation) AS delegations,
    sum(operations.istransaction) AS transactions,
    sum(operations.isactivate_account) AS activate_accounts,
    sum(operations.isballot) AS ballots,
    sum(operations.isorigination) AS originations,
    sum(operations.isreveal) AS reveals,
    sum(operations.isdouble_baking_evidence) AS double_baking_evidences,
    sum(operations.isdouble_endorsement_evidence) AS double_endorsement_evidences
   FROM tezos.operations_for_counters operations
  GROUP BY operations.block_level;

CREATE TABLE tezos.double_baking_evidences (
    operation_id integer PRIMARY KEY,
    dbe_block_hash character varying NOT NULL,
    dbe_block_level integer NOT NULL,
    dbe_denounced_level integer NOT NULL,
    dbe_offender character varying NOT NULL,
    dbe_priority integer NOT NULL,
    dbe_evidence_baker character varying NOT NULL,
    dbe_baker_reward numeric NOT NULL,
    dbe_lost_deposits numeric NOT NULL,
    dbe_lost_rewards numeric NOT NULL,
    dbe_lost_fees numeric NOT NULL,
    CONSTRAINT double_baking_evidences_block_foreign FOREIGN KEY (dbe_block_level) REFERENCES tezos.blocks (level),
    CONSTRAINT double_baking_evidences_denounced_block_foreign FOREIGN KEY (dbe_denounced_level) REFERENCES tezos.blocks (level),
    CONSTRAINT double_baking_evidences_operation_foreign FOREIGN KEY (operation_id) REFERENCES tezos.operations (operation_id)
);

