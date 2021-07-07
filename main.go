package main

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	router.Post("/api/users", func(w http.ResponseWriter, r *http.Request) {
		// 注册
		// TODO: 参数校验
		// 1. 必填字段不能为空
		// 2. 邮箱校验
		data := new(UserLoginRequest)
		if err := render.Bind(r, data); err != nil {
			// err: unable to automatically decode the request content type
			fmt.Println(err.Error())
			// TODO: 错误处理
		}
		// TODO: 执行业务逻辑

		// TODO: 返回值封装
		userWithToken := UserWithToken{
			Email:    "email",
			Token:    "token",
			Username: "username",
			Bio:      nil,
			Image:    nil,
		}
		render.Render(w, r, &UserRegisterResponse{&userWithToken})
	})

	router.Post("/api/users/login", func(w http.ResponseWriter, r *http.Request) {
		// TODO: 1. 请求参数校验
		data := new(UserLoginRequest)
		if err := render.Bind(r, data); err != nil {
			fmt.Println(err.Error())
		}
		// TODO: 2. 业务逻辑
		// TODO: 3. 返回值
		w.Write([]byte("result"))
	})

	http.ListenAndServe(":8080", router)
}

// --
// 用户请求
// --

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterRequest struct {
	*User `json:"user"`
}

// Bind on UserRegisterRequest will run after the unmarshalling is complete
// its a good time to focus some post-processing after a decoding
func (u *UserRegisterRequest) Bind(r *http.Request) error {
	// u.User is nil if no User fields are sent in the request. Return an error to avoid a nil pointer dereference.
	if u.User == nil {
		return errors.New("missing required User fields")
	}
	fmt.Println("after decoding")
	return nil
}

type UserWithToken struct {
	Email    string
	Token    string
	Username string
	Bio      *string
	Image    *string
}

type UserRegisterResponse struct {
	*UserWithToken `json:"user"`
}

// Render manage response payloads
func (u *UserRegisterResponse) Render(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("render")
	return nil
}

type UserLoginParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginRequest struct {
	*UserLoginParam `json:"user"`
}

func (u UserLoginRequest) Bind(r *http.Request) error {
	if u.UserLoginParam == nil {
		fmt.Println("ff")
	}
	return nil
}
