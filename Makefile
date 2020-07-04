clean-mr:
	rm -r mr-tmp
	rm cmd/mrmaster/mrmaster
	rm cmd/mrsequential/mrsequential
	rm cmd/mrworker/mrworker
	rm mrapps/crash/crash.so
	rm mrapps/indexer/indexer.so
	rm mrapps/mtiming/mtiming.so
	rm mrapps/nocrash/nocrash.so
	rm mrapps/rtiming/rtiming.so
	rm mrapps/wc/wc.so