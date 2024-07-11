package gqlgen_plugin

import (
	_ "embed"
	"log"
	"path/filepath"
	"strings"

	"github.com/99designs/gqlgen/codegen"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/99designs/gqlgen/plugin"

	_ "github.com/darashevcstbg/gqlgen/pkg/meetups"
)

//go:embed meetup.resolvers.gotpl
var resolversTemplate string

type GqlgenPlugin struct {
	filename        string
	clientModelPath string
}

func New(filename string, clientModelPath string) *GqlgenPlugin {
	return &GqlgenPlugin{filename: filename, clientModelPath: clientModelPath}
}

var (
	_ plugin.CodeGenerator = &GqlgenPlugin{}
	_ plugin.ConfigMutator = &GqlgenPlugin{}
)

func (m *GqlgenPlugin) Name() string {
	return "gqlgen_plugin"
}

func (m *GqlgenPlugin) MutateConfig(cfg *config.Config) error {
	return nil
}

func (m *GqlgenPlugin) GenerateCode(data *codegen.Data) error {
	log.Printf("GenerateCode called with filename: %s", m.filename)

	abs, err := filepath.Abs(m.filename)
	if err != nil {
		log.Printf("Error getting absolute path: %v", err)
		return err
	}
	pkgName := nameForDir(filepath.Dir(abs))
	// pkgName = "main"

	resolverBuild := &ResolverBuild{
		Data:            data,
		ResolverType:    "Resolver",
		PackageName:     pkgName,
		ClientModelPath: m.clientModelPath,
	}

	return templates.Render(templates.Options{
		PackageName:     pkgName,
		Filename:        m.filename,
		Data:            resolverBuild,
		GeneratedHeader: true,
		Packages:        data.Config.Packages,
		Template:        resolversTemplate,
	})
}

type ResolverBuild struct {
	*codegen.Data
	ResolverType    string
	PackageName     string
	ClientModelPath string
}

// nameForDir returns a valid package name for the given directory
func nameForDir(dir string) string {
	base := filepath.Base(dir)
	return strings.Map(func(r rune) rune {
		if r == '-' {
			return '_'
		}
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' {
			return r
		}
		return -1
	}, base)
}
