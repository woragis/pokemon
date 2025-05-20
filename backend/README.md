# Backend

## Packages

```sh
go get github.com/gofiber/fiber/v2 # API
go get github.com/gofiber/fiber/v2/middleware # Middleware
go get github.com/gofiber/fiber/v2/middleware/logger # Logger for fiber
go get github.com/joho/godotenv # .env
go get github.com/golang-jwt/jwt/v5 # json web token
go get golang.org/x/crypto/bcrypt # password hashing
go get github.com/redis/go-redis/v9 # Redis cache
go get github.com/jackc/pgx/v5 # Postgres
go get github.com/jackc/pgx/v5/pgxpool # Postgres but focused on pooling
go get github.com/go-playground/validator/v10 # Validator
go get go.uber.org/zap # Logging for general purpouse
go get github.com/google/uuid # Uuid for database
go get github.com/sashabaranov/go-openai # GPT AI
go get github.com/go-resty/resty/v2 # Gemini AI
go get github.com/stripe/stripe-go/v78 # Stripe payments
go get github.com/gofiber/websocket/v2 # Websocket
go get gorm.io/gorm # Postgres Gorm
```

## Pokedex

### Planning

```js
PostgreSQL (Relational)
├── trainers
├── games
├── pokemon_species
├── user_pokedex_progress
└── user_game_completion (optional for living/shiny/etc metrics)

MongoDB (Flexible)
└── user_journeys
    └── [
          {
            trainer_id: "abc123",
            game: "Platinum",
            type: "achievement" | "todo" | "strategy",
            content: "...",
            tags: ["team building", "elite4"],
            created_at: ...
          }
        ]
```
