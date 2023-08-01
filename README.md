<h1 align="center">
  🚀 🐹 Repository Pattern in Go 🐹 🚀 
</h1>

<p align="center">
    <a href="#"><img src="https://img.shields.io/badge/technology-go-blue.svg" alt="Go"/></a>
</p>

<p align="center">
  Project to explore some techniques to implement the repository pattern in go.
</p>

## 🤔 What's the repository pattern

The repository pattern is a design pattern that isolates the application/business layer from the data layer.

This generate an anti-corruption layer between the application and the data layer and isolates infrastructure details from services.

## 🧲 Environment Setup

### 🛠️ Needed tools

1. Go 1.20 (recommended)

### 🏃🏻 Application execution

1. Make sure to download all Needed tools
2. Clone the repository
```bash
git clone https://github.com/dasalgadoc/repository-go.git
```
3. Build up go project
```bash
go mod download
go get .
```

## 📚 References

- [Repository pattern by Martin Fowler](https://martinfowler.com/eaaCatalog/repository.html)
- [Repository pattern by Microsoft](https://learn.microsoft.com/en-us/dotnet/architecture/microservices/microservice-ddd-cqrs-patterns/infrastructure-persistence-layer-design)