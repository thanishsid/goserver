package mock

//go:generate moq -pkg mock -out db_mock.go ../infrastructure/db DB Querier Transactioner
//go:generate moq -pkg mock -out tokenizer_mock.go ../infrastructure/tokenizer Tokenizer
//go:generate moq -pkg mock -out mailer_mock.go ../infrastructure/mailer Mailer
//go:generate moq -pkg mock -out user_service_mock.go ../domain UserService
//go:generate moq -pkg mock -out session_service_mock.go ../domain SessionService
