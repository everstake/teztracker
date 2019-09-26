generate_api:
	swagger generate server -t gen -f ./swagger/swagger.yml --exclude-main -A TezTracker

serve_ui:
	swagger serve swagger/swagger.yml

validate:
	swagger validate swagger/swagger.yml
