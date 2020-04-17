CREATE TABLE tezos.double_operation_evidences (
    operation_id integer PRIMARY KEY,
    doe_block_hash character varying NOT NULL,
    doe_type character varying NOT NULL,
    doe_block_level integer NOT NULL,
    doe_denounced_level integer NOT NULL,
    doe_offender character varying NOT NULL,
    doe_priority integer NOT NULL,
    doe_evidence_baker character varying NOT NULL,
    doe_baker_reward numeric NOT NULL,
    doe_lost_deposits numeric NOT NULL,
    doe_lost_rewards numeric NOT NULL,
    doe_lost_fees numeric NOT NULL,
    CONSTRAINT double_operation_evidences_block_foreign FOREIGN KEY (doe_block_level) REFERENCES tezos.blocks (level),
    CONSTRAINT double_operation_evidences_denounced_block_foreign FOREIGN KEY (doe_denounced_level) REFERENCES tezos.blocks (level),
    CONSTRAINT double_operation_evidences_operation_foreign FOREIGN KEY (operation_id) REFERENCES tezos.operations (operation_id)
);

INSERT INTO tezos.double_operation_evidences(operation_id, doe_block_hash, doe_block_level, doe_denounced_level,
                                             doe_offender, doe_priority, doe_evidence_baker, doe_baker_reward,
                                             doe_lost_deposits, doe_lost_rewards, doe_lost_fees, doe_type)
SELECT *, 'baking'
from tezos.double_baking_evidences;