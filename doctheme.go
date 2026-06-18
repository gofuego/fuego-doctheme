// Package doctheme provides the shared look-and-feel for the Fuego
// documentation sites as two Fuego format packs.
//
// Each pack carries an embedded theme FS (base.html, layouts/, renderers/, and
// a static/ subtree with the stylesheet). A site consumes
// one with eng.Use(doctheme.Public()) or eng.Use(doctheme.Blueprint()) and
// supplies only the genuinely site-specific bits — its topbar and sidebar
// partials and its content. Anything a site puts in its own theme/ directory
// overrides the pack, because user theme files always win over pack themes.
//
//   - Public is the marketing-adjacent look for user-facing docs (hero home,
//     accent buttons, RSS link). Its home layout reads call-to-action buttons
//     from the home page's "cta" envelope list, so it stays site-agnostic.
//   - Blueprint is the muted, monospace "maintainer documentation" look for the
//     private *-internal sites, with an Internal banner and no marketing chrome.
package doctheme

import (
	"embed"
	"io/fs"

	"github.com/gofuego/fuego/core"
)

//go:embed public
var publicFS embed.FS

//go:embed blueprint
var blueprintFS embed.FS

// Public returns the user-facing documentation theme pack.
func Public() core.Pack {
	return core.Pack{Name: "doctheme-public", Theme: sub(publicFS, "public")}
}

// Blueprint returns the maintainer "internal docs" theme pack.
func Blueprint() core.Pack {
	return core.Pack{Name: "doctheme-blueprint", Theme: sub(blueprintFS, "blueprint")}
}

// sub roots an embedded theme FS at its variant directory so base.html sits at
// the FS root, as the engine's theme layering expects.
func sub(f embed.FS, dir string) fs.FS {
	sf, err := fs.Sub(f, dir)
	if err != nil {
		panic("doctheme: " + err.Error())
	}
	return sf
}
