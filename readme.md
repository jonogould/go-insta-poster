# go-insta-poster

A proof of concept Instagram photo-uploader using the unofficial [go-insta](https://github.com/ahmdrz/goinsta) library.

>NOTE: I do not suggest you use this for anything other than playing around. You may find your account banned, who knows. All I know is that I acccept no responsibility for anything that may occur as a result of using it.

### Installation

```sh
cd $GOPATH/src/github.com/jonogould
git clone git@github.com:github.com/jonogould/go-insta-poster.git
cd go-insta-poster
dep ensure
```

### Configuration

You have to configure some env variables in order to authenticate with your Instagram account.

```sh
cp .env.example .env
```

Add your Instagram `username` and `password` to the required vars in the `.env` file

### Run-it

Now that you are configured, you can post a photo to your account.

1. Add your photo to the `/example` dir
2. Change the values for `filename`, `caption`, `quality` and `filter` in the `main.go` file (lines `39-43`)
3. `go run main.go`

A successful post would result in `Result: ok`
