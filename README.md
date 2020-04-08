## Getting Started

For running this Go app you should first build its image by running

`docker build -t qok-url-shortener-v0 . `

and then run 

```
docker run -p 8787:8787  qok-url-shortener-v0
```

`create database qok_url_shortener;`

```
CREATE TABLE urls ( id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY, long_url VARCHAR(100) NOT NULL, short_url VARCHAR(50), reg_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP );
```