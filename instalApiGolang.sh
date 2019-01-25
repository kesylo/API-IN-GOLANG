sudo apt-get install redis-server -y	
echo "installation redis-server finie "

go get -u github.com/tidwall/gjson
echo "installation tdwall/gjson terminée"

go get -u google.golang.org/api/calendar/v3
echo "installation calendar terminée"
go get golang.org/x/oauth2	
mkdir -p API-IN-GOLANG
echo "installation oauth2/google terminée"
	


