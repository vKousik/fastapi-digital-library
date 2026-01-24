import time
from fastapi import  Request

async def logging_middleware(request: Request, call_next):
    start_time = time.perf_counter()
    user_agent = request.headers.get("User-Agent", "Unknown")    
    print(f"[LOG] Request received from: {user_agent}")

    response = await call_next(request)
    process_time = time.perf_counter() - start_time
    response.headers["X-Process-Time"] = f"{process_time:.6f}"
    
    return response