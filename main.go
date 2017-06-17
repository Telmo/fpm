package main

import (
	"os"
	"time"

	"github.com/urfave/cli"
)

const (
	// VERSION is the release version
	VERSION = "v0.0.0-alpha"
)

func main() {
	app := cli.NewApp()
	app.Name = "gfpm"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Telmo",
			Email: "no-reply@example.com",
		},
	}
	app.Compiled = time.Now()
	app.Version = VERSION
	app.Usage = "golang version of Jordan Sisselˈs FPM"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "output-type, t",
			Usage: "the type of package you want to create (deb, rpm, solaris, etc)",
		},
		cli.StringFlag{
			Name:  "input-type, s",
			Usage: "the package type to use as input (gem, rpm, python, etc)",
		},
		cli.StringFlag{
			Name:  "chdir, C",
			Usage: "Change directory to here before searching for files",
		},
		cli.StringFlag{
			Name:  "prefix",
			Usage: "A path to prefix files with when building the target package. This may not be necessary for all input packages. For example, the 'gem' type will prefix with your gem directory automatically.",
		},
		cli.StringFlag{
			Name:  "package, p",
			Usage: "The package file path to output.",
		},
		cli.StringFlag{
			Name:  "force, f",
			Usage: "Force output even if it will overwrite an existing file",
		},
		cli.StringFlag{
			Name:  "name, n",
			Usage: "The name to give to the package",
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "Enable debug output",
		},
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Enable verbose output",
		},
		cli.BoolFlag{
			Name:  "debug-workspace",
			Usage: "Keep any file workspaces around for debugging. This will disable automatic cleanup of package staging and build paths. It will also print which directories are available.",
		},
		cli.StringFlag{
			Name:  "iteration",
			Usage: "The iteration to give to the package. RPM calls this the 'release'. FreeBSD calls it 'PORTREVISION'. Debian calls this 'debian_revision'",
		},
		cli.StringFlag{
			Name:  "epoch",
			Usage: "The epoch value for this package. RPM and Debian calls this 'epoch'. FreeBSD calls this 'PORTEPOCH'",
		},
		cli.StringFlag{
			Name:  "license",
			Usage: "(optional) license name for this package",
		},
		cli.StringFlag{
			Name:  "vendor",
			Usage: "(optional) vendor name for this package",
		},
		cli.StringFlag{
			Name:  "category",
			Usage: "(optional) category this package belongs to",
		},
		cli.StringSliceFlag{
			Name:  "depends, d",
			Usage: "A dependency. This flag can be specified multiple times. Value is usually in the form of: -d 'name' or -d 'name > version'",
		},
		cli.BoolFlag{
			Name:  "no-depends",
			Usage: "Do not list any dependencies in this package",
		},
		cli.BoolFlag{
			Name:  "no-auto-depends",
			Usage: "Do not list any dependencies in this package automatically",
		},
		cli.StringSliceFlag{
			Name:  "provides",
			Usage: "What this package provides (usually a name). This flag can be specified multiple times.",
		},
		cli.StringSliceFlag{
			Name:  "conflicts",
			Usage: "Other packages/versions this package conflicts with. This flag can be specified multiple times.",
		},
		cli.StringSliceFlag{
			Name:  "replaces",
			Usage: "Other packages/versions this package replaces. Equivalent of rpm's 'Obsoletes'. This flag can be specified multiple times.",
		},
		cli.StringFlag{
			Name:  "config-files",
			Usage: "Mark a file in the package as being a config file. This uses 'conffiles' in debs and %config in rpm. If you have multiple files to mark as configuration files, specify this flag multiple times.  If argument is directory all files inside it will be recursively marked as config files.",
		},
		cli.StringSliceFlag{
			Name:  "directories",
			Usage: "Recursively mark a directory as being owned by the package. Use this flag multiple times if you have multiple directories and they are not under the same parent directory ",
		},
		cli.StringFlag{
			Name:  "architecture, a",
			Usage: "The architecture name. Usually matches 'uname -m'. For automatic values, you can use '-a all' or '-a native'. These two strings will be translated into the correct value for your platform and target package type.",
		},
		cli.StringFlag{
			Name:  "maintainer, m",
			Usage: "The maintainer of this package.",
		},
		cli.StringFlag{
			Name:  "package-name-suffix, S",
			Usage: "a name suffix to append to package and dependencies.",
		},
		cli.BoolFlag{
			Name:  "edit, e",
			Usage: "Edit the package spec before building.",
		},
		cli.StringSliceFlag{
			Name:  "exclude, x",
			Usage: "Exclude paths matching pattern (shell wildcard globs valid here). If you have multiple file patterns to exclude, specify this flag excludes",
		},
		cli.StringFlag{
			Name:  "exclude-file",
			Usage: "The path to a file containing a newline-sparated list of patterns to exclude from input.",
		},
		cli.StringFlag{
			Name:  "description",
			Usage: "Add a description for this package. You can include '\\n' sequences to indicate newline breaks.",
		},
		cli.StringFlag{
			Name:  "url",
			Usage: "Add a url for this package.",
		},
		cli.StringFlag{
			Name:  "after-install",
			Usage: "A script to be run after package installation",
		},
		cli.StringFlag{
			Name:  "before-install",
			Usage: "A script to be run before package installation",
		},
		cli.StringFlag{
			Name:  "after-remove",
			Usage: "A script to be run after package removal",
		},
		cli.StringFlag{
			Name:  "before-remove",
			Usage: "A script to be run before package removal",
		},
		cli.StringFlag{
			Name:  "after-upgrade",
			Usage: "A script to be run after package upgrade. If not specified,\n --before-install, --after-install, --before-remove, and \n-after-remove will behave in a backwards-compatible manner\n(they will not be upgrade-case aware).\nCurrently only supports deb, rpm and pacman packages.",
		},
		cli.StringFlag{
			Name:  "before-upgrade",
			Usage: "A script to be run before package upgrade. If not specified,\n --before-install, --after-install, --before-remove, and \n --after-remove will behave in a backwards-compatible manner\n (they will not be upgrade-case aware).\n Currently only supports deb, rpm and pacman packages.",
		},
		cli.StringFlag{
			Name:  "workdir",
			Usage: "The directory you want fpm to do its work in, where 'work' is any file copying, downloading, etc. Roughly any scratch space fpm needs to build your package.",
		},
		cli.StringFlag{
			Name:  "template-scripts",
			Usage: "(NOT SUPPORTED), Allow scripts to be templated. This lets you use ERB to template your packaging scripts (for --after-install, etc). For example, you can do things like <%= name %> to get the package name. For more information, see the fpm wiki: https://github.com/jordansissel/fpm/wiki/Script-Templates",
		},
		cli.StringFlag{
			Name:  "template-value",
			Usage: "(NOT SUPPORTED), Make 'key' available in script templates, so <%= key %> given will be the provided value. Implies --template-scripts",
		},
	}

	app.Run(os.Args)
}
