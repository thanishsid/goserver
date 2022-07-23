#!/bin/sh

echo "generating mock files..."

mockgen -package mockpostgres -destination mock/mockpostgres/mock_postgres.go github.com/thanishsid/goserver/infrastructure/postgres Querier

mockgen -package mockmailer -destination mock/mockmailer/mock_mailer.go github.com/thanishsid/goserver/internal/mailer Mailer

# Search Index Mocks

mockgen -package mocksearch -destination mock/mocksearch/mock_user_search.go github.com/thanishsid/goserver/infrastructure/search UserSearcher

# Repository Mocks

mockgen -package mockrepository -destination mock/mockrepository/mock_repository.go github.com/thanishsid/goserver/repository Repository,TxRepository

mockgen -package mockrepository -destination mock/mockrepository/mock_user_repository.go github.com/thanishsid/goserver/domain UserRepository

# Serices Mocks

mockgen -package mockservice -destination mock/mockservice/mock_service.go github.com/thanishsid/goserver/service Service

mockgen -package mockservice -destination mock/mockservice/mock_user_service.go github.com/thanishsid/goserver/domain UserService

echo "mock files generated successfully"
