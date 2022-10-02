tracker: FORCE
	tinygo flash --target wioterminal --size short ./wioterminal/tracker/

fukuwarai: FORCE
	tinygo flash --target wioterminal --size short --opt 2 ./wioterminal/fukuwarai/

smoketest: FORCE
	tinygo build -o /tmp/out.uf2 --target wioterminal --size short ./wioterminal/tracker/
	tinygo build -o /tmp/out.uf2 --target wioterminal --size short --opt 2 ./wioterminal/fukuwarai/
	make --directory chap07 smoketest

FORCE:
