current=`date "+%Y-%m-%d %H:%M:%S"`
timestamp=`date -d "$current" +%s`
today=`date "+%Y%m%d"`
echo $today
echo $timestamp

docker build -t xxxxxxxxxx:$today-$timestamp -f Dockerfile .