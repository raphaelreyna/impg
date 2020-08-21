impg:
	go build -ldflags "-X 'main.CreateTablesSQL=$(cat psql/create_tables/*.sql)'" -o impg
