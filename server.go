package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

func RunServer() {
	app := iris.Default()
	{
		app.Get("/api/user/{username:string}", CachedHandler(UserAPI))
		app.Get("/api/repo/{username:string}/{repo:string}", CachedHandler(RepoAPI))
		app.Get("/api/contributor/{username:string}/{repo:string}", CachedHandler(ContributorAPI))
	}
	app.Listen(fmt.Sprintf(":%d", conf.Server.Port))
}

func UserAPI(ctx iris.Context) {
	username := ctx.Params().Get("username")
	data := AnalysisUser(username)
	EndBodyWithCache(ctx, data)
}

func RepoAPI(ctx iris.Context) {
	username, repo := ctx.Params().Get("username"), ctx.Params().Get("repo")
	data := AnalysisRepo(username, repo)
	EndBodyWithCache(ctx, data)
}

func ContributorAPI(ctx iris.Context) {
	username, repo := ctx.Params().Get("username"), ctx.Params().Get("repo")
	data := AnalysisContributor(username, repo)
	EndBodyWithCache(ctx, data)
}
