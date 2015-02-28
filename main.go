package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Github API Json 2 Struct
type UserKeys []UserKey
type UserKey struct {
	ID  float64 `json:"id"`
	Key string  `json:"key"`
}

func main() {
	app := cli.NewApp()
	app.Usage = "fetch users public key from github.com"
	app.Version = "0.0.1"
	app.Author = "Joshua Mervine <joshua@mervine.net>"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "username, u",
			Usage: "target user's github username",
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "public key output file, default: {username}.pub",
		},
		cli.BoolFlag{
			Name:  "append,a",
			Usage: "append key to output file",
		},
	}

	app.Action = func(c *cli.Context) {
		usage := func() {
			cli.ShowAppHelp(c)
			os.Exit(1)
		}

		assert := func(e error) {
			if e != nil {
				fmt.Println(e)
				fmt.Println("---\n")
				os.Exit(1)
			}
		}

		username := c.String("username")

		if username == "" {
			usage()
		}

		output := c.String("output")
		append := c.Bool("append")

		// fetch keys
		endpoint := fmt.Sprintf("https://api.github.com/users/%s/keys", username)
		resp, err := http.Get(endpoint)
		defer resp.Body.Close()
		assert(err)

		body, err := ioutil.ReadAll(resp.Body)
		assert(err)

		var userKeys UserKeys
		err = json.Unmarshal(body, &userKeys)
		assert(err)

		keys := fmt.Sprintf("# git-my-key: github.com/%s\n#\ton: %s\n", username, time.Now())
		for i := range userKeys {
			keys += fmt.Sprintf("%s %s@api.github.com\n", userKeys[i].Key, username)
		}

		// print keys
		if output == "" {
			fmt.Printf(keys)
			os.Exit(0)
		}

		// write keys
		if !append {
			err = ioutil.WriteFile(output, []byte(keys), 0600)
			assert(err)
			fmt.Printf("# git-my-key: github.com/%s \n\t-> %s\n", username, output)
			os.Exit(0)
		}

		// append keys
		fd, err := os.OpenFile(output, os.O_RDWR|os.O_APPEND, 0600)
		assert(err)
		defer fd.Close()

		_, err = fd.WriteString(keys)
		assert(err)

		fd.Sync()
		fmt.Printf("# git-my-key: github.com/%s \n\t-> %s\n", username, output)
		os.Exit(0)
	}

	app.Run(os.Args)
}
