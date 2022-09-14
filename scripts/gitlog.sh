#/bin/sh -e

git log -n 15 --pretty=format:'%h -%d %s (%cr) <%an>' --abbrev-commit |\
go run main.go |\
bat --style=changes -l gitlog
