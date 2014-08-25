package gogitlab

import (
	"encoding/json"
	"fmt"
	"net/url"
  "strconv"
  "log"
)

const (
	users_url        = "/users"     // Get users list
	user_url         = "/users/:id" // Get a single user.
	current_user_url = "/user"      // Get current user
)

type User struct {
	Id             int    `json:"id,omitempty"`
	Username       string `json:"username,omitempty"`
	Email          string `json:"email,omitempty"`
	Name           string `json:"name,omitempty"`
	State          string `json:"state,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	Bio            string `json:"bio,omitempty"`
	Skype          string `json:"skype,omitempty"`
	LinkedIn       string `json:"linkedin,omitempty"`
	Twitter        string `json:"twitter,omitempty"`
	ExternUid      string `json:"extern_uid,omitempty"`
	Provider       string `json:"provider,omitempty"`
	ThemeId        int    `json:"theme_id,omitempty"`
	ColorSchemeId  int    `json:"color_scheme_id,color_scheme_id"`
  Password       string `json:"password"`
  Admin          bool   `json:"admin"`
  CreateGroup    bool   `json:"can_create_group"`
}


func (g *Gitlab) Users() ([]*User, error) {

	url := g.ResourceUrl(users_url, nil)

	var users []*User

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &users)
	}

	return users, err
}

/*
Get a single user.

    GET /users/:id

Parameters:

    id The ID of a user

Usage:

	user, err := gitlab.User("your_user_id")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v\n", user)
*/
func (g *Gitlab) User(id string) (*User, error) {

	url := g.ResourceUrl(user_url, map[string]string{":id": id})

	user := new(User)

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &user)
	}

	return user, err
}

func (g *Gitlab) DeleteUser(id string) error {
	url := g.ResourceUrl(user_url, map[string]string{":id": id})
	var err error
	_, err = g.buildAndExecRequest("DELETE", url, nil)
	return err
}

func (g *Gitlab) CurrentUser() (User, error) {
	url := g.ResourceUrl(current_user_url, nil)
	var user User

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err == nil {
		err = json.Unmarshal(contents, &user)
	}

	return user, err
}

/*
Create a new user

  POST /users

Parameters:
  email The email address of the new user
  password The password of the new user
  username The username of the new user
  name The name of the new user
  
  admin True/False new user is admin?
  can_create_group True/False new user can create groups

*/

func (g *Gitlab) AddUser(u User)  error {
	path := g.ResourceUrl(users_url, nil)

	var err error
	v := url.Values{}
	v.Set("email", u.Email)
	v.Set("name", u.Name)
	v.Set("username", u.Username)
  v.Set("admin", strconv.FormatBool(u.Admin))
  v.Set("password", u.Password)
  v.Set("can_create_group", strconv.FormatBool(u.CreateGroup))
	body := v.Encode()

  log.Printf("Request body: %s", body)
	_, err = g.buildAndExecRequest("POST", path, []byte(body))
  if err != nil {
    fmt.Printf("There was an error\n\n %s", err)
  	return err
  }
  return nil
}

