package devbox

// Manifest contains a map of ManifestItem values by their type key.
type Manifest = map[ManifestType][]ManifestItem

// ManifestType identifies the type of ManifestItem values.
type ManifestType = string

// ManifestItem contains optional path and commands to use when setting up
// that item.
type ManifestItem struct {
	Path     string   `yaml:"path,omitempty"`
	Commands []string `yaml:"commands,omitempty"`
}

const breakCommand = "break"

// ManifestTypes contains the list of defaultManifest types in install order.
var ManifestTypes = []string{
	"bash",
	"zsh",
	"git",
	"ssh",
	"tmux",
	"vim",
	"emacs",
}

var emptyCommands = make([]string, 0)

var defaultManifest = Manifest{
	"bash": []ManifestItem{
		{
			Path:     "~/.bash_profile",
			Commands: emptyCommands,
		},
		{
			Path:     "~/.bashrc",
			Commands: emptyCommands,
		},
		{
			Path:     "~/.profile",
			Commands: emptyCommands,
		},
	},
	"zsh": []ManifestItem{
		{
			Path:     "~/.zshrc",
			Commands: emptyCommands,
		},
		{
			Path:     "~/.zshenv",
			Commands: emptyCommands,
		},
		{
			Path:     "~/.zlogin",
			Commands: emptyCommands,
		},
		{
			Path:     "~/.zlogout",
			Commands: emptyCommands,
		},
		{
			Path:     "~/.zprofile",
			Commands: emptyCommands,
		},
		{
			Path:     "~/.oh-my-zsh/",
			Commands: emptyCommands,
		},
	},
	"git": []ManifestItem{
		{
			Path:     "~/.gitconfig",
			Commands: emptyCommands,
		},
		{
			Path:     "~/.gitignore",
			Commands: emptyCommands,
		},
		{
			Path:     "~/.gitattributes",
			Commands: emptyCommands,
		},
	},
	"ssh": []ManifestItem{
		{
			Path:     "~/.ssh/",
			Commands: emptyCommands,
		},
	},
	"tmux": []ManifestItem{
		{
			Path:     "~/.tmux.conf",
			Commands: emptyCommands,
		},
		{
			Path:     "~/.tmux/",
			Commands: emptyCommands,
		},
	},
	"vim": []ManifestItem{
		{
			Path:     "~/.vimrc",
			Commands: emptyCommands,
		},
		{
			Path:     "~/.viminfo",
			Commands: emptyCommands,
		},
		{
			Path:     "~/.vim/",
			Commands: emptyCommands,
		},
	},
	"emacs": []ManifestItem{
		{
			Path: "~/.spacemacs",
			Commands: []string{
				"git clone https://github.com/syl20bnr/spacemacs /home/{box.User}/.emacs.d",
				breakCommand,
			},
		},
		{
			Path: "~/.doom.d/",
			Commands: []string{
				"git clone https://github.com/hlissner/doom-emacs /home/{box.User}/.emacs.d",
				breakCommand,
			},
		},
		{
			Path:     "~/.emacs.d/",
			Commands: emptyCommands,
		},
	},
}
