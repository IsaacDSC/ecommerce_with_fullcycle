# ecommerce_with_fullcycle


adicionar depois do package e import no resolver.go
//go:generate go run github.com/99designs/gqlgen generate


command generate scaffold 
```sh
go generate ./...
```


### createProduct
```
mutation createProduct($data: NewProduct!){ 
  createProduct(input: $data){
    id
    name
  }
}

{
 "data": {
  "name": "camisa azul",
  "imageUrl": "image",
  "price": 9999,
  "description": "desc",
  "categoryId": "6372313d-1afb-480d-8015-86b26f9400ab",
  "active": true
  } 
}
```


### createCategory
```
mutation createCategory($data: NewCategory!){ 
  createCategory(input: $data){
    id
    name
  }
}

{
  "data": {
    "name": "Short"
  }
}
```

### GetCatalog
```
query GetCatalog {
  products{
     id
    name
    Category {
      id
      name
    }
  }
  categories{
    id
    name
  }
}
```

### GetProduct
```
query GetProduct($data: RetrieveByID) {
   getProduct(input: $data){
    id
    name
    imageUrl
    price
    description
    code
    Category{
       id
      name
    }
  }
}

{
  "data": {
    "ID": "9a82ed8d-ef56-46e9-8ca1-83c0e106f582"
  }
}
```
