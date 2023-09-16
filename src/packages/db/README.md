## Examples
```docker
docker run --name telebot-mongo -v $(pwd)/data:/data -p 27017:27017 -d mongo
```

```golang
  // Find documents
  query := bson.M{"key": "value"}
  results, err := db.Find(query)
  if err != nil {
    log.Fatal(err)
  }
  for _, result := range results {
    fmt.Println(result)
  }

  // Update documents
  filter := bson.M{"key": "value"}
  update := bson.M{"$set": bson.M{"key": "new_value"}}
  err = db.Update(filter, update)
  if err != nil {
    log.Fatal(err)
  }

  // Delete documents
  deleteQuery := bson.M{"key": "new_value"}
  err = db.Delete(deleteQuery)
  if err != nil {
    log.Fatal(err)
  }
```