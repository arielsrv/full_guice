# Dependency Injection Example

This project demonstrates the use of dependency injection with the `go.uber.org/dig` package.

## Project Structure

The project is organized into three main packages:

- `di`: Contains the dependency injection components
- `services`: Contains the notification service
- `workers`: Contains the worker interface and implementations

## Dependency Graph Visualization

The application displays a dependency graph in DOT format when it starts. This graph shows the relationships between the components in the application.

### How to Interpret the Graph

The graph is displayed in DOT format, which can be visualized using tools like Graphviz. The graph shows:

- **Constructors**: Functions that create and provide dependencies (shown as rectangles)
- **Components**: The actual dependencies provided by constructors (shown as nodes)
- **Dependencies**: Arrows showing which components depend on other components

In our application:

1. The `ProvideWorkers` constructor from the `awesomeProject19/di` package provides two workers:
   - `workers.Worker[name=email_worker]`: The email worker
   - `workers.Worker[name=sms_worker]`: The SMS worker

2. The `NewNotificationService` constructor from the `awesomeProject19/services` package:
   - Depends on both workers
   - Provides a `*services.NotificationService`

### Visualizing with External Tools

For a more visual representation, you can save the DOT output to a file and use Graphviz to generate an image:

```bash
go run main.go > graph.dot
dot -Tpng graph.dot -o graph.png
```

## Running the Application

To run the application:

```bash
go run main.go
```

This will:
1. Display the dependency graph
2. Execute the notification service, which will use both workers to send notifications

## Testing

The project includes tests for all components. To run the tests:

```bash
go test ./...
```