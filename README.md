# ecommerce-go

## 🧠 Understanding Pointers and Constructors in Go (Using `repository_sql.go`)

In Go, understanding how to use `*` (dereference) and `&` (address-of) is essential when working with dependency injection and clean architecture.

---

### 📌 Key Concepts

| Symbol | Meaning                         | Usage                                |
|--------|----------------------------------|--------------------------------------|
| `&`    | Address of a value               | Used when passing a pointer          |
| `*`    | Dereferencing / Pointer receiver | Used in function signatures to work with the real value |

---

### ✅ Constructor Pattern in Go

Go doesn’t support traditional constructors like C# or Java, but by convention we use `NewType()` functions that return initialized struct instances (often as interfaces).

Here’s an example using a **SQL repository** for the `category` entity:

```go
func NewSQLRepository(db *sql.DB) Storage {
    return &repositorySQL{db: db}
}


/*
╔══════════════════════════════════════════════════════════════════╗
║     💡 Go Pointers and Constructor Usage Explanation            ║
╚══════════════════════════════════════════════════════════════════╝

📌 This file defines a SQL-based repository for categories.
   It uses a constructor function: NewSQLRepository(db *sql.DB)

🔹 Pointers:
   - Go functions can receive either values or pointers.
   - If a function or method expects a pointer (*Type), you usually pass it using &.

🔸 Constructor: NewSQLRepository
   - Signature: func NewSQLRepository(db *sql.DB) Storage
   - Parameter: expects a pointer to a sql.DB object.
   - Return: a pointer to a repositorySQL struct that implements the Storage interface.

✅ DO:
   db, _ := sql.Open("mysql", "...") // db is already *sql.DB
   repo := NewSQLRepository(db)      // Correct: db is already a pointer

🚫 DON'T:
   db := sql.DB{}                    // db is a value
   repo := NewSQLRepository(db)      // Error: expected *sql.DB

✅ FIX:
   repo := NewSQLRepository(&db)     // Correct: pass pointer explicitly

📎 Summary:
   - Use `&` when you have a value and need a pointer.
   - Use `*` to define functions that expect or operate on pointers.
   - `sql.Open(...)` returns a *sql.DB, so no `&` is needed in normal usage.

*/
```

---

### 💻 Example Code

```go
package category

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/ddessilvestri/ecommerce-go/models"
)

// This struct acts like a "class" in Go.
// It implements the Storage interface for SQL-based storage.
type repositorySQL struct {
	db *sql.DB // Dependency to the database connection
}

/*
	🔹 C# Equivalent:

	public class CategoryRepository {
	    private SqlConnection _db;
	    public CategoryRepository(SqlConnection db) {
	        _db = db;
	    }
	}
*/

// Constructor-like function (Go does not support constructors like C# or Java).
// By convention, we use New<Name>() to instantiate and return the interface type.
func NewSQLRepository(db *sql.DB) Storage {
	// We return a pointer to the struct instance
	return &repositorySQL{db: db}
}

/*
	🔹 C# Equivalent:

	public interface ICategoryStorage {
	    long InsertCategory(Category c);
	}

	public class CategoryRepository : ICategoryStorage {
	    public long InsertCategory(Category c) {
	        // SQL logic here
	    }
	}
*/

// Method bound to the repositorySQL struct.
// The receiver is a pointer (*repositorySQL), which allows modifying internal state
// and avoids copying the struct on each method call.
func (r *repositorySQL) InsertCategory(c models.Category) (int64, error) {
	// Build a safe SQL INSERT query using the squirrel package
	query, args, err := squirrel.
		Insert("category").
		Columns("Categ_Name", "Categ_Path").
		Values(c.CategName, c.CategPath).
		ToSql()

	if err != nil {
		return 0, err
	}

	// Execute the query with the generated SQL and arguments
	result, err := r.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	// Return the last inserted ID
	return result.LastInsertId()
}
```

---

### 🧪 Pointer Theory in Practice

```go
type User struct {
	Name string
}

func updateName(u *User) {
	u.Name = "Alice"
}

func main() {
	user := User{Name: "Original"}
	updateName(&user) // Pass address to modify original value
}
```

## 🧱 Design Patterns in Use

This project follows a **Clean Architecture** approach and implements several key design patterns to ensure maintainability, testability, and scalability.

### ✅ Patterns Currently Applied

| Pattern                 | Purpose                                                                 | Implementation Example                                              |
|-------------------------|-------------------------------------------------------------------------|----------------------------------------------------------------------|
| **Repository Pattern**   | Abstracts data access to support multiple storage backends              | [`category/repository_sql.go`](./internal/category/repository_sql.go) implements `Storage` interface |
| **Factory Pattern**      | Standardized way to construct components (similar to constructors)      | `NewSQLRepository`, `NewCategoryService`, `NewCategoryHandler`      |
| **Dependency Injection** | Injects dependencies from higher layers to lower layers, improves testing | Passed from `main.go → router → handler → service`                  |
| **Pointer Injection**    | Go idiomatic way to share resources like `*sql.DB` without copying      | Used across `repository`, `service`, `handler` layers               |

---

### ✨ Example: Repository Pattern

```go
type Storage interface {
    InsertCategory(c models.Category) (int64, error)
}

type repositorySQL struct {
    db *sql.DB
}

func NewSQLRepository(db *sql.DB) Storage {
    return &repositorySQL{db: db}
}
```

This design allows you to swap `repositorySQL` with a `MongoRepository` or `MockRepository` without changing the business logic that depends on the interface.

---

### 🔜 Coming Next: Middleware Pattern

We plan to introduce the **Middleware Pattern** to centralize and simplify cross-cutting concerns such as:

