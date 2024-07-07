package types

type RegisterPayload struct {
  Name string `json:"name" validate:"required"`
  Email string `json:"email" validate:"required"`
  Password string `json:"password" validate:"required"`
}

type LoginPayload struct {
  Email string `json:"email" validate:"required"`
  Password string `json:"password" validate:"required"`
}
