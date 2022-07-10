# recipes-api

go get -u github.com/gin-gonic/gin
go get -u github.com/rs/xid

init
``` text
 1010  git clone git@github.com:similarface/recipes-api.git
 1011  git status
 1013  cd recipes-api
 1014  ls
 1015  git checkout -b preprod
 1016  git push origin preprod
 1017  git checkout -b develop
 1018  git push origin develop
 1019  go  mod init
```


POSTMAN 使用：

``` bash
“curl --location --request POST 'http://localhost:8080/recipes' \
--header 'Content-Type: application/json' --data-raw '{
   "name": "Homemade Pizza",
   "tags" : ["italian", "pizza", "dinner"],
   "ingredients": [
       "1 1/2 cups (355 ml) warm water (105°F-115°F)",
       "1 package (2 1/4 teaspoons) of active dry yeast",
       "3 3/4 cups (490 g) bread flour",
       "feta cheese, firm mozzarella cheese, grated"
   ],
   "instructions": [
       "Step 1.",
       "Step 2.",
       "Step 3."
   ]
}' | jq -r
```

GET
``` postman
curl -s --location --request GET 'http://localhost:8080/recipes' \
--header 'Content-Type: application/json
```

``` postman
curl -s -X GET 'http://localhost:8080/recipes' | jq length
```

