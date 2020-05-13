# TezTracker
  TezTracker is a open-source [Tezos](https://tezos.com) explorer based on Conseil indexer. Developed and supported by [Everstake](https://everstake.one) team. 
  
## Local deployment
### Environment variables
All project variables should be configured be environment.

Environment variables divides into 2 groups:
1. Conseil variables

	  Database Config:

		DB_HOST - default: db
		DB_USER - default: user
		DB_PASSWORD - default: password
		DB_DATABASE - default: conseil

    Tezos Node Config:

		XTZ_SCHEME - http or https, default : http
		XTZ_HOST - default: node
		XTZ_PORT - default 8732
    Also Conseil can be configured for Carthage testnet by adding CARTHAGENET_ prefix

2. TezTracker API server variables

        TEZTRACKER_PORT	- api port    
        TEZTRACKER_MAINNET_SQLCONNECTIONSTRING - raw Postgres connection string. example: postgresql://user:pass@127.0.0.1:5432/conseil?sslmode=disable
        TEZTRACKER_LOG_LEVEL - default: info
        TEZTRACKER_COUNTERINTERVALHOURS - update interval of chain counters. example: 2
        TEZTRACKER_FUTURERIGHTSINTERVALMINUTES - check interval of future baking/endorsement rights
        TEZTRACKER_SNAPSHOTCHECKINTERVALMINUTES - check interval of snapshots
        TEZTRACKER_DOUBLEBAKINGCHECKINTERVALMINUTES - check interval of double baking operations
        TEZTRACKER_DOUBLEENDORSEMENTCHECKINTERVALMINUTES - check interval of double endorsement operations

### Build and deploy Conseil
TezTracker depends on the [Conseil](https://github.com/Cryptonomic/Conseil) indexer. Follow the steps to deploy Conseil from our instructions or read through the [README](https://github.com/Cryptonomic/Conseil/blob/master/README.md) in the Conseil GitHub repository.   
Current explorer state work with [2020-january-release-19](https://github.com/Cryptonomic/Conseil/releases/tag/2020-january-release-19) Conseil releas, so use correct conseil.sql file for db init.

Clone the teztracker repository and cd into the cloned folder.

    git clone https://github.com/everstake/teztracker
    cd teztracker
    
Configure Conseil environment variables.

Build the Conseil docker image:

	docker-compose build conseil-lorre
  
Run the Conseil instance

	docker-compose up -d conseil-lorre
  
### Tezos Node
  We recommend use public archive nodes as `mainnet.tezos.org.ua` for save you time and disc space.
  
### Build and deploy TezTracker
 
   Clone the teztracker repository and cd into the cloned folder.

    git clone https://github.com/everstake/teztracker
    cd teztracker
    
   If support for multiple networks is needed add SQLCONNECTIONSTRING with required tezos net instead of MAINNET.
   
   Manualy exec sql migrations from `/repos/migrations` on PostgreSQL Conseil DB.
   
   From the root of the teztracker folder, execute the following command to build teztracker on Docker:
   
    docker-compose build api-server
    
   To start teztracker on localhost, execute:
   
    docker-compose up -d api-server
    
