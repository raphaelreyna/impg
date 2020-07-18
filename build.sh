go build -ldflags "-X 'main.CreateTablesSQL=$(cat psql/create_tables/*)'" -o impg .
