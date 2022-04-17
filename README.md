# mdm
master data modeling

## data source

## data model

suppose we are building an online website for searching books. We are going to model this with
`Book`, `Author`, `Publisher`.

These data model has the following schema:

```
DataModel Book {
    id string, required, pk
    name string, required, unique
    authorId string, required, Ref<Author> 
    publisherId string, required, Ref<Publisher>
    publishedAt timestamp 
}

DataModel Author {
    id string, required, pk
    name string, required
}

DataModel Publisher {
    id string, required, pk
    name string, required
}

Relation AuthorPublisher {
    id string, required, pk
    authorId string, required, Ref<Author>
    publisherId string, required, Ref<Publisher>
}
```





