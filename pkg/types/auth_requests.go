package types

type RegisterPayload struct {
<<<<<<< HEAD:pkg/types/auth_requests.go
  Name string `json:"name" validate:"required"`
  Email string `json:"email" validate:"required"`
  Password string `json:"password" validate:"required"`
  Gender bool `json:"gender" validate:"required"`
=======
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
>>>>>>> 4f0aafb (added cors config):pkg/types/request.go
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

