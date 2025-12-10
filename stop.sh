#!/bin/bash

# Eskio Stop Script
# Stops the backend server and frontend client

echo "🛑 Stopping Eskio..."
echo ""

# Get the script directory
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Stop backend
if [ -f "$SCRIPT_DIR/.backend.pid" ]; then
    BACKEND_PID=$(cat "$SCRIPT_DIR/.backend.pid")
    if ps -p $BACKEND_PID > /dev/null 2>&1; then
        echo "🔧 Stopping backend server (PID: $BACKEND_PID)..."
        kill $BACKEND_PID
        echo "✅ Backend stopped"
    else
        echo "⚠️  Backend process not found"
    fi
    rm "$SCRIPT_DIR/.backend.pid"
else
    echo "⚠️  No backend PID file found"
fi
echo ""

# Stop frontend
if [ -f "$SCRIPT_DIR/.frontend.pid" ]; then
    FRONTEND_PID=$(cat "$SCRIPT_DIR/.frontend.pid")
    if ps -p $FRONTEND_PID > /dev/null 2>&1; then
        echo "⚛️  Stopping frontend server (PID: $FRONTEND_PID)..."
        kill $FRONTEND_PID
        echo "✅ Frontend stopped"
    else
        echo "⚠️  Frontend process not found"
    fi
    rm "$SCRIPT_DIR/.frontend.pid"
else
    echo "⚠️  No frontend PID file found"
fi
echo ""

# Optionally stop PostgreSQL container
read -p "📦 Do you want to stop the PostgreSQL container? (y/N): " -n 1 -r
echo ""
if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo "🔄 Stopping PostgreSQL container..."
    cd "$SCRIPT_DIR/server" && docker-compose down
    echo "✅ PostgreSQL container stopped"
else
    echo "ℹ️  PostgreSQL container left running"
fi
echo ""

echo "✨ Eskio has been stopped!"
