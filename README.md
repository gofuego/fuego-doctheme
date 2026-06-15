# fuego-doctheme

Shared look-and-feel for the [Fuego](https://github.com/gofuego/fuego)
documentation sites, packaged as two Fuego **format packs**.

```go
import "github.com/gofuego/fuego-doctheme"

eng.Use(doctheme.Public())     // user-facing docs (hero home, accent buttons)
eng.Use(doctheme.Blueprint())  // private *-internal maintainer docs (muted, banner)
```

Each pack carries an embedded theme FS — `base.html`, `layouts/`, `renderers/`,
and a `static/` subtree (the stylesheet). The engine layers it
*under* the consuming site's own `theme/` directory, so a site supplies only the
genuinely site-specific bits (its `partials/topbar.html` and a curated
`partials/sidebar.html`) plus its content, and inherits the rest. Anything a site
puts in its own `theme/` wins over the pack.

| Pack | Look | Used by |
|------|------|---------|
| `Public()` | Marketing-adjacent: hero home with envelope-driven CTA buttons, accent links, RSS. | `fuego-docs`, `fuego-studio-docs` |
| `Blueprint()` | Muted monospace "maintainer documentation" with an *Internal* banner, no marketing chrome. | `fuego-docs-internal`, `fuego-studio-docs-internal` |

## Conventions the theme reads

- `layout: home` — hero (Public) / plain heading (Blueprint).
- `layout: doc` — a documentation article; `summary` renders as a lede, `tags` as chips (Blueprint).
- `nav_section` + `nav_weight` — drive a curated sidebar (each site overrides `partials/sidebar.html`).
- Home page `cta:` envelope list (`label`, `url`, optional `ghost: true`) — Public hero buttons.

This module is intentionally tiny and stdlib-plus-`core` only.
