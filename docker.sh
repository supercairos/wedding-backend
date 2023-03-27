docker run --name mysql \
    -e MYSQL_ROOT_PASSWORD=root-password \
    -e MYSQL_DATABASE=wishlist \
    -e MYSQL_USER=wishlist \
    -e MYSQL_PASSWORD=password \
    -p 3306:3306 \
    -d mysql:8