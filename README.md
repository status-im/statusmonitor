

```
df = read.csv("~/Downloads/data.csv")
df$timestamp = as.POSIXct(df$timestamp, origin="1970-01-01")
```

Draw the plot;
```
plot(df, type="l", main="status CPU usage", col = 'green', lwd=2, cex.main=1.5, cex.axis=1, xlab="time", xaxt="n")
axis.POSIXct(1, df$timestamp, format="%H:%M:%S")
```

or interactive version with plotly:

```
library(plotly)
plot_ly(df, x = df$timestamp, y = df$cpu, type = 'scatter', mode = 'lines', fill = 'tozeroy')
```