- ✅ JWT Token validation
- ✅ Admin access checks (`UserIsAdmin`)
- 🧪 Future logging, panic recovery, etc.

This will ensure cleaner and reusable handler logic.

#### 💡 Planned Branch

```
feature/add-middleware-auth
```

The goal is to enable middleware chaining that wraps handlers like:

```go
handlerWithMiddleware := middleware.Authenticate(UserIsAdmin)(handler.Post)
```

## 📋 Complete Repository Overview

This is a **Go-based ecommerce API** built for **AWS Lambda** using a clean architecture pattern. Here's the complete breakdown:

### 🏗️ **Architecture & Design Patterns**

The project follows **Clean Architecture** principles with these key patterns:

1. **Repository Pattern** - Abstracts data access through interfaces
2. **Dependency Injection** - Services depend on interfaces, not concrete implementations
3. **Factory Pattern** - Constructor functions like `NewSQLRepository()`
4. **Layered Architecture** - Handler → Service → Repository → Database

### 📁 **Project Structure**

```
ecommerce-go/
├── main.go                 # Lambda entry point
├── go.mod                  # Go module dependencies
├── README.md              # Comprehensive documentation
├── deploy.sh              # Deployment script
├── build.sh               # Build script
├── internal/              # Business logic modules
│   ├── category/          # Category management
│   ├── product/           # Product management  
│   ├── user/              # User management
│   ├── order/             # Order management
│   ├── address/           # Address management
│   ├── stock/             # Stock management
│   ├── admin/             # Admin functionality
│   └── config/            # Configuration management
├── routers/               # HTTP routing layer
├── models/                # Data structures
├── db/                    # Database utilities
├── auth/                  # Authentication & authorization
├── tools/                 # Utility functions
├── awsgo/                 # AWS SDK initialization
└── secretm/               # AWS Secrets Manager integration
```

### 🎯 **Key Technologies**

- **Language**: Go 1.23.5
- **Runtime**: AWS Lambda
- **Database**: MySQL (via RDS)
- **Authentication**: JWT tokens
- **Cloud**: AWS (Lambda, RDS, Secrets Manager)
- **SQL Builder**: Squirrel (for safe SQL queries)

### 🎯 **Core Features**

#### **Entities Managed:**
1. **Categories** - Product categorization
2. **Products** - Product catalog with search
3. **Users** - User management
4. **Orders** - Order processing
5. **Addresses** - Shipping addresses
6. **Stock** - Inventory management
7. **Admin** - Administrative functions

#### **API Endpoints:**
- `GET/POST/PUT/DELETE /category` - Category CRUD
- `GET/POST/PUT/DELETE /product` - Product CRUD with search
- `GET/POST/PUT/DELETE /order` - Order management
- `GET/POST/PUT/DELETE /user` - User management
- `GET/POST/PUT/DELETE /address` - Address management
- `GET/POST/PUT/DELETE /stock` - Stock management
- `GET/POST/PUT/DELETE /admin/users` - Admin user management

### 🏛️ **Architecture Layers**

#### **1. Router Layer** (`routers/`)
- Routes HTTP requests to appropriate entity handlers
- Handles authentication (JWT validation)
- Supports public endpoints (product/category GET)

#### **2. Handler Layer** (`internal/*/handler.go`)
- HTTP request/response handling
- JSON parsing and validation
- Error response formatting

#### **3. Service Layer** (`internal/*/service.go`)
- Business logic implementation
- Input validation
- Orchestrates repository calls

#### **4. Repository Layer** (`internal/*/repository_sql.go`)
- Data access abstraction
- SQL query building (using Squirrel)
- Database interaction

#### **5. Interface Layer** (`internal/*/interface.go`)
- Defines contracts between layers
- Enables dependency injection
- Supports testing and mocking

### 🎯 **Authentication & Security**

- **JWT Token Validation** - Extracts and validates tokens from Authorization header
- **Admin Role Checking** - `UserIsAdmin()` function for admin-only endpoints
- **AWS Secrets Manager** - Secure database credentials storage
- **Input Validation** - Service layer validation for all inputs

### 🎯 **Database Design**

The project uses a MySQL database with these key tables:
- `category` - Product categories
- `product` - Product catalog
- `users` - User accounts
- `orders` - Order records
- `order_details` - Order line items
- `address` - Shipping addresses
- `stock` - Inventory levels

### 🚀 **Deployment**

- **AWS Lambda** - Serverless deployment
- **Binary Packaging** - Compiled for Linux x86_64
- **Environment Variables** - Configuration via Lambda environment
- **AWS Secrets Manager** - Database credentials

### 🎯 **Key Dependencies**

```go
github.com/aws/aws-lambda-go v1.48.0    # Lambda runtime
github.com/Masterminds/squirrel v1.5.4   # SQL query builder
github.com/go-sql-driver/mysql v1.9.2    # MySQL driver
github.com/aws/aws-sdk-go-v2/*           # AWS SDK
```

### 🎨 **Code Quality Features**

1. **Clean Architecture** - Separation of concerns
2. **Dependency Injection** - Testable and maintainable
3. **Error Handling** - Comprehensive error management
4. **Input Validation** - Service layer validation
5. **SQL Injection Prevention** - Using Squirrel for safe queries
6. **JWT Token Security** - Proper token validation
7. **Admin Role Management** - Role-based access control

### 🎯 **Request Flow**

```
Lambda Request → Router → Handler → Service → Repository → Database
                ↓
            Authentication → Admin Check → Business Logic → Response
```

This is a well-structured, production-ready ecommerce API that demonstrates Go best practices, clean architecture principles, and AWS serverless deployment patterns. The code is well-documented, follows consistent patterns across all modules, and includes comprehensive error handling and security measures.