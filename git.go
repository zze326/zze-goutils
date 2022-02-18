package zze_goutils

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"log"
	"os"
)

type GitRepository struct {
	Username string
	Password string
	GitUrl   string
	GitDir   string
}

//
//  NewGitRepository
//  @Description: 初始化一个 *GitRepository 实例
//  @param username Git 用户名
//  @param password Git 用户密码
//  @param gitUrl Git 仓库 Url
//  @param gitDir Git 本地目录
//  @return *GitRepository
//
func NewGitRepository(username, password, gitUrl, gitDir string) *GitRepository {
	return &GitRepository{
		Username: username,
		Password: password,
		GitUrl:   gitUrl,
		GitDir:   gitDir,
	}
}

//
//  GitPull
//  @Description: 从一个已有的 Git 目录执行 Pull 操作
//  @param gitDir 已存在的 Git 目录
//  @param username Git 用户名
//  @param password Git 用户密码
//  @return error 错误信息
//
func (gp *GitRepository) GitPull() (err error) {
	r, err := git.PlainOpen(gp.GitDir)
	if err != nil {
		return
	}

	w, err := r.Worktree()
	if err != nil {
		return
	}

	err = w.Pull(&git.PullOptions{RemoteName: "origin", Force: true, Auth: &http.BasicAuth{
		Username: gp.Username,
		Password: gp.Password,
	}})

	if err != nil {
		return
	}

	ref, err := r.Head()
	if err != nil {
		return
	}

	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		return
	}
	log.Println(commit)
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
func (gp *GitRepository) GitClone() (err error) {
	_, err = git.PlainClone(gp.GitDir, false, &git.CloneOptions{
		URL:      gp.GitUrl,
		Progress: os.Stdout,
		Auth: &http.BasicAuth{
			Username: gp.Username,
			Password: gp.Password,
		},
		//ReferenceName: "master",
		//SingleBranch:  true,
	})

	if err != nil {
		return
	}
	return
}
