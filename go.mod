module github.com/phirmware/kubectl-gui

go 1.16

require (
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
)

replace github.com/phirmware/kubectl-gui/handlers => ./kubectl-gui/handlers
