HTMX is a JS library that detects new attributes that can be set on HTML element (like a HTML extension).
Specifically it aims to simplify the process of fetching data (HTTP AJAX requests) and rendering it.
2 usage options:
- HTMX in a decoupled frontend: like React, however HTMX focuses on updating the UI with AJAX requests.
- HTMX in an integrated fullstack app: best use-case is with a tightly connected backend (that renders pages). Indeed HTMX expects HTML code as response content.

# install
- Add a CDN link.
- Download HTMX and host it on the server, then import and load it from there (https://htmx.org/docs/#download-a-copy).

# hx-trigger
hx-trigger default is `submit` for a form and `click` for most other elements.

It can also be used to change how often the event triggers the request, or some other element should trigger it, or some "indirect" event (e.g., the element being scrolled into view) should trigger it.

`hx-trigger="click once"` to ensure that only the first click on a button loads data.
`hx-trigger="click delay:1s"` to add a delay until the request is sent.
Change the element that triggers the request:
```html
<button>Load Data</button>
<div hx-get="/data" hx-trigger="click from:previous button">
    Nothing there :(
</div>
```
Trigger the request once the element is scrolled into the viewport:
```html
<div hx-get="/data" hx-trigger="revealed">
    Nothing there :(
</div>
```
https://htmx.org/docs/#triggers modifiers & variations in the official docs.

# hx-select
When we wish to only use a fragment of the html page received, we can target that fragment with `hx-select`. Then the other directives (like `hx-target`) will only use that HTML to do their action.