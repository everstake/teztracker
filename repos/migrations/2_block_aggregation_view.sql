CREATE OR REPLACE VIEW public.block_aggregation_view
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
    sum(operations.isdouble_baking_evidence) AS double_baking_evidences
   FROM operations_for_counters operations
  GROUP BY operations.block_level;



