#!/bin/sh

#
# basic map-reduce test
#

RACE=
WC_SO="../mrapps/wc/wc.so"
INDEXER_SO="../mrapps/indexer/indexer.so"
MTIMING_SO="../mrapps/mtiming/mtiming.so"
RTIMING_SO="../mrapps/rtiming/rtiming.so"
CRASH_SO="../mrapps/crash/crash.so"
NOCRASH_SO="../mrapps/nocrash/nocrash.so"
MR_MASTER="../cmd/mrmaster/mrmaster"
MR_WORKER="../cmd/mrworker/mrworker"
MR_SEQUENTIAL="../cmd/mrsequential/mrsequential"
ALLDATA="../data/pg*txt"

# uncomment this to run the tests with the Go race detector.
#RACE=-race

# run the test in a fresh sub-directory.
rm -rf mr-tmp
mkdir mr-tmp || exit 1
cd mr-tmp || exit 1
rm -f mr-*

# make sure software is freshly built.
(cd ../mrapps/wc && go build $RACE -buildmode=plugin wc.go) || exit 1
(cd ../mrapps/indexer && go build $RACE -buildmode=plugin indexer.go) || exit 1
(cd ../mrapps/mtiming && go build $RACE -buildmode=plugin mtiming.go) || exit 1
(cd ../mrapps/rtiming && go build $RACE -buildmode=plugin rtiming.go) || exit 1
(cd ../mrapps/crash && go build $RACE -buildmode=plugin crash.go) || exit 1
(cd ../mrapps/nocrash && go build $RACE -buildmode=plugin nocrash.go) || exit 1
(cd ../cmd/mrmaster && go build $RACE mrmaster.go) || exit 1
(cd ../cmd/mrworker && go build $RACE mrworker.go) || exit 1
(cd ../cmd/mrsequential && go build $RACE mrsequential.go) || exit 1

failed_any=0

# first word-count

# generate the correct output
$MR_SEQUENTIAL $WC_SO $ALLDATA || exit 1
sort mr-out-0 > mr-correct-wc.txt
rm -f mr-out*

echo '***' Starting wc test.

timeout -k 2s 180s $MR_MASTER $ALLDATA &

# give the master time to create the sockets.
sleep 1

# start multiple workers.
timeout -k 2s 180s $MR_WORKER $WC_SO &
timeout -k 2s 180s $MR_WORKER $WC_SO &
timeout -k 2s 180s $MR_WORKER $WC_SO &

# wait for one of the processes to exit.
# under bash, this waits for all processes,
# including the master.
wait

# the master or a worker has exited. since workers are required
# to exit when a job is completely finished, and not before,
# that means the job has finished.

sort mr-out* | grep . > mr-wc-all
if cmp mr-wc-all mr-correct-wc.txt
then
  echo '---' wc test: PASS
else
  echo '---' wc output is not the same as mr-correct-wc.txt
  echo '---' wc test: FAIL
  failed_any=1
fi

# wait for remaining workers and master to exit.
wait ; wait ; wait

# now indexer
rm -f mr-*

# generate the correct output
$MR_SEQUENTIAL $INDEXER_SO $ALLDATA || exit 1
sort mr-out-0 > mr-correct-indexer.txt
rm -f mr-out*

echo '***' Starting indexer test.

timeout -k 2s 180s $MR_MASTER $ALLDATA &
sleep 1

# start multiple workers
timeout -k 2s 180s $MR_WORKER $INDEXER_SO &
timeout -k 2s 180s $MR_WORKER $INDEXER_SO

sort mr-out* | grep . > mr-indexer-all
if cmp mr-indexer-all mr-correct-indexer.txt
then
  echo '---' indexer test: PASS
else
  echo '---' indexer output is not the same as mr-correct-indexer.txt
  echo '---' indexer test: FAIL
  failed_any=1
fi

wait ; wait


echo '***' Starting map parallelism test.

rm -f mr-out* mr-worker*

timeout -k 2s 180s $MR_MASTER $ALLDATA &
sleep 1

timeout -k 2s 180s $MR_WORKER $MTIMING_SO &
timeout -k 2s 180s $MR_WORKER $MTIMING_SO

NT=`cat mr-out* | grep '^times-' | wc -l | sed 's/ //g'`
if [ "$NT" != "2" ]
then
  echo '---' saw "$NT" workers rather than 2
  echo '---' map parallelism test: FAIL
  failed_any=1
fi

if cat mr-out* | grep '^parallel.* 2' > /dev/null
then
  echo '---' map parallelism test: PASS
else
  echo '---' map workers did not run in parallel
  echo '---' map parallelism test: FAIL
  failed_any=1
fi

wait ; wait


echo '***' Starting reduce parallelism test.

rm -f mr-out* mr-worker*

timeout -k 2s 180s $MR_MASTER $ALLDATA &
sleep 1

timeout -k 2s 180s $MR_WORKER $RTIMING_SO &
timeout -k 2s 180s $MR_WORKER $RTIMING_SO

NT=`cat mr-out* | grep '^[a-z] 2' | wc -l | sed 's/ //g'`
if [ "$NT" -lt "2" ]
then
  echo '---' too few parallel reduces.
  echo '---' reduce parallelism test: FAIL
  failed_any=1
else
  echo '---' reduce parallelism test: PASS
fi

wait ; wait


# generate the correct output
$MR_SEQUENTIAL $NOCRASH_SO $ALLDATA || exit 1
sort mr-out-0 > mr-correct-crash.txt
rm -f mr-out*

echo '***' Starting crash test.

rm -f mr-done
(timeout -k 2s 180s $MR_MASTER $ALLDATA ; touch mr-done ) &
sleep 1

# start multiple workers
timeout -k 2s 180s $MR_WORKER $CRASH_SO &

# mimic rpc.go's masterSock()
SOCKNAME=/var/tmp/824-mr-`id -u`

( while [ -e $SOCKNAME -a ! -f mr-done ]
  do
    timeout -k 2s 180s $MR_WORKER $CRASH_SO
    sleep 1
  done ) &

( while [ -e $SOCKNAME -a ! -f mr-done ]
  do
    timeout -k 2s 180s $MR_WORKER $CRASH_SO
    sleep 1
  done ) &

while [ -e $SOCKNAME -a ! -f mr-done ]
do
  timeout -k 2s 180s $MR_WORKER $CRASH_SO
  sleep 1
done

wait
wait
wait

rm $SOCKNAME
sort mr-out* | grep . > mr-crash-all
if cmp mr-crash-all mr-correct-crash.txt
then
  echo '---' crash test: PASS
else
  echo '---' crash output is not the same as mr-correct-crash.txt
  echo '---' crash test: FAIL
  failed_any=1
fi

if [ $failed_any -eq 0 ]; then
    echo '***' PASSED ALL TESTS
else
    echo '***' FAILED SOME TESTS
    exit 1
fi
