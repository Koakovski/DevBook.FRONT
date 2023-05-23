package model

import (
	"devbook-front/src/infra/config"
	request "devbook-front/src/infra/requests"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	Id           uint64        `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	NickName     string        `json:"nickName,omitempty"`
	Email        string        `json:"email,omitempty"`
	Password     string        `json:"password,omitempty"`
	CreatedAt    time.Time     `json:"createdAt,omitempty"`
	Followers    []User        `json:"followers"`
	Following    []User        `json:"following"`
	Publications []Publication `json:"publications"`
}

func GetCompleteUser(userId uint64, r *http.Request) (User, error) {
	userDataChannel := make(chan User)
	userFollowersChannel := make(chan []User)
	userFollowingChannel := make(chan []User)
	userPublicationsChannel := make(chan []Publication)

	go getUserData(userDataChannel, userId, r)
	go getUserFollowersData(userFollowersChannel, userId, r)
	go getUserFollowingData(userFollowingChannel, userId, r)
	go getUserPublicationsData(userPublicationsChannel, userId, r)

	var (
		user             User
		userFollowers    []User
		userFollowing    []User
		userPublications []Publication
	)

	for i := 0; i < 4; i++ {
		select {
		case userFetched := <-userDataChannel:
			if userFetched.Id == 0 {
				return user, errors.New("erro em buscar usuário")
			}
			user = userFetched

		case userFollowersFetched := <-userFollowersChannel:
			if userFollowersFetched == nil {
				return user, errors.New("erro em buscar seguidores")
			}

			userFollowers = userFollowersFetched

		case userFollowingFetched := <-userFollowingChannel:
			if userFollowingFetched == nil {
				return user, errors.New("erro em buscar usuários que usuário está seguindo")
			}
			userFollowing = userFollowingFetched

		case userPublicationsFetched := <-userPublicationsChannel:
			if userPublicationsFetched == nil {
				return user, errors.New("erro em buscar publicações")
			}
			userPublications = userPublicationsFetched
		}
	}

	user.Followers = userFollowers
	user.Following = userFollowing
	user.Publications = userPublications

	return user, nil
}

func getUserData(channel chan<- User, userId uint64, r *http.Request) {
	var user User

	url := fmt.Sprintf("%s/user/%d", config.ApiUrl, userId)
	response, err := request.RequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- user
		return
	}
	defer response.Body.Close()

	_ = json.NewDecoder(response.Body).Decode(&user)

	channel <- user
}

func getUserFollowersData(channel chan<- []User, userId uint64, r *http.Request) {
	var users []User

	url := fmt.Sprintf("%s/user/%d/followers", config.ApiUrl, userId)
	response, err := request.RequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- users
		return
	}
	defer response.Body.Close()

	_ = json.NewDecoder(response.Body).Decode(&users)

	channel <- users
}

func getUserFollowingData(channel chan<- []User, userId uint64, r *http.Request) {
	var users []User

	url := fmt.Sprintf("%s/user/%d/following", config.ApiUrl, userId)
	response, err := request.RequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- users
		return
	}
	defer response.Body.Close()

	_ = json.NewDecoder(response.Body).Decode(&users)

	channel <- users
}

func getUserPublicationsData(channel chan<- []Publication, userId uint64, r *http.Request) {
	var publications []Publication

	url := fmt.Sprintf("%s/user/%d/publication", config.ApiUrl, userId)
	response, err := request.RequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- publications
		return
	}
	defer response.Body.Close()

	_ = json.NewDecoder(response.Body).Decode(&publications)

	channel <- publications
}
