module github.com/omniboost/go-fourth-pos

go 1.23.6

require (
	github.com/gocarina/gocsv v0.0.0-00010101000000-000000000000
	github.com/gorilla/schema v1.4.1
	github.com/hashicorp/go-multierror v1.1.1
	github.com/joefitzgerald/passwordcredentials v0.3.0
	golang.org/x/oauth2 v0.28.0
	gopkg.in/guregu/null.v3 v3.5.0
)

require github.com/hashicorp/errwrap v1.0.0 // indirect

replace github.com/gorilla/schema => github.com/omniboost/schema v1.1.1-0.20211111150515-2e872025e306

replace github.com/gocarina/gocsv => github.com/omniboost/gocsv v0.0.0-20230406101422-6445c2b15027
