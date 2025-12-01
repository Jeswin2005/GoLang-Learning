package filehandler

import (
	"linux_agent_framework/src/errors"
	"os"
	"path/filepath"
	"testing"
)

func TestRenameFileAndMoveFile(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(string) (string, string)
		expectedError bool
	}{
		{
			"SuccessfulRename/Move",
			func(dir string) (string, string) {
				oldPath := filepath.Join(dir, "old.txt")
				newPath := filepath.Join(dir, "new.txt")

				if err := os.WriteFile(oldPath, []byte("hello"), 0644); err != nil {
					t.Fatalf("failed to create test file: %v", err)
				}
				return oldPath, newPath
			},
			false,
		},
		{
			"SourceNotExists",
			func(dir string) (string, string) {
				oldPath := filepath.Join(dir, "not-exist.txt")
				newPath := filepath.Join(dir, "new.txt")
				return oldPath, newPath
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// RenameFile
			dir := t.TempDir()
			oldpath, newpath := tt.setup(dir)

			err := RenameFile(oldpath, newpath)

			if tt.expectedError {
				if err == errors.ERROR_NO_ERROR {
					t.Fatalf("Expected error, got no error")
				}
			} else {
				if err != errors.ERROR_NO_ERROR {
					t.Fatalf("expected no error, got %v", err)
				}
			}

			// MoveFile
			dir = t.TempDir()
			oldpath, newpath = tt.setup(dir)
			err = MoveFile(oldpath, newpath)

			if tt.expectedError {
				if err == errors.ERROR_NO_ERROR {
					t.Fatalf("Expected error, got no error")
				}
			} else {
				if err != errors.ERROR_NO_ERROR {
					t.Fatalf("expected no error, got %v", err)
				}
			}
		})
	}
}

func TestRemoveFile(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(string) string
		expectedError bool
	}{
		{
			"FileExists",
			func(dir string) string {
				path := filepath.Join(dir, "test.txt")
				err := os.WriteFile(path, []byte("hello"), 0644)
				if err != nil {
					t.Fatalf("Failed to create test file")
				}
				return path
			},
			false,
		},
		{
			"EmptyDirectory",
			func(dir string) string {
				return dir
			},
			false,
		},
		{
			"FileNotExists",
			func(dir string) string {
				path := filepath.Join(dir, "test.txt")
				return path
			},
			true,
		},
		{
			"DirectoryNotEmpty",
			func(dir string) string {
				path := filepath.Join(dir, "test.txt")
				err := os.WriteFile(path, []byte("hello"), 0644)
				if err != nil {
					t.Fatalf("Failed to create test file")
				}
				return dir
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			path := tt.setup(dir)

			err := RemoveFile(path)
			if tt.expectedError {
				if err == errors.ERROR_NO_ERROR {
					t.Fatalf("Expected error, got no error")
				}
			} else {
				if err != errors.ERROR_NO_ERROR {
					t.Fatalf("expected no error, got %v", err)
				}
			}
		})
	}
}

func TestRemoveDirectory(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(string) string
		expectedError bool
	}{
		{
			"FileExists",
			func(dir string) string {
				path := filepath.Join(dir, "test.txt")
				err := os.WriteFile(path, []byte("hello"), 0644)
				if err != nil {
					t.Fatalf("Failed to create test file")
				}
				return path
			},
			false,
		},
		{
			"EmptyDirectory",
			func(dir string) string {
				return dir
			},
			false,
		},
		{
			"PathNotExists",
			func(dir string) string {
				path := filepath.Join(dir, "test.txt")
				return path
			},
			false,
		},
		{
			"DirectoryWithSubFolders",
			func(dir string) string {
				path := filepath.Join(dir, "sub1", "sub2")
				err := os.MkdirAll(path, 0755)
				if err != nil {
					t.Fatalf("Failed to create subdirectories: %v", err)
				}
				return dir
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			path := tt.setup(dir)

			err := RemoveDirectory(path)
			if tt.expectedError {
				if err == errors.ERROR_NO_ERROR {
					t.Fatalf("Expected error, got no error")
				}
			} else {
				if err != errors.ERROR_NO_ERROR {
					t.Fatalf("expected no error, got %v", err)
				}
			}
		})
	}
}

