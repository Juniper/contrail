package contrailutil

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	ContrailUtil.AddCommand(packageCmd)
}

var (
	pkgArch string
	// deb & rpm does not support semver so have to handle their version a little differently
	linuxPackageVersion   = "0.0.1"
	linuxPackageIteration string
	binaries              = []string{"contrail"}
)

const (
	errorGetWorkingDirectory = "Cannot get working directory"
)

// createLinuxPackages is based on Grafana project.
func createLinuxPackages() {
	if err := clean(); err != nil {
		log.Fatalf("Cannot clean workspace: %s", err)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("%s: %s", errorGetWorkingDirectory, err)
	}

	distDir := filepath.Join(wd, "dist")
	if err := os.MkdirAll(distDir, 0700); err != nil {
		log.Fatalf("Cannot create dist directory: %s", err)
	}

	outputDir := filepath.Join(wd, "tmp/bin/")
	runPrint("gox", "-osarch=linux/amd64", "--output", filepath.Join(outputDir, "contrail"), "./cmd/contrail")
	runPrint("gox", "-osarch=linux/amd64", "--output", filepath.Join(outputDir, "contrailutil"), "./cmd/contrailutil")
	createDebPackages()
	createRpmPackages()
}

func rmrf(paths ...string) error {
	for _, path := range paths {
		log.Println("rm -rf", path)
		if err := os.RemoveAll(path); err != nil {
			return err
		}
	}

	return nil
}

func clean() error {
	if err := rmrf("dist"); err != nil {
		return fmt.Errorf("remove dist directory: %s", err)
	}
	if err := rmrf("tmp"); err != nil {
		return fmt.Errorf("remove tmp directory: %s", err)
	}
	return nil
}

type linuxPackageOptions struct {
	packageType            string
	homeDir                string
	binPath                string
	configDir              string
	etcDefaultPath         string
	etcDefaultFilePath     string
	initdScriptFilePath    string
	systemdServiceFilePath string

	postinstSrc    string
	initdScriptSrc string
	defaultFileSrc string
	systemdFileSrc string

	depends []string
}

func createPackage(options linuxPackageOptions) {
	packageRoot, err := ioutil.TempDir("", "contrail-linux-pack")
	if err != nil {
		log.Fatalf("Cannot create temporary directory: %s", err)
	}

	homeDir := filepath.Join(packageRoot, options.homeDir)
	configDir := filepath.Join(packageRoot, options.configDir)
	// create directories
	runPrint("mkdir", "-p", homeDir)
	runPrint("mkdir", "-p", configDir)
	runPrint("mkdir", "-p", filepath.Join(packageRoot, "/etc/init.d"))
	runPrint("mkdir", "-p", filepath.Join(packageRoot, options.etcDefaultPath))
	runPrint("mkdir", "-p", filepath.Join(packageRoot, "/usr/lib/systemd/system"))
	runPrint("mkdir", "-p", filepath.Join(packageRoot, "/usr/sbin"))

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("%s: %s", errorGetWorkingDirectory, err)
	}

	// copy binary
	for _, binary := range binaries {
		runPrint(
			"cp",
			"-p",
			filepath.Join(wd, "tmp/bin/"+binary), filepath.Join(packageRoot, "/usr/sbin/"+binary),
		)
	}

	// copy init.d script
	runPrint("cp", "-p", options.initdScriptSrc, filepath.Join(packageRoot, options.initdScriptFilePath))
	// copy environment var file
	runPrint("cp", "-p", options.defaultFileSrc, filepath.Join(packageRoot, options.etcDefaultFilePath))
	// copy systemd file
	runPrint("cp", "-p", options.systemdFileSrc, filepath.Join(packageRoot, options.systemdServiceFilePath))
	// copy release files
	runPrint("cp", "-a", filepath.Join(wd, "tmp")+"/.", filepath.Join(packageRoot, options.homeDir))
	// copy json schema
	runPrint("cp", "-rp", filepath.Join(wd, "public"), homeDir)
	// default config
	runPrint("cp", "-p", filepath.Join(wd, "tools", "init.sql"), homeDir)
	runPrint("cp", "-p", filepath.Join(wd, "packaging", "apisrv.yml"), homeDir)
	// remove bin path
	runPrint("rm", "-rf", filepath.Join(packageRoot, options.homeDir, "bin"))

	args := []string{
		"-s", "dir",
		"--description", "contrail",
		"-C", packageRoot,
		"--vendor", "Juniper Networks",
		"--url", "https://juniper.net",
		"--license", "\"Apache 2.0\"",
		"--maintainer", "nueno@juniper.net",
		"--config-files", options.initdScriptFilePath,
		"--config-files", options.etcDefaultFilePath,
		"--config-files", options.systemdServiceFilePath,
		"--after-install", options.postinstSrc,
		"--name", "contrail",
		"--version", linuxPackageVersion,
		"-p", "./dist",
	}

	if options.packageType == "rpm" {
		args = append(args, "--rpm-posttrans", "packaging/rpm/control/posttrans")
		args = append(args, "--rpm-os", "linux")
	}

	if options.packageType == "deb" {
		args = append(args, "--deb-no-default-config-files")
	}

	if pkgArch != "" {
		args = append(args, "-a", pkgArch)
	}

	if linuxPackageIteration != "" {
		args = append(args, "--iteration", linuxPackageIteration)
	}

	// add dependenciesj
	for _, dep := range options.depends {
		args = append(args, "--depends", dep)
	}

	args = append(args, ".")

	fmt.Println("Creating package: ", options.packageType)
	runPrint("fpm", append([]string{"-t", options.packageType}, args...)...)
}

func createDebPackages() {
	createPackage(linuxPackageOptions{
		packageType:            "deb",
		homeDir:                "/usr/share/contrail",
		binPath:                "/usr/sbin",
		configDir:              "/etc/contrail",
		etcDefaultPath:         "/etc/default",
		etcDefaultFilePath:     "/etc/default/contrail",
		initdScriptFilePath:    "/etc/init.d/contrail",
		systemdServiceFilePath: "/usr/lib/systemd/system/contrail.service",

		postinstSrc:    "packaging/deb/control/postinst",
		initdScriptSrc: "packaging/deb/init.d/contrail",
		defaultFileSrc: "packaging/deb/default/contrail",
		systemdFileSrc: "packaging/deb/systemd/contrail.service",

		depends: []string{"adduser"},
	})
}

func createRpmPackages() {
	createPackage(linuxPackageOptions{
		packageType:            "rpm",
		homeDir:                "/usr/share/contrail",
		binPath:                "/usr/sbin",
		configDir:              "/etc/contrail",
		etcDefaultPath:         "/etc/sysconfig",
		etcDefaultFilePath:     "/etc/sysconfig/contrail",
		initdScriptFilePath:    "/etc/init.d/contrail",
		systemdServiceFilePath: "/usr/lib/systemd/system/contrail.service",

		postinstSrc:    "packaging/rpm/control/postinst",
		initdScriptSrc: "packaging/rpm/init.d/contrail",
		defaultFileSrc: "packaging/rpm/sysconfig/contrail",
		systemdFileSrc: "packaging/rpm/systemd/contrail.service",

		depends: []string{"/sbin/service"},
	})
}

func runPrint(cmd string, args ...string) {
	log.Println(cmd, strings.Join(args, " "))
	ecmd := exec.Command(cmd, args...)
	ecmd.Stdout = os.Stdout
	ecmd.Stderr = os.Stderr
	err := ecmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "make a deb and rpm package",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		createLinuxPackages()
	},
}
