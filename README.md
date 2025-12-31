# Eino Feature Testing Project

A demonstration project showcasing **Cloudwego Eino** framework capabilities for LLM application development. This project implements various Eino patterns including chains, parallel execution, branching, tool integration, and graph-based workflows.

## ✨ Features

### Eino Framework Demonstrations

This project showcases the following Eino capabilities:

- **Basic Chat** - Simple LLM conversation with message history
- **Prompt Templates** - Go template-based prompt construction with variable interpolation
- **Parallel Execution** - Concurrent execution of multiple LLM operations
- **Branch Logic** - Conditional workflow branching based on runtime conditions
- **Tool Integration** - External tool calling (DuckDuckGo web search)
- **Graph Workflows** - Node-based execution graphs with explicit edge definitions
- **Complex Graphs** - Advanced graph workflows with conditional branching

### Architecture Features

- **Clean 3-Layer Architecture** (Handler → UseCase → Repository)
- **Custom Error Handling** system with 4-digit error codes aligned to HTTP status codes
- **Structured Logging** with TrID (Transaction ID) tracking using Zerolog
- **HTTP Middleware Stack** (CORS, logging, request ID, recovery)
- **Standard JSON Response Format** with transaction ID and status codes

## 🛠 Tech Stack

- **Go** 1.24+
- **Cloudwego Eino** v0.6.0 - LLM orchestration framework
- **Eino-ext** - Ollama model adapter and tool extensions
- **Ollama** - Local LLM inference engine
- **Chi Router** - Lightweight, idiomatic HTTP router
- **Zerolog** - Structured JSON logging
- **LangChainGo** - Additional LLM utilities
- **go.uber.org/mock** - Mock generation for testing

## 📋 Prerequisites

- **Go 1.24 or higher**
- **Ollama** - For running LLMs locally
- **Make** (optional, for convenience commands)

## 🚀 Quick Start

### 1. Install Ollama