func TestIsPathExists(t *testing.T) {
	tests := []struct {
		name           string
		setup          func(string) string
		expectedOutput bool
	}{
		{
			"PathExists",
			func(dir string) string {
				path := filepath.Join(dir, "test.txt")
				err := os.WriteFile(path, []byte("hello"), 0644)
				if err != nil {
					t.Fatalf("Failed to create test file")
				}
				return path
			},
			true,
		},
		{
			"PathNotExists",
			func(dir string) string {
				path := filepath.Join(dir, "test.txt")
				return path
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			path := tt.setup(dir)

			actualOutput := IsPathExists(path)
			if actualOutput != tt.expectedOutput {
				t.Errorf("Expected: %t, Actual: %t", tt.expectedOutput, actualOutput)
			}
		})
	}
}

func TestReadDirectory(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(string) string
		expectedError bool
	}{
		{
			"EmptyDirectory",
			func(dir string) string {
				return dir
			},
			false,
		},
		{
			"DirectoryNotExists",
			func(dir string) string {
				return filepath.Join(dir, "not-exist")
			},
			true,
		},
		{
			"DirectoryWithSubFolders",
			func(dir string) string {
				path := filepath.Join(dir, "sub1", "sub2")
				err := os.MkdirAll(path, 0755)
				if err != nil {
					t.Fatalf("Failed to create subdirectories: %v", err)
				}
				return dir
			},
			false,
		},
		{
			"PathIsFile",
			func(dir string) string {
				file := filepath.Join(dir, "file.txt")
				if err := os.WriteFile(file, []byte("hello"), 0644); err != nil {
					t.Fatalf("Failed to create file: %v", err)
				}
				return file
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			path := tt.setup(dir)

			_, err := ReadDirectory(path)
			if tt.expectedError {
				if err == errors.ERROR_NO_ERROR {
					t.Fatalf("Expected error, got no error")
				}
			} else {
				if err != errors.ERROR_NO_ERROR {
					t.Fatalf("expected no error, got %v", err)
				}
			}
		})
	}
}

func TestChangeWorkingDirectory(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(string) string
		expectedError bool
	}{
		{
			"EmptyDirectory",
			func(dir string) string {
				return dir
			},
			false,
		},
		{
			"DirectoryNotExists",
			func(dir string) string {
				return filepath.Join(dir, "not-exist")
			},
			true,
		},
		{
			"PathIsFile",
			func(dir string) string {
				file := filepath.Join(dir, "file.txt")
				if err := os.WriteFile(file, []byte("hello"), 0644); err != nil {
					t.Fatalf("Failed to create file: %v", err)
				}
				return file
			},
			true,
		},
		{
			"NilInput",
			func(dir string) string {
				return ""
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			path := tt.setup(dir)

			err := ChangeWorkingDirectory(path)
			if tt.expectedError {
				if err == errors.ERROR_NO_ERROR {
					t.Fatalf("Expected error, got no error")
				}
			} else {
				if err != errors.ERROR_NO_ERROR {
					t.Fatalf("expected no error, got %v", err)
				}
			}
		})
	}
}

