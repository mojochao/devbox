package devbox

// Manifest contains a map of ManifestItem values by their type key.
type Manifest = map[ManifestType][]ManifestItem

// ManifestType identifies the type of ManifestItem values.
type ManifestType = string

// ManifestItem contains optional path and commands to use when setting up
// that item.
type ManifestItem struct {
	Path     string
	Commands []string
}

const done = "done"

// ManifestTypes contains the list of manifest types in install order.
var ManifestTypes = []string{
	"bash",
	"zsh",
	"git",
	"ssh",
	"tmux",
	"vim",
	"emacs",
}

var manifest = Manifest{
	"bash": []ManifestItem{
		{Path: "~/.bash_profile"},
		{Path: "~/.bashrc"},
		{Path: "~/.profile"},
	},
	"zsh": []ManifestItem{
		{Path: "~/.zshrc"},
		{Path: "~/.zshenv"},
		{Path: "~/.zlogin"},
		{Path: "~/.zlogout"},
		{Path: "~/.zprofile"},
		{Path: "~/.oh-my-zsh/"},
	},
	"git": []ManifestItem{
		{Path: "~/.gitconfig"},
		{Path: "~/.gitignore"},
		{Path: "~/.gitattributes"},
	},
	"ssh": []ManifestItem{
		{Path: "~/.ssh/"},
	},
	"tmux": []ManifestItem{
		{Path: "~/.tmux.conf"},
		{Path: "~/.tmux/"},
	},
	"vim": []ManifestItem{
		{Path: "~/.vimrc"},
		{Path: "~/.viminfo"},
		{Path: "~/.vim/"},
	},
	"emacs": []ManifestItem{
		{
			Path: "~/.spacemacs",
			Commands: []string{
				"git clone https://github.com/syl20bnr/spacemacs /home/{box.User}/.emacs.d",
				done,
			},
		},
		{
			Path: "~/.doom.d/",
			Commands: []string{
				"git clone https://github.com/hlissner/doom-emacs /home/{box.User}/.emacs.d",
				done,
			},
		},
		{
			Path: "~/.emacs.d/",
		},
	},
}
