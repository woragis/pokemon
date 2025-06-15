#/usr/bin/bash

# Run backend
(cd backend && ./build/main) &

# Run frontend
(cd frontend && bun dev) &

# Wait for both processes
wait