func TestChangeOwnership(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(dir string) string
		uid, gid      int
		expectedError bool
	}{
		{
			"ValidInput",
			func(dir string) string {
				file := filepath.Join(dir, "test.txt")
				os.WriteFile(file, []byte("data"), 0644)
				return file
			},
			os.Getuid(), os.Getgid(),
			false,
		},
		{
			"InvalidFilePath",
			func(dir string) string {
				return filepath.Join(dir, "not-exist.txt")
			},
			os.Getuid(), os.Getgid(),
			true,
		},
		{
			"InvalidUser",
			func(dir string) string {
				file := filepath.Join(dir, "abc.txt")
				os.WriteFile(file, []byte("hi"), 0644)
				return file
			},
			99999, 99999,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			dir := t.TempDir()
			path := tt.setup(dir)

			err := ChangeOwnership(path, tt.uid, tt.gid)

			if tt.expectedError {
				if err == errors.ERROR_NO_ERROR {
					t.Fatalf("Expected error, got no error")
				}
			} else {
				if err != errors.ERROR_NO_ERROR {
					t.Fatalf("Expected no error, got %v", err)
				}
			}
		})
	}
}

func TestChangePermission(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(dir string) string
		permission    int
		expectedError bool
	}{
		{
			"ValidInput",
			func(dir string) string {
				file := filepath.Join(dir, "test.txt")
				os.WriteFile(file, []byte("data"), 0644)
				return file
			},
			0755,
			false,
		},
		{
			"InvalidFilePath",
			func(dir string) string {
				return filepath.Join(dir, "not-exist.txt")
			},
			0644,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			dir := t.TempDir()
			path := tt.setup(dir)

			err := ChangePermission(path, tt.permission)

			if tt.expectedError {
				if err == errors.ERROR_NO_ERROR {
					t.Fatalf("Expected error, got no error")
				}
			} else {
				if err != errors.ERROR_NO_ERROR {
					t.Fatalf("Expected no error, got %v", err)
				}
			}
		})
	}
}

func TestIsDirectory(t *testing.T) {
	tests := []struct {
		name           string
		setup          func(dir string) string
		expectedOutput bool
	}{
		{
			"ValidInput",
			func(dir string) string {
				return dir
			},
			true,
		},
		{
			"InvalidInput-FilePath",
			func(dir string) string {
				return filepath.Join(dir, "test.txt")
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			dir := t.TempDir()
			path := tt.setup(dir)

			actualOutput := IsDirectory(path)

			if actualOutput != tt.expectedOutput {
				t.Errorf("Expected: %t, Actual: %t", tt.expectedOutput, actualOutput)
			}
		})
	}
}

func TestIsRegularFile(t *testing.T) {
	tests := []struct {
		name           string
		setup          func(dir string) string
		expectedOutput bool
	}{
		{
			"ValidInput",
			func(dir string) string {
				path := filepath.Join(dir, "test.txt")
				os.WriteFile(path, []byte("hello"), 0644)
				return path
			},
			true,
		},
		{
			"InvalidInput-Directory",
			func(dir string) string {
				return dir
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			dir := t.TempDir()
			path := tt.setup(dir)

			actualOutput := IsRegularFile(path)

			if actualOutput != tt.expectedOutput {
				t.Errorf("Expected: %t, Actual: %t", tt.expectedOutput, actualOutput)
			}
		})
	}
}

func TestGetSHA1SumOfFileAndGetSHA256SumOfFileAndGetMD5SumOfFile(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(dir string) string
		expectedError bool
	}{
		{
			"ValidInput",
			func(dir string) string {
				path := filepath.Join(dir, "test.txt")
				os.WriteFile(path, []byte("hello"), 0644)
				return path
			},
			false,
		},
		{
			"InvalidInput-Directory",
			func(dir string) string {
				return dir
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			dir := t.TempDir()
			path := tt.setup(dir)

			err, _ := GetSHA1SumOfFile(path)
			if tt.expectedError {
				if err == errors.ERROR_NO_ERROR {
					t.Fatalf("Expected error, got no error")
				}
			} else {
				if err != errors.ERROR_NO_ERROR {
					t.Fatalf("Expected no error, got %v", err)
				}
			}

			err, _ = GetSHA256SumOfFile(path)
			if tt.expectedError {
				if err == errors.ERROR_NO_ERROR {
					t.Fatalf("Expected error, got no error")
				}
			} else {
				if err != errors.ERROR_NO_ERROR {
					t.Fatalf("Expected no error, got %v", err)
				}
			}

			err, _ = GetMD5SumOfFile(path)
			if tt.expectedError {
				if err == errors.ERROR_NO_ERROR {
					t.Fatalf("Expected error, got no error")
				}
			} else {
				if err != errors.ERROR_NO_ERROR {
					t.Fatalf("Expected no error, got %v", err)
				}
			}

		})
	}
}

