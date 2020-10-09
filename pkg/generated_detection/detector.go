package generated_detection

import (
	"path/filepath"

	linesinternal "github.com/dominikeinkemmer/generated-file-detection/pkg/lines"
)

type fileName struct {
	Full string
	Base string
	Ext  string
}

type detector struct {
	FileName *fileName
	Lines linesinternal.Line
}

func New() detector {
	return detector{}
}

func (d detector) IsGenerated(filePath string) (bool, error) {
	d.FileName = &fileName{
		Full: filePath,
		Base: filepath.Base(filePath),
		Ext:  filepath.Ext(filePath),
	}

	// Do all name based checks first
	isGenerated, err := d.isGeneratedByName()
	if isGenerated || err != nil {
		return isGenerated, err
	}

	// If name based checks did not result in a detection, read the file and pass it down
	lines := linesinternal.New(filePath)

	err = lines.ReadFile()

	if err != nil {
		return false, err
	}

	d.Lines = lines

	isGenerated = d.isMinified() ||
		d.hasSourceMap() ||
		d.isSourceMap() ||
		d.isCompiledCoffeeScript() ||
		d.isDotNetDocFile()

	return isGenerated, nil
}

func (d detector) isGeneratedByName() (bool, error) {
	isGenerated, err := d.isXcodeFile()

	if isGenerated {
		return isGenerated, err
	}

	isGenerated, err = d.isCocoaPods()

	if isGenerated {
		return isGenerated, err
	}

	isGenerated, err = d.isCarthageBuild()

	if isGenerated {
		return isGenerated, err
	}

	return d.isDotNetDesignerFile() || d.isDotNetSpecFlowFeatureFile(), nil
}

// isXcodeFile checks if the file is an xcode file by checking the file extension
func (d detector) isXcodeFile() (bool, error) {
	extMap := map[string]bool{
		".nib":             true,
		".xcworkspacedata": true,
		".xcuserstate":     true,
	}

	return extMap[d.FileName.Ext], nil
}

// isCocoaPods checks if the file is part of Pods/,
// which contains dependencies not meant for humans in pull requests.
func (d detector) isCocoaPods() (bool, error) {
	return filepath.Match("(^Pods|/Pods)/", d.FileName.Base)
}

// isCarthageBuild checks if the file is part of Carthage/Build/,
// which contains dependencies not meant for humans in pull requests.
func (d detector) isCarthageBuild() (bool, error) {
	return filepath.Match("(^|/)Carthage/Build/", d.FileName.Base)
}
