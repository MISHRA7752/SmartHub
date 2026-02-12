# SmartLoad Optimization API

A stateless microservice built in Go that optimizes truck loads by selecting the most profitable combination of orders while respecting weight, volume, and compatibility constraints.

## Features

- **Revenue Maximization**: Uses a recursive backtracking algorithm with pruning to find the optimal order combination.
- **Constraints Handling**:
  - Weight and Volume limits.
  - Hazmat Isolation (Hazmat vs Non-Hazmat).
  - Route Compatibility (grouping by Origin -> Destination).
- **Performance**: Capable of handling N=22 orders in < 100ms.
- **Stateless**: No database required, completely in-memory.

## Technologies

- **Language**: Go 1.24
- **Containerization**: Docker & Docker Compose
- **Architecture**: Clean Architecture (Handlers -> Service -> Domain)

## How to Run

### Prerequisites
- Docker & Docker Compose installed.

### Steps
1. Clone the repository:
   ```bash
   git clone <your-repo-url>
   cd SmartLoad
   ```

2. Start the service:
   ```bash
   docker compose up --build
   ```
   *Note: If you see `unknown flag: --build` or `docker: 'compose' is not a docker command`, try using the legacy command:*
   ```bash
   docker-compose up --build
   ```
   
   If neither works, you can build and run manually:
   ```bash
   docker build -t smartload .
   docker run -p 8080:8080 smartload
   ```

   The service will start on port `8080`.

3. Check health:
   ```bash
   curl http://localhost:8080/healthz
   ```

## API Usage

### Endpoint
`POST /api/v1/load-optimizer/optimize`

### Request Example
```json
{
  "truck": {
    "id": "truck-123",
    "max_weight_lbs": 44000,
    "max_volume_cuft": 3000
  },
  "orders": [
    {
      "id": "ord-001",
      "payout_cents": 250000,
      "weight_lbs": 18000,
      "volume_cuft": 1200,
      "origin": "Los Angeles, CA",
      "destination": "Dallas, TX",
      "pickup_date": "2025-12-05",
      "delivery_date": "2025-12-09",
      "is_hazmat": false
    }
    // ... more orders
  ]
}
```

### Response Example
```json
{
  "truck_id": "truck-123",
  "selected_order_ids": ["ord-001"],
  "total_payout_cents": 250000,
  "total_weight_lbs": 18000,
  "total_volume_cuft": 1200,
  "utilization_weight_percent": 40.91,
  "utilization_volume_percent": 40.0
}
```