func TestGetFileNameFromAbsolutePath(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			"ValidPath",
			"/home/test/sample.txt",
			"sample.txt",
		},
		{
			"PathWithoutFile",
			"/home/test",
			"test",
		},
		{
			"PathWithTrailingSlash",
			"/home/test/",
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualOutput := GetFileNameFromAbsolutePath(tt.input)
			if actualOutput != tt.expectedOutput {
				t.Errorf("Expected: %s, Actual: %s", tt.expectedOutput, actualOutput)
			}
		})
	}
}

func TestCreateDirectoryIfNeeded(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(string) string
		expectedError bool
	}{
		{
			"CreateNewDirectory",
			func(dir string) string {
				return filepath.Join(dir, "new_folder")
			},
			false,
		},
		{
			"CreateNestedDirectories",
			func(dir string) string {
				return filepath.Join(dir, "a", "b", "c")
			},
			false,
		},
		{
			"DirectoryAlreadyExists",
			func(dir string) string {
				existing := filepath.Join(dir, "exists")
				if err := os.Mkdir(existing, 0755); err != nil {
					t.Fatalf("Setup failed: could not create directory %v", err)
				}
				return existing
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := t.TempDir()
			path := tt.setup(root)

			err := CreateDirectoryIfNeeded(path)
			if tt.expectedError {
				if err == errors.ERROR_NO_ERROR {
					t.Fatalf("Expected error, got no error")
				}
			} else {
				if err != errors.ERROR_NO_ERROR {
					t.Fatalf("Expected no error, got %v", err)
				}
			}

		})
	}
}

func TestOpenFile(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(string) string
		expectedError bool
	}{
		{
			"OpenExistingFile",
			func(dir string) string {
				Path := filepath.Join(dir, "test.txt")

				if err := os.WriteFile(Path, []byte("hello"), 0644); err != nil {
					t.Fatalf("failed to create test file: %v", err)
				}
				return Path
			},
			false,
		},
		{
			"OpenDirectory",
			func(dir string) string {
				return dir
			},
			true,
		},
		{
			"CreateFile",
			func(dir string) string {
				Path := filepath.Join(dir, "test.txt")
				return Path
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			Path := tt.setup(dir)

			_, err := OpenFile(Path, os.O_CREATE|os.O_RDONLY)

			if tt.expectedError {
				if err == errors.ERROR_NO_ERROR {
					t.Fatalf("Expected error, got no error")
				}
			} else {
				if err != errors.ERROR_NO_ERROR {
					t.Fatalf("expected no error, got %v", err)
				}
			}

		})
	}
}

func TestReadFileAsBytesAndAsString(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(string) string
		expectedError bool
	}{
		{
			"OpenExistingFile",
			func(dir string) string {
				Path := filepath.Join(dir, "test.txt")

				if err := os.WriteFile(Path, []byte("hello"), 0644); err != nil {
					t.Fatalf("failed to create test file: %v", err)
				}
				return Path
			},
			false,
		},
		{
			"OpenDirectory",
			func(dir string) string {
				return dir
			},
			true,
		},
		{
			"FileNotExist",
			func(dir string) string {
				Path := filepath.Join(dir, "test.txt")
				return Path
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			Path := tt.setup(dir)

			// ReadFileAsBytes
			_, err := ReadFileAsBytes(Path)

			if tt.expectedError {
				if err == errors.ERROR_NO_ERROR {
					t.Fatalf("Expected error, got no error")
				}
			} else {
				if err != errors.ERROR_NO_ERROR {
					t.Fatalf("expected no error, got %v", err)
				}
			}

			// ReadFileAsString
			_, err = ReadFileAsString(Path)

			if tt.expectedError {
				if err == errors.ERROR_NO_ERROR {
					t.Fatalf("Expected error, got no error")
				}
			} else {
				if err != errors.ERROR_NO_ERROR {
					t.Fatalf("expected no error, got %v", err)
				}
			}
		})
	}
}

