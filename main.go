package main

import "fmt"

func main() {
	fmt.Println(getURLsFromHTML(`
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`, "https://blog.boot.dev"))
}
