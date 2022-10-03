clear

while true
do
	printf "Enter a number of years or -1 to quit: "
	read years

	if [ $years -eq -1 ]
	then
		break
	fi
	go run main.go $years

done
