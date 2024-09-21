grep -o '" [0-9]\{3\} ' dummy.log | awk '{print $2}' #grep extract the 3 digit number with " and awk prints the second element
