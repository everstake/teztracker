ALTER TABLE tezos.accounts_history
    ADD COLUMN  is_baker boolean NOT NULL DEFAULT false;

ALTER TABLE tezos.accounts 
    ADD COLUMN  is_baker boolean NOT NULL DEFAULT false;
