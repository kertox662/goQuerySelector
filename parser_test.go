package goqs

import "testing"

func TestParseHTML(t *testing.T) {
	htmlStr := "<html><head></head><body><header><div id=\"left div\" class=\"class\"><img id = \"inner\" class=\"pretty cool class\"></div><div id=\"center\"><h1>Website</h1></div><section></section><div id=\"right\" class = \"pretty class\"><p>It's so goo</p></div><img id=\"logo\"></header><main></main><footer></footer></body></html>"
	ParseHTML(htmlStr)
}