# Vexera/LogoGen
**Swaggy custom vexera logos in seconds**

---

#### Use and install
This can be a library or a seperate web application. For the library simply call `LogoGen.CreateLogo(...)`

The assets folder is required and has to be in the current working directory (where the app is ran from).

- To install the library simply `go get github.com/Vexera/LogoGen`.
- For the webapp you can run `go get github.com/Vexera/LogoGen/cmd/LogoGenWebApp`, and the binary will end up as `$GOPATH/bin/LogoGenWebApp`.
  - The webapp has a -listen parameter which can be used for chaning the listen address.