# SDK Testing

- Dev should create a branch from main and change the required package for testing.
- Make certain that the 'dev' environment and API tokens are not exposed.

## Installing branch as package

- In modules, enter the name of your branch to install it as a package.
- If you haven't already created a testing repo or branch, please fork the respective main repo and add a new branch with the appropriate Epic name and use the same below.

Production Repos :

- [pixelbin-js-admin](https://github.com/pixelbin-dev/pixelbin-js-admin)
- [pixelbin-python-sdk](https://github.com/pixelbin-dev/pixelbin-python-sdk)
- [pixelbin-go](https://github.com/pixelbin-dev/pixelbin-go)

### Node JS:

```
"@pixelbin/admin": "github:<username>/<repo-name>#<branch-name>"
```

Example :

```
"@pixelbin/admin": "github:rishabhkanojiyafynd/pixelbin-js-admin-temp#PIXB-2489"
```

### Python :

```
git+https://github.com/<username>/<repo-name>@<branch-name>#egg=pixelbin
```

Example :

```
git+https://github.com/rishabhkanojiyafynd/pixelbin-python-sdk.git@PIXB-2489#egg=pixelbin
```

### Go

```
go get -u "github.com/pixelbin-dev/pixelbin-go/v2@<branch-name>"
```

Example :

```
go get -u "github.com/pixelbin-dev/pixelbin-go/v2@PIXB-2491"
```

### Miscellaneous

- Replace `API_TOKEN` with the correct one.

- _No tokens or changes should be pushed._
