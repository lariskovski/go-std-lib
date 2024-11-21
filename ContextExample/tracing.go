package main

import (
    "context"
    "log"
)

// Define a key type to avoid context key collisions
type contextKey string

const userKey contextKey = "user"

// Function to handle a request
func handleRequest(ctx context.Context, user string) {
    // Create a new context with user information
    ctx = context.WithValue(ctx, userKey, user)

    // Call multiple functions, passing the context
    result1 := function1(ctx)
    result2 := function2(ctx)

    // Log the results
    log.Printf("User: %s, Result1: %s, Result2: %s", user, result1, result2)
}

// Function 1
func function1(ctx context.Context) string {
    // Extract user information from context
    user := ctx.Value(userKey).(string)
    // Log the user and function call
    log.Printf("Function1 called by user: %s", user)
    // Return some result
    return "Result from function1"
}

// Function 2
func function2(ctx context.Context) string {
    // Extract user information from context
    user := ctx.Value(userKey).(string)
    // Log the user and function call
    log.Printf("Function2 called by user: %s", user)
    // Return some result
    return "Result from function2"
}

func main() {
    // Simulate a user request
    user := "user123"
    handleRequest(context.Background(), user)
}