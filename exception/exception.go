package exception

type(
    Exception struct {
        Code int `json:"code"`
        Message string `json:"message"`
    }
)

func New(code int, message string) *Exception {
    return &Exception{
        Code: code,
        Message: message,
    }
}
