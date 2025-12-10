#!/bin/bash
# Eskio Startup Script
# Starts the database, backend server, and frontend client

echo "🚀 Starting Eskio..."
echo ""

# Get the script directory
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Check if PostgreSQL container is running
echo "📦 Checking PostgreSQL container..."
if docker ps | grep -q eskio-postgres; then
    echo "✅ PostgreSQL container is already running"
else
    echo "🔄 Starting PostgreSQL container..."
    cd "$SCRIPT_DIR/server" && docker-compose up -d
    echo "⏳ Waiting for PostgreSQL to be ready..."
    sleep 3
fi
echo ""

# Start the backend server
echo "🔧 Starting Go backend server on :8080..."
cd "$SCRIPT_DIR/server/cmd"
go run api/main.go &
BACKEND_PID=$!
echo "✅ Backend started (PID: $BACKEND_PID)"
echo ""

# Wait a moment for backend to initialize
sleep 2

# Start the frontend client
echo "⚛️  Starting Next.js frontend on :3000..."
cd "$SCRIPT_DIR/client"
npm run dev &
FRONTEND_PID=$!
echo "✅ Frontend started (PID: $FRONTEND_PID)"
echo ""

# Save PIDs for the stop script
echo "$BACKEND_PID" > "$SCRIPT_DIR/.backend.pid"
echo "$FRONTEND_PID" > "$SCRIPT_DIR/.frontend.pid"

echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "✨ Eskio is now running!"
echo ""
echo "📱 Frontend:  http://localhost:3000"
echo "🔌 Backend:   http://localhost:8080"
echo "🗄️  Database:  postgres://localhost:5432"
echo ""
echo "To stop all services, run: ./stop.sh"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""

# Keep script running and show logs
echo "📋 Showing logs (Ctrl+C to exit logs, services will keep running)..."
echo ""
wait
