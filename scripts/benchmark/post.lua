wrk.method = "POST"
wrk.body = '{"conversation_id": "[UUID]","sender_id": "[UUID]","content": "test message"}'
wrk.headers["Content-Type"] = "application/json"
