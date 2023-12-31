package badge

import (
	"strings"
	"views-counter/src/formatter"
)

type Create func(int) string

func MakeCreate() Create {
	return func(message int) string {

		counter := formatter.ShortNumber(float64(message), 2)

		badge := `
			<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="113.5" height="28" role="img" aria-label="VIEWS: 3">
				<g shape-rendering="crispEdges">
					<rect width="81.25" height="28" fill="#640464"/>
					<rect x="81.25" width="32.25" height="28" fill="#7c007c"/>
				</g>
				<g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="100">
					<image x="9" y="7" width="14" height="14" xlink:href="data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIGZpbGw9IndoaXRlc21va2UiIHZlcnNpb249IjEuMSIgd2lkdGg9IjE2IiBoZWlnaHQ9IjE2IiB2aWV3Qm94PSIwIDAgMTYgMTYiIGNsYXNzPSJvY3RpY29uIG9jdGljb24tZXllIiBhcmlhLWhpZGRlbj0idHJ1ZSI+PHN0eWxlPiogeyBmaWxsOiB3aGl0ZSFpbXBvcnRhbnQ7IH08L3N0eWxlPjxwYXRoIGQ9Ik04IDJjMS45ODEgMCAzLjY3MS45OTIgNC45MzMgMi4wNzggMS4yNyAxLjA5MSAyLjE4NyAyLjM0NSAyLjYzNyAzLjAyM2ExLjYyIDEuNjIgMCAwIDEgMCAxLjc5OGMtLjQ1LjY3OC0xLjM2NyAxLjkzMi0yLjYzNyAzLjAyM0MxMS42NyAxMy4wMDggOS45ODEgMTQgOCAxNGMtMS45ODEgMC0zLjY3MS0uOTkyLTQuOTMzLTIuMDc4QzEuNzk3IDEwLjgzLjg4IDkuNTc2LjQzIDguODk4YTEuNjIgMS42MiAwIDAgMSAwLTEuNzk4Yy40NS0uNjc3IDEuMzY3LTEuOTMxIDIuNjM3LTMuMDIyQzQuMzMgMi45OTIgNi4wMTkgMiA4IDJaTTEuNjc5IDcuOTMyYS4xMi4xMiAwIDAgMCAwIC4xMzZjLjQxMS42MjIgMS4yNDEgMS43NSAyLjM2NiAyLjcxN0M1LjE3NiAxMS43NTggNi41MjcgMTIuNSA4IDEyLjVjMS40NzMgMCAyLjgyNS0uNzQyIDMuOTU1LTEuNzE1IDEuMTI0LS45NjcgMS45NTQtMi4wOTYgMi4zNjYtMi43MTdhLjEyLjEyIDAgMCAwIDAtLjEzNmMtLjQxMi0uNjIxLTEuMjQyLTEuNzUtMi4zNjYtMi43MTdDMTAuODI0IDQuMjQyIDkuNDczIDMuNSA4IDMuNWMtMS40NzMgMC0yLjgyNS43NDItMy45NTUgMS43MTUtMS4xMjQuOTY3LTEuOTU0IDIuMDk2LTIuMzY2IDIuNzE3Wk04IDEwYTIgMiAwIDEgMS0uMDAxLTMuOTk5QTIgMiAwIDAgMSA4IDEwWiI+PC9wYXRoPjwvc3ZnPg=="/>
					<text transform="scale(.1)" x="491.25" y="175" textLength="402.5" fill="#fff">VIEWS</text>
					<text id="counter" transform="scale(.1)" x="973.75" y="175" fill="#fff" font-weight="bold">{counter}</text>
				</g>
			</svg>
			`
		badge = strings.Replace(badge, "{counter}", counter, -1)

		return badge
	}
}
