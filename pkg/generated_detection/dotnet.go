package generated_detection

import (
	"regexp"
	"strings"
)

// isDotNetDocFile Is this a generated documentation file for a .NET assembly?
//
// .NET developers often check in the XML Intellisense file along with an
// assembly - however, these don't have a special extension, so we have to
// dig into the contents to determine if it's a docfile. Luckily, these files
// are extremely structured, so recognizing them is easy.
func (d detector) isDotNetDocFile() bool {
	if !d.hasExtension(".xml") || d.Lines.Len() < 3 {
		return false
	}

	lines := d.Lines.All()

	secondLine := lines[1]
	thirdLine := lines[2]
	secondLastLine := lines[len(lines)-2]

	// .NET Docfiles always open with <doc> and their first tag is an
	// <assembly> tag
	return strings.Contains(secondLine, "<doc>") &&
		strings.Contains(thirdLine, "<assembly>") &&
		strings.Contains(secondLastLine, "</doc>")
}

// isDotNetDesignerFile checks if the file is a codegen file for a .NET project
//
// Visual Studio often uses code generation to generate partial classes, and
// these files can be quite unwieldy. Let's hide them.
func (d detector) isDotNetDesignerFile() bool {
	re := regexp.MustCompile("(?i)\\.designer\\.(cs|vb)$")

	return re.MatchString(d.FileName.Base)
}

// isDotNetSpecFlowFeatureFile checks if the file is a codegen file for Specflow feature file
//
// Visual Studio's SpecFlow extension generates *.feature.cs files
// from *.feature files, they are not meant to be consumed by humans.
// Let's hide them.
func (d detector) isDotNetSpecFlowFeatureFile() bool {
	re := regexp.MustCompile("(?i)\\.feature\\.cs$")

	return re.MatchString(d.FileName.Base)
}
