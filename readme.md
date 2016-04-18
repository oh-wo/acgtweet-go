# Acgtweet

## Requirements
- [Postgress](http://postgresapp.com/)
- [Google app engine SDK](https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go).
- mysql 
 `brew install mysql`
 `unset TMPDIR`
  `mysql_install_db --verbose --user=`whoami` --basedir="$(brew --prefix mysql)" --datadir=/usr/local/var/mysql`
- To use `psql acgtweet` on the command line:  
  `echo 'export PATH="/Applications/Postgres.app/Contents/Versions/9.4/bin:$PATH"' >> ~/.bash_profile`

## Installation
1. All packages will automatically download (are already downloaded?) with the exception of the following:
 * `go get github.com/go-sql-driver/mysql`

## Development
1. Start mysql `mysql.server start`
1. To run the local environment `goapp serve hello.go`
 * This will start a local server running at `localhost:8080`
 * Will watch files and auto recompile.

## Deployment
1. To upload to google app engine: `appcfg.py -A acgtweet -V v1 update ./`
 (`goapp deploy` may also work if your pc is setup...)


## Mysql
 - Connect `mysql -u root -p` (password==password).
 - Create user       `CREATE USER 'user'@'localhost' IDENTIFIED BY 'password';`
                     `GRANT ALL PRIVILEGES ON * . * TO 'user'@'localhost';`
   Reload privileges `FLUSH PRIVILEGES;`
 - `connect acgtweet`
 - 