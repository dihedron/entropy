NAME := entropy
DESCRIPTION := A tool to calculate the entropy of a bunch of data.
COPYRIGHT := 2024 © Andrea Funtò
LICENSE := AGPL 3.0
LICENSE_URL := https://www.gnu.org/licenses/agpl-3.0.en.html
VERSION_MAJOR := 0
VERSION_MINOR := 0
VERSION_PATCH := 2
VERSION=$(VERSION_MAJOR).$(VERSION_MINOR).$(VERSION_PATCH)
MAINTAINER=dihedron.dev@gmail.com
VENDOR=dihedron.dev@gmail.com
PRODUCER_URL=https://github.com/dihedron/
DOWNLOAD_URL=$(PRODUCER_URL)${NAME}
METADATA_PACKAGE=$$(grep "module .*" go.mod | sed 's/module //gi')/version

_RULES_MK_MINIMUM_VERSION=202408011410
#_RULES_MK_ENABLE_CGO=1
#_RULES_MK_ENABLE_GOGEN=1

include rules.mk
