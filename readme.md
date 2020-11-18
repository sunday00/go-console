# install
```
go get -u -v github.com/sunday00/go-console
```

# import 
```
import (
	"github.com/sunday00/go-console"
)
```

# usage

## colored
```
console.PrintColored("apple", console.Danger)
```

<div style="background-color:black; padding:1rem;">
$>> <span style="color:red">apple</span>
</div>

with Ln func is break line
```
console.PrintColoredLn("apple", console.Danger)
```

## emphasize
```
console.PrintEmp("banana")
```
<div style="background-color:black; padding:1rem;">
$>>

&#61;&#61;&#61;&#61;&#61;&#61;&#61;&#61;&#61;&#61;&#61;&#61;&#61;&#61;&#61;&#61;

banana

&#61;&#61;&#61;&#61;&#61;&#61;&#61;&#61;&#61;&#61;&#61;&#61;&#61;&#61;&#61;&#61;

</div>

# color types
<p style="color:white;">White</p>
<p style="color:black; background-color:white;">Black</p>
<p style="color:gold;">Warning</p>
<p style="color:cyan;">Info</p>
<p style="color:darkOrange;">Danger</p>
<p style="color:red;">Panic</p>
<p style="color:magenta;">Cute</p>
<p style="color:lightgreen;">Success</p>

