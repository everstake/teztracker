create index blocks_meta_cycle_index
	on tezos.blocks (meta_cycle desc);

CREATE OR REPLACE FUNCTION tezos.blockInCycle(integer)
    RETURNS RECORD
    LANGUAGE plpgsql
AS $$
    DECLARE
    ret RECORD;

    BEGIN
    IF $1 >= 388 THEN
        SELECT 8192 :: integer, '30 seconds'::interval INTO ret;
      ELSE
        SELECT 4096 :: integer, '1 minute'::interval INTO ret;
    END IF;
    RETURN ret;
END $$;

CREATE OR REPLACE FUNCTION tezos.cycles(integer, integer)
   RETURNS TABLE(cycle integer, cycle_start TIMESTAMP WITHOUT TIME ZONE, cycle_end TIMESTAMP WITHOUT TIME ZONE)
   LANGUAGE plpgsql
AS $$
    declare
    start TIMESTAMP WITHOUT TIME ZONE;
    last_block_time TIMESTAMP WITHOUT TIME ZONE;
    cycle_position integer;
    blocksPerCycle integer;
    blockInterval interval;
    BEGIN

    SELECT blocks, interv into blocksPerCycle, blockInterval from tezos.blockInCycle($1) AS (blocks integer, interv interval);

    select min(timestamp), max(meta_cycle_position), max(timestamp) into start, cycle_position, last_block_time from tezos.blocks where meta_cycle = $1;
    cycle := $1;
    cycle_start := start;
    cycle_end := last_block_time + (blocksPerCycle - cycle_position -1) * blockInterval;
    return next;

    FOR l_counter IN $1+1..$2
    LOOP
      SELECT blocks, interv into blocksPerCycle, blockInterval from tezos.blockInCycle(l_counter) AS (blocks integer, interv interval);

      cycle := l_counter;

      cycle_start :=  cycle_end + blockInterval;
      cycle_end := cycle_start + (blocksPerCycle * blockInterval);  --interval '2 days 20 hours 16 minutes';
    return next;
    END LOOP;
END $$;

CREATE TABLE tezos.cycle_periods (
 cycle integer,
 cycle_start TIMESTAMP WITHOUT TIME ZONE,
 cycle_end TIMESTAMP WITHOUT TIME ZONE);

ALTER TABLE tezos.cycle_periods
	add constraint cycle_periods_pk
		primary key (cycle);

CREATE VIEW tezos.cycle_periods_view AS
select * from tezos.cycle_periods
UNION
select * from tezos.cycles( (select max(cycle) from tezos.cycle_periods), (select max(cycle) + 6 from tezos.cycle_periods) )
order by cycle desc;

CREATE VIEW tezos.snapshots_view AS
SELECT * FROM tezos.snapshots
 LEFT JOIN tezos.cycle_periods_view cp on snp_cycle = cp.cycle;

CREATE OR REPLACE FUNCTION tezos.cycle_periods()
RETURNS trigger
LANGUAGE plpgsql
AS
$$
BEGIN
  IF NEW.meta_cycle_position = 0 THEN
   INSERT INTO tezos.cycle_periods
    SELECT meta_cycle, min(timestamp), max(timestamp)
     FROM tezos.blocks
     where meta_cycle = NEW.meta_cycle - 1 group by meta_cycle;
  END IF;
  RETURN NEW;
END
$$;

CREATE TRIGGER cycle_periods_insert
  AFTER INSERT
  ON tezos.blocks
  FOR EACH ROW
EXECUTE PROCEDURE tezos.cycle_periods();

INSERT INTO tezos.cycle_periods
 SELECT meta_cycle, min(timestamp), max(timestamp)
 FROM tezos.blocks
 where meta_cycle >= 0 and meta_cycle <= (select max(meta_cycle) -1 FROM tezos.blocks) group by meta_cycle;