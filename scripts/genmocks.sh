#!/bin/sh

echo "generating mock files..."

mockgen -package mockdb -destination mock/mockdb/mock_db.go github.com/thanishsid/goserver/infrastructure/db DB,Querier,Transactioner

mockgen -package mockmailer -destination mock/mockmailer/mock_mailer.go github.com/thanishsid/goserver/infrastructure/mailer Mailer

mockgen -package mocktoken -destination mock/mocktoken/mock_token.go github.com/thanishsid/goserver/infrastructure/tokenizer Tokenizer

# Serices Mocks

mockgen -package mockservice -destination mock/mockservice/mock_services.go github.com/thanishsid/goserver/domain \
    UserService,ImageService,SessionService

echo "mock files generated successfully"
