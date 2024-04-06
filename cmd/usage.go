package main

const usage string = `mvnc - A client for Apache Maven HTTP repositories

Usage:
    mvnc -r=<url> -t=<token> <command>
    mvnc -r=<url> -u=<user> -p=<pass> <command>

Options:
    -r, -repo       Repo base url
    -t, -token      Auth token
    -u, -user       Auth username
    -p, -pass       Auth password
    -v, -verbose    More logs
    -h, -help       Show help
    -V, -version    Show version

Commands:
    ls - Show artifact ids, versions or artifacts
    Usage: mvnc ls <gid> [<aid>] [<ver>]

    dl - Fetch artifact(s) from remote server
    Usage: mvnc dl <gid> <aid> <ver> [<file>]

    up - Upload artifact to remote server
    Usage: mvnc up <gid> <aid> <ver> <file>

    rm - Remove artifact from remote server
    Usage: mvnc rm <gid> <aid> <ver> <file>
`
