#!/usr/bin/env python3
# using 'flush': See https://stackoverflow.com/questions/74171437

import sys, time

status=0 if(len(sys.argv)==1) else int(sys.argv[1])

for a in range(status):
	if(a%2==0): print(a, file=sys.stdout, flush=True)
	else: print(a, file=sys.stderr, flush=True)
	time.sleep(0.1)

if(len(sys.argv)==1):
	print(f"This is an STDOUT message, Exit status: {status}", flush=True)
	print("This is an STDERR message", file=sys.stderr, flush=True)
else:
	print("This is an STDOUT message", flush=True)
	print(f"This is an STDERR message, Exit status: {status}", flush=True, file=sys.stderr)

sys.exit(status)

