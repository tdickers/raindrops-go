## Golang raindrops

Research into how to consume netlink inetdiag info to get metrics like Ruby's raindrops gem provides.

Current code is workable, rather than other libraries that were less usable but could be used to fetch only the desired data.

Ideal is sending inetdiag messages asking only for the desired data, for the desired bind addresses.

# Resources
https://veithen.io/2014/01/01/how-tcp-backlog-works-in-linux.html
https://yhbt.net/raindrops.git/tree/lib/raindrops