# Elastic API

[tutorials](/README.md) / databases / elastic

API Documentation:

- https://www.elastic.co/guide/en/elasticsearch/reference/current/rest-apis.html

## CURL commands

- `-v` Verbose
- `-u "elasticUser:elasticPass"`
- `-k` Ignore unsafe certificate
- `-X GET` to use http method GET

First set the environment variables as desired:

```shell
$ ES_AUTH="elastic:<password>"
$ ES_OPTS='-vkX GET'
$ ES_HOST=https://localhost:9200
```

Try some commands:

```shell
$ curl -u $ES_AUTH $ES_OPTS "$ES_HOST/_security/privilege?pretty"
$ curl -u $ES_AUTH $ES_OPTS "$ES_HOST/_all/_mapping?pretty"
```
