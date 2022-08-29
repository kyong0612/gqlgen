package models

import "github.com/kyong0612/gqlgen/integration/remote_api"

type Viewer struct {
	User *remote_api.User
}
