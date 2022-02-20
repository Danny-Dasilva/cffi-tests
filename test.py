import faster_than_requests
import time
start_time = time.time()
for i in range(1000):
    resp = faster_than_requests.get("http://localhost:5000/")

print(resp)
print("--- %s seconds ---" % (time.time() - start_time))