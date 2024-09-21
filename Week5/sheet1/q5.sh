jq -r '.[].url' images.json | xargs -n 1 -P 4 wget -q # xargs picks 1 url at a time from n urls and -P 4 runs upto 4 download task in parallel wget download quitly
