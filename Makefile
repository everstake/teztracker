generate_api:
	swagger generate server -t gen -f ./swagger/swagger.yml --exclude-main -A TezTracker

generate_rpc_client:
	swagger generate client -t services/rpc_client -f ./services/rpc_client/client.yml -A tezosrpc

serve_ui:
	swagger serve swagger/swagger.yml

validate:
	swagger validate swagger/swagger.yml