Download and install Ollama from [https://ollama.com](https://ollama.com)

```bash
# macOS/Linux
curl -fsSL https://ollama.com/install.sh | sh

# Or download from website for your platform
```

### 2. Pull the LLM Model

The project is configured to use `gemma3:1b` by default:

```bash
ollama pull gemma3:1b
```

Verify Ollama is running:

```bash
ollama list
# Should show gemma3:1b in the list
```

### 3. Clone the Repository

```bash
git clone <repository-url>
cd eino-sample
```

### 4. Configure Environment

Create a `.env.local` file in the project root:

```env
PORT=8080
ENV=local
```

You can modify these values as needed for your environment.

### 5. Setup and Build

Run the all-in-one setup command:

```bash
make all
```

This will:

- Install development tools (golangci-lint, mockgen)
- Initialize Go modules
- Tidy dependencies
- Vendor dependencies
- Build the application

### 6. Run the Server

```bash
# Run the built binary
./bin/server

# Or run directly with go (for development)
go run cmd/server/main.go
```

You should see:

```
[ASCII art banner]
Configuration loaded: ENV=local, PORT=8080, ...
HTTP server starting on :8080
```

## 🧪 Testing the API

### Health Check

```bash
curl http://localhost:8080/healthz
```

**Response:**

```json
{
  "trid": "2025103112345678901234",
  "code": "0200",
  "result": {
    "status": "ok"
  }
}
```

### Basic Chat

Test the basic Eino chat functionality:

```bash
curl -X POST http://localhost:8080/basic-chat \
  -H "Content-Type: application/json" \
  -d '{
    "msg": "Hello"
  }'
```

**Response:**

```json
{
  "trid": "2025103112345678901235",
  "code": "0200",
  "result": "The three main functions of LangChain are: 1) Prompt Management, 2) Chain Composition, and 3) Agent Development."
}
```

This endpoint demonstrates:

- Multi-turn conversation with system and user messages
- Message history management
- Basic LLM Generate() call

### Prompt Template Chat

Test prompt template with variable interpolation and JSON parsing:

```bash
curl -X POST http://localhost:8080/basic-chat/prompt-template \
  -H "Content-Type: application/json" \
  -d '{
    "msg": "Generate report"
  }'
```

**Response:**

```json
{
  "trid": "2025103112345678901236",
  "code": "0200",
  "result": "The user is feeling neutral."
}
```

This endpoint demonstrates:

- Go template-based prompt construction
- Variable interpolation (user, company, date)
- JSON response parsing with custom parser
- Chain composition: Template → Model → Parser
- Graph workflow with emotion analysis and conditional branching

## 📁 Project Structure

```
├── cmd/
│   └── server/                # HTTP server entrypoint
│       └── main.go            # Application bootstrap
├── internal/                  # Private application code
│   ├── config/                # Configuration management
│   │   ├── config.go          # Environment variable loading
│   │   └── banner.asc         # ASCII art banner
│   ├── constants/             # Error codes and constants
│   │   ├── errors.go          # Error code definitions
│   │   └── app.go             # Application constants
│   ├── database/              # LLM initialization
│   │   └── llm.go             # Ollama LLM configuration
│   ├── domain/                # (Reserved for future domain models)
│   ├── handler/               # HTTP layer
│   │   └── http/
│   │       ├── basic_chat_controller.go  # Chat endpoints
│   │       ├── health_controller.go      # Health check
│   │       ├── routes.go                 # Route definitions
│   │       ├── dto/                      # Data transfer objects
│   │       └── middleware/               # HTTP middleware
│   │           ├── cors.go               # CORS handling
│   │           ├── http_logger.go        # Request logging
│   │           └── tr_id.go              # Transaction ID injection
│   ├── repository/            # Eino implementations
│   │   ├── repository.go      # Repository interfaces
│   │   └── langchain/ollama/
│   │       └── basic_chat_repo.go  # All Eino feature implementations
│   │   └── shared/
│   │       └── parser.go      # JSON parser utilities
│   ├── usecase/               # Business logic
│   │   ├── service.go         # Service interfaces
│   │   └── basic_chat_service.go  # Service implementation
│   └── shared/                # Shared utilities
│       └── utils/
├── pkg/                       # Reusable packages
│   ├── constants/             # Shared constants
│   ├── errors/                # Custom error system
│   │   └── errors.go          # Error codes and wrapping
│   ├── logger/                # Logging utilities
│   │   └── logger.go          # Zerolog wrapper with TrID
│   └── utils/                 # Common utilities
│       ├── http.go            # HTTP helpers
│       ├── id.go              # ID generation
│       └── string.go          # String utilities
├── mock/                      # Generated mocks
├── test/                      # Tests and test utilities
├── main.go                    # Simple Ollama test script
├── Makefile                   # Development commands
├── go.mod                     # Go module definition
├── go.sum                     # Go module checksums
└── README.md                  # This file
```

## 🏗 Architecture

### 3-Layer Architecture

```
┌─────────────────────────────────────────┐
│         Handler Layer (HTTP)            │
│  • HTTP Controllers                     │
│  • Middleware (CORS, Logging, TrID)     │
│  • Request/Response DTOs                │
│  • Error handling & logging             │
└──────────────┬──────────────────────────┘
               │ depends on
               ▼
┌─────────────────────────────────────────┐
│       UseCase Layer (Service)           │
│  • Business orchestration               │
│  • Service interfaces                   │
│  • Application workflows                │
└──────────────┬──────────────────────────┘
               │ depends on
               ▼
┌─────────────────────────────────────────┐
│     Repository Layer (Eino LLM)         │
│  • Eino chains, graphs, workflows       │
│  • LLM model interactions               │
│  • Tool integrations                    │
│  • Prompt template management           │
└─────────────────────────────────────────┘
```

### Dependency Flow

```
Handler (HTTP) → UseCase (Service) → Repository (Eino)
     ↓                  ↓                    ↓
  [Log]          [Orchestrate]        [LLM Chains/Graphs]
```

## 🎯 Eino Features in Detail

All Eino feature implementations are in [`internal/repository/langchain/ollama/basic_chat_repo.go`](internal/repository/langchain/ollama/basic_chat_repo.go). While only 2 are exposed via HTTP endpoints, you can explore all 7 implementations:

### 1. Basic Chat (`AskBasicChat`)

Simple conversation with LLM using message history.

```go
messages := []*schema.Message{
    {Role: schema.System, Content: "You are a helpful assistant."},
    {Role: schema.User, Content: "Please explain about langchain."},
    {Role: schema.Assistant, Content: "LangChain is a library..."},
    {Role: schema.User, Content: "Please answer the 3 main function."},
}
resp, err := r.ollamaLLM.Generate(ctx, messages)
```

**Key concepts:** Message history, system prompts, basic Generate() call

### 2. Prompt Template (`AskBasicPromptTemplateChat`)

Go template-based prompts with variable interpolation and JSON parsing.

```go
template := prompt.FromMessages(
    schema.GoTemplate,
    schema.SystemMessage("You are a JSON-only response assistant..."),
    schema.UserMessage("Generate a report for {{.user}} on {{.date}}..."),
)

chain, err := compose.NewChain[map[string]any, *JSONResponse]().
    AppendChatTemplate(template).
    AppendChatModel(r.ollamaLLM).
    AppendLambda(jsonParserLambda).
    Compile(ctx)

result, err := chain.Invoke(ctx, variables)
```

**Key concepts:** Template variables, chain composition, custom parsers

### 3. Parallel Execution (`AskBasicParallelChat`)

Concurrent execution of multiple operations.

```go
finalChain, err := compose.NewChain[map[string]any, map[string]any]().
    AppendParallel(
        compose.NewParallel().
            AddGraph("ask", askChain).           // LLM call
            AddLambda("length", lengthLambda).   // String length
            AddLambda("upper", upperLambda),     // Uppercase
    ).
    Compile(ctx)
```

**Key concepts:** Parallel component, concurrent operations, result aggregation

### 4. Branch Logic (`AskBasicBranchChat`)

Conditional workflow branching.

```go
roleCond := func(ctx context.Context, kvs map[string]any) (string, error) {
    if kvs["word"] == "a" {
        return "dog", nil
    }
    return "cat", nil
}

chain, err := compose.NewChain[map[string]any, *schema.Message]().
    AppendBranch(
        compose.NewChainBranch(roleCond).
            AddLambda("dog", dog).
            AddLambda("cat", cat),
    ).
    AppendChatTemplate(template).
    AppendChatModel(r.ollamaLLM).
    Compile(ctx)
```

**Key concepts:** Conditional branching, dynamic workflow paths

### 5. Tool Integration (`AskWithTool`)

Integration of external tools (web search).

```go
searchTool, err := duckduckgo.NewTextSearchTool(ctx, &duckduckgo.Config{
    MaxResults: 3,
    Region:     duckduckgo.RegionWT,
})

llmWithTools, err := r.ollamaLLM.WithTools([]*schema.ToolInfo{toolInfo})

chain, err := compose.NewChain[map[string]any, *schema.Message]().
    AppendChatTemplate(initialPrompt).
    AppendChatModel(llmWithTools).
    AppendToolsNode(toolsNode).
    AppendChatModel(r.ollamaLLM).
    Compile(ctx)
```

**Key concepts:** Tool calling, ToolNode, multi-step chains

### 6. Graph Workflow (`AskWithGraph`)

Node-based execution graphs.

```go
g := compose.NewGraph[map[string]any, *schema.Message]()
g.AddLambdaNode(greetingNode, greeting)
g.AddLambdaNode(processNode, process)
g.AddEdge(compose.START, greetingNode)
g.AddEdge(greetingNode, processNode)
g.AddEdge(processNode, compose.END)

res, err := g.Compile(ctx)
result, err := res.Invoke(ctx, input)
```

**Key concepts:** Graph nodes, edges, START/END nodes

### 7. Graph with Branch (`AskWithGraphWithBranch`)

Complex graph with conditional branching (emotion analysis).

```go
const (
    nodeOfPrompt    = "prompt"
    nodeOfModel     = "model"
    nodeOfEmotion   = "emotion"
    nodeOfPositive  = "positive"
    nodeOfNegative  = "negative"
    nodeOfNeutral   = "neutral"
)

g.AddBranch(nodeOfEmotion, compose.NewGraphBranch(cond, map[string]bool{
    "positive": true,
    "negative": true,
    "neutral":  true,
}))
```

**Key concepts:** GraphBranch, emotion classification, dynamic routing

## ⚙️ Configuration

### Environment Variables

| Variable | Description | Default/Example | Required |
| -------- | ----------- | --------------- | -------- |
| `PORT`   | Server port | `8080`          | Yes      |
| `ENV`    | Environment | `local`         | Yes      |

### Ollama Configuration

The Ollama LLM is configured in [`internal/database/llm.go`](internal/database/llm.go):

```go
ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
    BaseURL: "http://localhost:11434",  // Ollama service address
    Timeout: 30 * time.Second,          // Request timeout
    Model:   "gemma3:1b",                // Model name
})
```

To use a different model:

1. Pull the model: `ollama pull <model-name>`
2. Update `Model` field in `llm.go`
3. Rebuild and restart the server

Popular models:

- `gemma3:1b` - Fast, lightweight
- `llama2` - General purpose
- `codellama` - Code-focused
- `mistral` - Balanced performance
- `deepseek-r1:8b` - Reasoning model

## 🔧 Development Guide

### Build and Run

```bash
# Build the application
make build

# Run the built binary
./bin/server

# Or run directly
make start

# Development mode (no build)
go run cmd/server/main.go
```

### Testing

```bash
# Run tests
make test

# Run tests with coverage
go test -v -cover ./...

# Generate mocks
make build-mocks

# Full test suite (tests + vet + fmt + lint)
make test-all
```

### Code Quality

```bash
# Format code
make fmt

# Vet code
make vet

# Lint code (requires golangci-lint)
make lint
```

### Development Tools

```bash
# Install all development tools
make tool

# This installs:
# - golangci-lint (linting)
# - mockgen (mock generation)
```

### Clean Up

```bash
# Clean build artifacts, mocks, vendor
make clean
```

### Vendor Dependencies

```bash
# Vendor all dependencies
make vendor
```

## 📡 API Documentation

### Endpoints

| Method | Path                          | Description                     |
| ------ | ----------------------------- | ------------------------------- |
| `GET`  | `/healthz`                    | Health check                    |
| `POST` | `/basic-chat`                 | Basic LLM chat                  |
| `POST` | `/basic-chat/prompt-template` | Prompt template with graph flow |

### Request/Response Format

All responses follow a standard format:

```json
{
  "trid": "string", // Transaction ID for request tracing
  "code": "string", // 4-digit status code (e.g., "0200", "0404", "0500")
  "result": {} // Response data or error message
}
```

#### Success Response

```json
{
  "trid": "2025103112345678901234",
  "code": "0200",
  "result": "Your LLM response here"
}
```

#### Error Response

```json
{
  "trid": "2025103112345678901235",
  "code": "0500",
  "result": {
    "msg": "failed to ask: LLM error details"
  }
}
```

### Error Codes

The custom error system uses 4-digit codes aligned with HTTP status codes:

| Code   | HTTP Status         | Description               |
| ------ | ------------------- | ------------------------- |
| `0200` | 200 OK              | Success                   |
| `0400` | 400 Bad Request     | Invalid input             |
| `0404` | 404 Not Found       | Resource not found        |
| `0409` | 409 Conflict        | Conflict (e.g. duplicate) |
| `0500` | 500 Internal Server | Server/LLM error          |

## 🔑 Key Conventions

### Error Handling

Always use the custom error system:

```go
// Create error with code
errors.New(constants.InternalServerError, "LLM failed", err)

// Wrap error with context
errors.Wrap(err, "failed to generate response")

// Extract error code in handler
code := errors.GetCode(err)  // Returns "0500"
```

### Logging Strategy

- **Log ONLY at handler layer** (controllers, middleware)
- **Never log in usecase or repository layers**
- All logs include TrID for request tracing
- Use structured JSON format

```go
logger.LogInfo(ctx, "request received")
logger.LogError(ctx, "LLM error", err)
```

### Transaction ID (TrID)

Every request gets a unique transaction ID for tracing:

```go
// Generated by middleware
trid := ctx.Value(constants.ContextKeyTrID).(string)

// Included in all logs and responses
```

## 🤝 Contributing

Contributions are welcome! When adding new features:

1. **Testing new Eino features**: Add new methods to `basic_chat_repo.go`
2. **Adding endpoints**: Update `routes.go` and create controller methods
3. **Improving prompts**: Enhance prompt templates for better LLM responses
4. **Documentation**: Update this README with new examples

### Development Flow

1. Follow conventions in `.cursor/rules/convention.mdc`
2. Write tests for new features (use `make build-mocks`)
3. Run `make test-all` before committing
4. Use descriptive commit messages
5. Keep functions short and focused
6. Document public APIs with GoDoc comments

## 📚 Resources

- [Cloudwego Eino Documentation](https://github.com/cloudwego/eino)
- [Ollama Documentation](https://github.com/ollama/ollama)
- [Go Chi Router](https://github.com/go-chi/chi)
- [Zerolog](https://github.com/rs/zerolog)

## 🔨 Development Commands Reference

| Command            | Description                     |
| ------------------ | ------------------------------- |
| `make build`       | Build the application           |
| `make start`       | Run the built binary            |
| `make test`        | Run unit tests                  |
| `make build-mocks` | Generate mock implementations   |
| `make test-all`    | Run tests + vet + fmt + lint    |
| `make fmt`         | Format code with go fmt         |
| `make vet`         | Vet code with go vet            |
| `make lint`        | Lint code with golangci-lint    |
| `make tool`        | Install development tools       |
| `make clean`       | Clean build artifacts           |
| `make vendor`      | Vendor dependencies             |
| `make all`         | Full setup: tool + init + build |

---

**Note**: This project is designed as a learning and testing environment for Cloudwego Eino framework. Feel free to experiment with different LLM models, prompts, and Eino patterns!
