from cycletls import get


import time
start_time = time.time()
test_dict = {
		"url": "http://localhost:5000",
		# "ja3": "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-21,29-23-24,0",
		# "userAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36"
	}

resp = get(test_dict)
for i in range(10000):
    resp = get(test_dict)

print(resp)
print("--- %s seconds ---" % (time.time() - start_time))