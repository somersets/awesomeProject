package controller

type AppController struct {
	User      interface{ User }
	Auth      interface{ Auth }
	Tokens    interface{ TokensAuth }
	UserPhoto interface{ UserPhoto }
	UserChat  interface{ UserChat }
}
