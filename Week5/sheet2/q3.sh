find ./ -type f | awk -F. '{ext[$NF]++} END {for (e in ext) print e, ext[e]}'