func TestWriteToFile(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(string) string
		expectedError bool
	}{
		{
			"FileExist",
			func(dir string) string {
				Path := filepath.Join(dir, "test.txt")

				if err := os.WriteFile(Path, []byte(""), 0644); err != nil {
					t.Fatalf("failed to create test file: %v", err)
				}
				return Path
			},
			false,
		},
		{
			"InValidInput",
			func(dir string) string {
				return dir
			},
			true,
		},
		{
			"FileNotExist",
			func(dir string) string {
				Path := filepath.Join(dir, "test.txt")
				return Path
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			Path := tt.setup(dir)

			err := WriteToFile(Path, "Hello")

			if tt.expectedError {
				if err == errors.ERROR_NO_ERROR {
					t.Fatalf("Expected error, got no error")
				}
			} else {
				if err != errors.ERROR_NO_ERROR {
					t.Fatalf("expected no error, got %v", err)
				}
			}

		})
	}
}

func TestCopyFile(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(string) (string, string)
		expectedError bool
	}{
		{
			"ValidInput",
			func(dir string) (string, string) {
				src := filepath.Join(dir, "src.txt")
				dest := filepath.Join(dir, "dest.txt")

				if err := os.WriteFile(src, []byte("hello"), 0644); err != nil {
					t.Fatalf("failed to create test file: %v", err)
				}
				return src, dest
			},
			false,
		},
		{
			"SourceNotExists",
			func(dir string) (string, string) {
				src := filepath.Join(dir, "src.txt")
				dest := filepath.Join(dir, "dest.txt")
				return src, dest
			},
			true,
		},
		{
			"NillInput",
			func(dir string) (string, string) {
				return "", ""
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			src, dest := tt.setup(dir)

			err := CopyFile(src, dest)

			if tt.expectedError {
				if err == errors.ERROR_NO_ERROR {
					t.Fatalf("Expected error, got no error")
				}
			} else {
				if err != errors.ERROR_NO_ERROR {
					t.Fatalf("expected no error, got %v", err)
				}
			}
		})
	}
}

func TestCreateSymLink(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(string) (string, string)
		expectedError bool
	}{
		{
			"ValidInput",
			func(dir string) (string, string) {
				src := filepath.Join(dir, "src.txt")
				dest := filepath.Join(dir, "dest.txt")

				if err := os.WriteFile(src, []byte("hello"), 0644); err != nil {
					t.Fatalf("failed to create test file: %v", err)
				}
				return src, dest
			},
			false,
		},
		{
			"SourceNotExists-BrokenSymLink",
			func(dir string) (string, string) {
				src := filepath.Join(dir, "src.txt")
				dest := filepath.Join(dir, "dest.txt")
				return src, dest
			},
			true,
		},
		{
			"NillInput",
			func(dir string) (string, string) {
				return "", ""
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			src, dest := tt.setup(dir)

			_ = CreateSymLink(src, dest)

			_, err := ReadFileAsString(dest)

			if tt.expectedError {
				if err == errors.ERROR_NO_ERROR {
					t.Fatalf("Expected error, got no error")
				}
			} else {
				if err != errors.ERROR_NO_ERROR {
					t.Fatalf("expected no error, got %v", err)
				}
			}
		})
	}
}
