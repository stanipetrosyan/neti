package external

type FileReader struct {
	Path string
}

func NewFileReader(path string) *FileReader {
	return &FileReader{
		Path: path,
	}
}

func (f *FileReader) ReadTheFile() (string, error) {
	// Implementation kept simple on purpose
	return "reading a file", nil
}
