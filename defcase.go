package defcase

import (
	"github.com/gopher-summer/defcase/internal"
)

type defCaseFn func(pkgPath, name string) string

var tagNode *internal.PkgNode = internal.NewPkgNode()

func DefCase(tag, pkgPath string, c Case) {
	tagNode.SetCase(tag, pkgPath, c)
}

func ToDefCase(tag, pkgPath, name string) string {
	c := tagNode.GetCase(tag, pkgPath)

	return Convert(name, c)
}

func DefCaseFn(tag string) defCaseFn {
	pkgNode := tagNode.GetOrSetSubNode(tag)

	return func(pkgPath, name string) string {
		c := pkgNode.GetCaseByPkg(pkgPath)

		return Convert(name, c)
	}
}
