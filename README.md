# go-project-modular

## Steps to create a local environment

1. Download Go: https://go.dev/dl/.
2. [Windows only]: If on Windows make sure to download direct to the C:// drive.
3. [Windows only]: Set the CGO_ENABLED environment variable: In your PowerShell window, run the following command: `$env:CGO_ENABLED = "1"`
4. [Windows only]: Download [TDM-GCC](https://jmeubank.github.io/tdm-gcc/download/) as a C compiler
5. [Windows only]: Set TDM-GCC as a new Environment Variable in Windows Control Panel. Add `C:\TDM-GCC` to the `PATH` variable
6. Run `go build`
7. Run `go run main.go`