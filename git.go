package zze_goutils

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/golang/glog"
	"os"
)

//
//  GitPull
//  @Description: 从一个已有的 Git 目录执行 Pull 操作
//  @param gitDir 已存在的 Git 目录
//  @param username Git 用户名
//  @param password Git 用户密码
//  @return error 错误信息
//
func GitPull(gitDir, username, password string) (err error) {
	r, err := git.PlainOpen(gitDir)
	if err != nil {
		return
	}

	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		return
	}

	// Pull the latest changes from the origin remote and merge into the current branch
	err = w.Pull(&git.PullOptions{RemoteName: "origin", Force: true, Auth: &http.BasicAuth{
		Username: username,
		Password: password,
	}})
	if err != nil {
		return
	}

	// Print the latest commit that was just pulled
	ref, err := r.Head()
	if err != nil {
		return
	}

	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		return
	}
	glog.Info(commit)
	return nil
}

//
//  GitClone
//  @Description: 从给定的额 GitURL 克隆一个仓库到本地指定目录
//  @param gitUrl Git 仓库的 URL 地址
//  @param gitDir 本地目标目录
//  @param username Git 用户名
//  @param password Git 用户密码
//
func GitClone(gitUrl, gitDir, username, password string) (err error) {
	_, err = git.PlainClone(gitDir, false, &git.CloneOptions{
		URL:      gitUrl,
		Progress: os.Stdout,
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		},
		//ReferenceName: "master",
		//SingleBranch:  true,
	})

	if err != nil {
		return
	}
	return
}
