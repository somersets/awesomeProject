package controller

type AppController struct {
	User   interface{ User }
	Auth   interface{ Auth }
	Tokens interface{ TokensAuth }
}
