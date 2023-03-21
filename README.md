Service for storing quiz questions.

# Endpoints

## GET /questions

### Path Parameters
| Name        | Type    | Optional |
|:------------|:--------|:---------|
| category_id | integer | Optional |
| limit       | integer | Optional |
| locale_code | string  | Optional |
| type_id     | integer | Optional |

### Example

http://localhost:8080/questions?category_id=1&limit=3&locale_code=ru_RU&type_id=1

```json
[  
  {  
    "id": 1,  
    "category_id": 1,  
    "type_id": 1,  
    "text": "Каким был первый полнометражный анимационный фильм?",  
    "media_type": 0,  
    "media_url": "",  
    "locale_code": "ru_RU",  
    "answers": [  
      {  
        "id": 1,  
        "question_id": 1,  
        "text": "Покахонтас",  
        "is_valid": false  
      },  
      {  
        "id": 2,  
        "question_id": 1,  
        "text": "Белоснежка и семь гномов",  
        "is_valid": true  
      },  
      {  
        "id": 3,  
        "question_id": 1,  
        "text": "Русалочка",  
        "is_valid": false  
      },  
      {  
        "id": 4,  
        "question_id": 1,  
        "text": "Золушка",  
        "is_valid": false  
      }  
    ],  
    "created_at": "0001-01-01T00:00:00Z",  
    "updated_at": "0001-01-01T00:00:00Z",  
    "DeletedAt": null  
  }  
]
```

## GET /categories

### Path Parameters
None

### Example

http://localhost:8080/categories

```json
[  
  {  
    "id": 1,  
    "text": "Test"  
  }  
]
```

## GET /locales

### Path Parameters
None

### Example

http://localhost:8080/locales

```json
[  
  {  
    "code": "en_US",  
    "language": "English",  
    "country": "United States"  
  },  
  {  
    "code": "ru_RU",  
    "language": "Russian",  
    "country": "Russian Federation"  
  }  
]
```

