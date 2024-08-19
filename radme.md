# Snippetbox

A simple & blazingly fast CRUD app made with Go; features include but not limited to: Authentication, Security best practices, and ease of use.

## Technologies

- Front-end: Server-side rendered HTML with templates using the standard package `html/template`
- Back-end: Built completely with Go with minimal reliance on third-party packages. A list of all the packages used could be found inside `go.mod` file at the root directory.

## Usage

To run this app locally, you need to have [Go](https://go.dev/doc/install) installed. \
Then, run these commands in your terminal:

```
git clone https://github.com/DrMorax/Snippetbox
```

Then install the required dependencies:

```
cd snippetbox && go mod tidy
```

Then run the app on your local machine:

```
go run ./cmd/web
```

After that, go to http://localhost:4000 Where the app should be running.
